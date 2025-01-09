package operative1

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/messaging"
	"reflect"
)

const (
	contentTypeObservation = "application/observation"
)

type observation struct {
	Latency  int
	Gradient int
}

func getObservation(h messaging.Notifier, agentId string, msg *messaging.Message) (observation, *core.Status) {
	if !msg.IsContentType(contentTypeObservation) {
		return observation{}, core.StatusNotFound()
	}
	if p, ok := msg.Body.(observation); ok {
		return p, core.StatusOK()
	}
	status := observationTypeErrorStatus(agentId, msg.Body)
	h.Notify(status)
	return observation{}, status
}

func observationTypeErrorStatus(agentId string, t any) *core.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}
