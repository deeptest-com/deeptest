package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type TestLog struct {
	BaseModel

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`

	// for interface
	InterfaceId   uint   `json:"interfaceId"`
	InterfaceReq  string `json:"interfaceReq,omitempty"`
	InterfaceResp string `json:"interfaceResp,omitempty"`

	// for processor
	ProcessorType  consts.ProcessorType `json:"processorType" yaml:"processorType"`
	ProcessId      uint                 `json:"processId,omitempty"`
	ProcessContent string               `json:"processContent,omitempty"`
	ProcessResult  string               `json:"processResult,omitempty"`

	Status    consts.ResultStatus `json:"status"`
	StartTime *time.Time          `json:"startTime"`
	EndTime   *time.Time          `json:"endTime"`

	ResultId uint `json:"resultId"`
	UseID    uint `json:"useId"`
}

func (TestLog) TableName() string {
	return "biz_test_log"
}
