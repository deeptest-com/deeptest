package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorLogicStage struct {
	stage *run.TStage
}

func (s *ProcessorLogicStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorLogicStage) Category() consts.ProcessorCategory {
	return consts.ProcessorLogic
}

func (s *ProcessorLogicStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorLogicStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorLogic)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorLogic struct {
	Id uint
	model.ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
