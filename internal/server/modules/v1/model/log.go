package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Log struct {
	BaseModel

	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" gorm:"default:pass"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	ParentId uint `json:"parentId"`
	ReportId uint `json:"reportId"`
	UseID    uint `json:"useId"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`

	// for interface
	InterfaceId           uint                `json:"interfaceId"`
	ReqContent            string              `json:"reqContent,omitempty"`
	RespContent           string              `json:"respContent,omitempty"`
	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	InterfaceExtractorsResult  []LogExtractor  `gorm:"-" json:"interfaceExtractorsResult,omitempty"`
	InterfaceCheckpointsResult []LogCheckpoint `gorm:"-" json:"interfaceCheckpointsResult,omitempty"`

	// for processor
	ProcessorType    consts.ProcessorType `json:"processorType" yaml:"processorType"`
	ProcessorId      uint                 `json:"processorId,omitempty"`
	ProcessorContent string               `json:"processorContent,omitempty"`
	ProcessorResult  string               `json:"processorResult,omitempty"`

	Summary string `json:"summary,omitempty"`
	Output  string `json:"output,omitempty"`

	Logs []*Log `gorm:"-" json:"logs"`
}

func (Log) TableName() string {
	return "biz_log"
}

type LogExtractor struct {
	BaseModel
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
	LogId       uint   `json:"logId"`
}

func (LogExtractor) TableName() string {
	return "biz_log_extractor"
}

type LogCheckpoint struct {
	BaseModel
	Type consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	ActualResult string              `json:"actualResult"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	InterfaceId  uint                `json:"interfaceId"`
	LogId        uint                `json:"logId"`
}

func (LogCheckpoint) TableName() string {
	return "biz_log_checkpoint"
}
