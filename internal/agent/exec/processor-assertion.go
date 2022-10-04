package agentExec

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (p ProcessorAssertion) Run(s *Session) (ret *Log, err error) {
	return
}
