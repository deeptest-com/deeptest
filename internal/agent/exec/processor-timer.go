package agentExec

import "github.com/aaronchen2k/deeptest/internal/agent/exec/domain"

type ProcessorTimer struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntityBase

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`
}

func (entity ProcessorTimer) Run(processor *Processor, session *Session) (log domain.Result, err error) {
	return
}
