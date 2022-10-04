package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorGroup struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (s *ProcessorGroup) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
