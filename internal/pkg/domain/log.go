package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Log struct {
	ID int `json:"id"`

	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	Logs []Log `json:"logs"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`

	// for interface
	InterfaceId   uint   `json:"interfaceId"`
	InterfaceReq  string `json:"interfaceReq,omitempty"`
	InterfaceResp string `json:"interfaceResp,omitempty"`

	// for processor
	ProcessorType  consts.ProcessorType `json:"processorType"`
	ProcessId      uint                 `json:"processId,omitempty"`
	ProcessContent string               `json:"processContent,omitempty"`
	ProcessResult  string               `json:"processResult,omitempty"`
}
