package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorPrintStage struct {
	stage *run.TStage
}

func (s *ProcessorPrintStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorPrintStage) Category() consts.ProcessorCategory {
	return consts.ProcessorPrint
}

func (s *ProcessorPrintStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorPrintStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorPrint)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorPrint struct {
	Id uint
	model.ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
