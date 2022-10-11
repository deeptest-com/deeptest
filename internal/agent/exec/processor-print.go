package agentExec

import "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorPrint) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	return
}
