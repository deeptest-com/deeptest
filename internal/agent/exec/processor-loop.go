package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
)

type ProcessorLoop struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

	Times        int    `json:"times" yaml:"times"` // time
	Range        string `json:"range" yaml:"range"` // range
	List         string `json:"list" yaml:"list"`   // in
	Step         string `json:"step" yaml:"step"`
	IsRand       bool   `json:"isRand" yaml:"isRand"`
	VariableName string `json:"variableName" yaml:"variableName"`

	UntilExpression   string `json:"untilExpression" yaml:"untilExpression"` // until
	BreakIfExpression string `json:"breakIfExpression" yaml:"breakIfExpression"`
}

func (s *ProcessorLoop) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
