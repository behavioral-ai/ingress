package operative1

import (
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/ingress/timeseries1"
	"time"
)

const (
	Class           = "operative1"
	defaultDuration = time.Second * 10
)

type service struct {
	running bool
	agentId string
	origin  core.Origin
	filter  messaging.TraceFilter

	duration time.Duration
	handler  messaging.OpsAgent
	emissary *communications
	master   *communications
}

func serviceAgentUri(origin core.Origin) string {
	return origin.Uri(Class)
}

// New - create a new operative1 agent
func New(origin core.Origin, handler messaging.OpsAgent, global messaging.Dispatcher) messaging.Agent {
	return newOp(origin, handler, global, newMasterDispatcher(false), newEmissaryDispatcher(false))
}

func newOp(origin core.Origin, handler messaging.OpsAgent, global messaging.Dispatcher, master, emissary dispatcher) *service {
	r := new(service)
	r.origin = origin
	r.agentId = serviceAgentUri(origin)
	r.duration = defaultDuration

	r.handler = handler
	r.emissary = newEmmissaryComms(global, emissary)
	r.master = newMasterComms(global, master)
	return r
}

// String - identity
func (s *service) String() string { return s.Uri() }

// Uri - agent identifier
func (s *service) Uri() string { return s.agentId }

// Message - message the agent
func (s *service) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.To() {
	case messaging.EmissaryChannel:
		s.emissary.send(m)
	case messaging.MasterChannel:
		s.master.send(m)
	default:
		s.emissary.send(m)
	}
}

// Run - run the agent
func (s *service) Run() {
	if s.running {
		return
	}
	go masterAttend(s)
	go emissaryAttend(s, timeseries1.Observe)
	s.running = true
}

// Shutdown - shutdown the agent
func (s *service) Shutdown() {
	if !s.running {
		return
	}
	s.running = false
	msg := messaging.NewControlMessage(s.Uri(), s.Uri(), messaging.ShutdownEvent)
	s.emissary.enable()
	s.emissary.send(msg)
	s.master.enable()
	s.master.send(msg)
}

func (s *service) IsFinalized() bool {
	return s.emissary.isFinalized() && s.master.isFinalized()
}
