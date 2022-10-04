package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorLogic struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (s *ProcessorLogic) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
