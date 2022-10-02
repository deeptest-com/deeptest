package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorInterfaceStage struct {
	stage *run.TStage
}

func (s *ProcessorInterfaceStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorInterfaceStage) Category() consts.ProcessorCategory {
	return consts.ProcessorInterface
}

func (s *ProcessorInterfaceStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorInterfaceStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorInterface)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorInterface struct {
	Id uint
	model.ProcessorEntity

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
