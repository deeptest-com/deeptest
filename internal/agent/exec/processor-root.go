package agentExec

type ProcessorRoot struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (p *ProcessorRoot) Run(s *Session) (ret *Log, err error) {
	return
}
