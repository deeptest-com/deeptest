package agentExec

type ProcessorTimer struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`
}

func (p ProcessorTimer) Run(s *Session) (log Log, err error) {
	return
}
