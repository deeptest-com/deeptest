package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorPrintStage struct {
	Stage *run.TStage
}

func (s *ProcessorPrintStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorPrintStage) Category() consts.ProcessorCategory {
	return consts.ProcessorPrint
}

func (s *ProcessorPrintStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorPrintStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorPrint)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
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
