package agentExec

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (p ProcessorPrint) Run(s *Session) (variableName string, variableValues []interface{}, err error) {
	return
}
