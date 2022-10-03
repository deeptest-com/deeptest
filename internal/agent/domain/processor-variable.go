package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ProcessorVariable struct {
	Id uint
	model.ProcessorEntity

	VariableName string `json:"variableName" yaml:"variableName"`
	RightValue   string `json:"rightValue" yaml:"rightValue"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (s *ProcessorVariable) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
