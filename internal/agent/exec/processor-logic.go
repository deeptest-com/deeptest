package agentExec

type ProcessorLogic struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (p ProcessorLogic) Run(s *Session) (log Log, variableName string, variableValues []interface{}, err error) {
	return
}
