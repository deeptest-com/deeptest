package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorVariable struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	VariableName string `json:"variableName" yaml:"variableName"`
	RightValue   string `json:"rightValue" yaml:"rightValue"`
}

func (s *ProcessorVariable) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
