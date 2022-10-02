package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorTimerStage struct {
	stage *run.TStage
}

func (s *ProcessorTimerStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorTimerStage) Category() consts.ProcessorCategory {
	return consts.ProcessorTimer
}

func (s *ProcessorTimerStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorTimerStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorTimer)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorTimer struct {
	Id uint
	model.ProcessorEntity

	SleepTime int `json:"sleepTime" yaml:"sleepTime"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
