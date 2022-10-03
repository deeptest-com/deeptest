package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorVariableStage struct {
	Stage *run.TStage
}

func (s *ProcessorVariableStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorVariableStage) Category() consts.ProcessorCategory {
	return consts.ProcessorVariable
}

func (s *ProcessorVariableStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorVariableStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorVariable)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
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
