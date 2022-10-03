package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorAssertion struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`
}

func (s *ProcessorAssertion) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
