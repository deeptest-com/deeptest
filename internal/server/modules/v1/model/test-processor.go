package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type TestProcessor struct {
	BaseModel

	Name          string               `json:"name" yaml:"name"`
	Desc          string               `json:"desc" yaml:"desc"`
	ProcessorType consts.ProcessorType `json:"processorType" yaml:"processorType"`
	ProcessorId   uint                 `json:"processorId" yaml:"processorId"`
}

func (TestProcessor) TableName() string {
	return "test_processor"
}

type ProcessorSimple struct {
	BaseModel
	ProcessorBase
}

func (ProcessorSimple) TableName() string {
	return "biz_test_processor_simple"
}

type ProcessorFlow struct {
	BaseModel
	ProcessorBase

	Condition string `json:"condition" yaml:"condition"`
}

func (ProcessorFlow) TableName() string {
	return "biz_test_processor_flow"
}

type ProcessorIterator struct {
	BaseModel
	ProcessorBase
}

func (ProcessorIterator) TableName() string {
	return "biz_test_processor_iterator"
}

type ProcessorTimer struct {
	BaseModel
	ProcessorBase
}

func (ProcessorTimer) TableName() string {
	return "biz_test_processor_timer"
}

type ProcessorAssertion struct {
	BaseModel
	ProcessorBase
}

func (ProcessorAssertion) TableName() string {
	return "biz_test_processor_assertion"
}

type ProcessorExtractor struct {
	BaseModel
	ProcessorBase

	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `json:"expression"`
	Prop       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string `json:"variable"`

	Result      string `json:"result"`
	InterfaceId uint   `json:"interfaceId"`
}

func (ProcessorExtractor) TableName() string {
	return "biz_test_processor_extractor"
}

type ProcessorData struct {
	BaseModel
	ProcessorBase

	Type consts.DataSource `json:"type,omitempty" yaml:"type,omitempty"`
	Path string            `json:"path,omitempty" yaml:"path,omitempty"`

	Loop           int    `json:"loop,omitempty" yaml:"loop,omitempty"`
	StartIndex     int    `json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	EndIndex       int    `json:"endIndex,omitempty" yaml:"endIndex,omitempty"`
	IsRand         bool   `json:"isRand,omitempty" yaml:"isRand,omitempty"`
	IsOnce         bool   `json:"isOnce,omitempty" yaml:"isOnce,omitempty"`
	VarNamePostfix string `json:"varNamePostfix,omitempty" yaml:"varNamePostfix,omitempty"`
}

func (ProcessorData) TableName() string {
	return "biz_test_processor_data"
}

type ProcessorCookie struct {
	BaseModel
	ProcessorBase
}

func (ProcessorCookie) TableName() string {
	return "biz_test_processor_cookie"
}

type ProcessorBase struct {
	Name string `json:"name" yaml:"name"`

	ParentId uint `json:"parentId" yaml:"parentId"`

	// interface or Processor
	Children []interface{} `json:"children" yaml:"children" gorm:"-"`
	// results
	Results []string `json:"results" yaml:"results" gorm:"-"`
}
