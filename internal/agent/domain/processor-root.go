package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorRootStage struct {
	Stage *run.TStage
}

func (s *ProcessorRootStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorRootStage) Category() consts.ProcessorCategory {
	return consts.ProcessorRoot
}

func (s *ProcessorRootStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorRootStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorRoot)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorRoot struct {
	Id uint

	model.ProcessorEntity

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
