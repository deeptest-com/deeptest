package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type ProcessorExtractor struct {
	ID uint `json:"id" yaml:"id"`
	ProcessorEntity

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
	InterfaceID uint   `json:"interfaceID"`
}

func (p ProcessorExtractor) Run(s *Session) (log Result, err error) {
	return
}
