package agentExec

import "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	Expression string `json:"expression" yaml:"expression"`
}

func (entity ProcessorAssertion) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	return
}
