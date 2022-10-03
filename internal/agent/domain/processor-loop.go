package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorLoopStage struct {
	Stage *run.TStage
}

func (s *ProcessorLoopStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorLoopStage) Category() consts.ProcessorCategory {
	return consts.ProcessorLoop
}

func (s *ProcessorLoopStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorLoopStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorLoop)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorLoop struct {
	Id uint
	model.ProcessorEntity

	Times        int    `json:"times" yaml:"times"` // time
	Range        string `json:"range" yaml:"range"` // range
	List         string `json:"list" yaml:"list"`   // in
	Step         string `json:"step" yaml:"step"`
	IsRand       bool   `json:"isRand" yaml:"isRand"`
	VariableName string `json:"variableName" yaml:"variableName"`

	UntilExpression   string `json:"untilExpression" yaml:"untilExpression"` // until
	BreakIfExpression string `json:"breakIfExpression" yaml:"breakIfExpression"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
