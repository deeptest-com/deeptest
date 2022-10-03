package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorGroupStage struct {
	Stage *run.TStage
}

func (s *ProcessorGroupStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorGroupStage) Category() consts.ProcessorCategory {
	return consts.ProcessorGroup
}

func (s *ProcessorGroupStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorGroupStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorGroup)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorGroup struct {
	Id uint
	model.ProcessorEntity

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
