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
