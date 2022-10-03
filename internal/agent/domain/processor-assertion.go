package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorAssertionStage struct {
	Stage *run.TStage
}

func (s *ProcessorAssertionStage) Name() string {
	return s.Stage.Name
}

func (s *ProcessorAssertionStage) Category() consts.ProcessorCategory {
	return consts.ProcessorAssertion
}

func (s *ProcessorAssertionStage) Struct() *run.TStage {
	return s.Stage
}

func (s *ProcessorAssertionStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.Stage.Processor.(ProcessorAssertion)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.Stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorAssertion struct {
	Id uint
	model.ProcessorEntity

	Expression string `json:"expression" yaml:"expression"`

	Children []interface{} `json:"children" yaml:"children"`
}
