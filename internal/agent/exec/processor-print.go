package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorPrint struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (s *ProcessorPrint) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
