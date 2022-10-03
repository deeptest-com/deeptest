package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ProcessorRoot struct {
	Id uint

	model.ProcessorEntity

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (s *ProcessorRoot) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
