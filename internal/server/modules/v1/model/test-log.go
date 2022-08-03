package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type TestLog struct {
	BaseModel

	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	ParentId uint `json:"parentId"`
	ResultId uint `json:"resultId"`
	UseID    uint `json:"useId"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`

	// for interface
	InterfaceId uint   `json:"interfaceId"`
	ReqContent  string `json:"reqContent,omitempty"`
	RespContent string `json:"respContent,omitempty"`

	// for processor
	ProcessorType  consts.ProcessorType `json:"processorType" yaml:"processorType"`
	ProcessId      uint                 `json:"processId,omitempty"`
	ProcessContent string               `json:"processContent,omitempty"`
	ProcessResult  string               `json:"processResult,omitempty"`
}

func (TestLog) TableName() string {
	return "biz_test_log"
}
