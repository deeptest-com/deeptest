package runDomain

import (
	"github.com/aaronchen2k/deeptest/internal/agent/run"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"log"
)

type ProcessorExtractorStage struct {
	stage *run.TStage
}

func (s *ProcessorExtractorStage) Name() string {
	return s.stage.Name
}

func (s *ProcessorExtractorStage) Category() consts.ProcessorCategory {
	return consts.ProcessorExtractor
}

func (s *ProcessorExtractorStage) Struct() *run.TStage {
	return s.stage
}

func (s *ProcessorExtractorStage) Run(r *run.SessionRunner) (ret *run.StageResult, err error) {
	processor, ok := s.stage.Processor.(ProcessorExtractor)
	if ok {
		log.Println(processor)
	}

	for _, child := range s.stage.Children {
		log.Println(child)

		child.Run(r)
	}

	return
}

type ProcessorExtractor struct {
	Id uint
	model.ProcessorEntity

	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"` // form header

	Expression string `json:"expression"`
	//Prop       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string `json:"variable"`

	Result      string `json:"result"`
	InterfaceId uint   `json:"interfaceId"`

	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
}
