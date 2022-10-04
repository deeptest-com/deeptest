package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorRoot struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (s *ProcessorRoot) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
