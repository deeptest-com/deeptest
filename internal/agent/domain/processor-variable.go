package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorVariableStage struct {
	stage *run.TStage
}

func (s *ProcessorVariableStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorVariableStage) Category() consts.ProcessorCategory {
	return consts.ProcessorVariable
}

func (s *ProcessorVariableStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorVariableStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorVariable)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorVariable struct {
	Id uint
	model.ProcessorEntity

	VariableName string `json:"variableName" yaml:"variableName"`
	RightValue   string `json:"rightValue" yaml:"rightValue"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
