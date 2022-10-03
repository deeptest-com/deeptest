package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorInterface struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity
}

func (s *ProcessorInterface) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
