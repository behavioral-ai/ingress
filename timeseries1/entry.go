package timeseries1

type Entry struct {
	Latency  int
	Gradient int
}

/*

func GetObservation(h messaging.Notifier, agentId string, msg *messaging.Message) (Entry, *core.Status) {
	if !msg.IsContentType(ContentTypeObservation) {
		return Entry{}, core.StatusNotFound()
	}
	if p, ok := msg.Body.(Entry); ok {
		return p, core.StatusOK()
	}
	status := observationTypeErrorStatus(agentId, msg.Body)
	h.Notify(status)
	return Entry{}, status
}

func observationTypeErrorStatus(agentId string, t any) *core.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}


*/
