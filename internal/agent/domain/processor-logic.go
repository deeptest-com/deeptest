package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ProcessorLogic struct {
	Id uint
	model.ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (s *ProcessorLogic) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
