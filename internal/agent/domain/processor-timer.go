package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorTimer struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`
}

func (s *ProcessorTimer) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
