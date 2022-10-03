package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type ProcessorTimer struct {
	Id uint
	model.ProcessorEntity

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}

func (s *ProcessorTimer) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	return
}
