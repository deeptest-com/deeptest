package agentExec

type ProcessorLogic struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (p ProcessorLogic) Run(s *Session) (log Result, err error) {
	return
}
