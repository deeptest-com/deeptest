package agentExec

type ProcessorRoot struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (p *ProcessorRoot) Run(s *Session) (variableName string, variableValues []interface{}, err error) {
	return
}
