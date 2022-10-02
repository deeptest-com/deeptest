package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"log"
)

type ProcessorGroupStage struct {
	stage *run.TStage
}

func (s *ProcessorGroupStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorGroupStage) Category() consts.ProcessorCategory {
	return consts.ProcessorGroup
}

func (s *ProcessorGroupStage) Type() consts.ProcessorType {
	return consts.ProcessorGroupDefault
}

func (s *ProcessorGroupStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorGroupStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorGroup)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorGroup struct {
	Id uint

	Name     string `gorm:"-" json:"name" yaml:"name"`
	Comments string `json:"comments" yaml:"comments"`
	Default  string `json:"default" yaml:"default"`

	ProcessorId       uint                     `json:"processorId" yaml:"processorId"`
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType" yaml:"processorType"`

	ParentId uint `json:"parentId" yaml:"parentId"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
