package agentExec

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (p ProcessorAssertion) Run(s *Session) (variableName string, variableValues []interface{}, err error) {
	return
}
