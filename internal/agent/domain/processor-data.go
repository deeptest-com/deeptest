package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorDataStage struct {
	stage *run.TStage
}

func (s *ProcessorDataStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorDataStage) Category() consts.ProcessorCategory {
	return consts.ProcessorData
}

func (s *ProcessorDataStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorDataStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorData)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorData struct {
	Id uint
	model.ProcessorEntity

	Type      consts.DataSource `json:"type,omitempty" yaml:"type,omitempty"`
	Url       string            `json:"url,omitempty" yaml:"url,omitempty"`
	Separator string            `json:"separator,omitempty" yaml:"separator,omitempty"`

	RepeatTimes int `json:"repeatTimes,omitempty" yaml:"repeatTimes,omitempty"`
	//StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	//EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`

	IsLoop int  `json:"isLoop,omitempty" yaml:"isLoop,omitempty"`
	IsRand bool `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce bool `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`

	VariableName string `json:"variableName,omitempty" yaml:"variableName,omitempty"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
