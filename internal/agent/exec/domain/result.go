package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Result struct {
	ID uint `json:"id" yaml:"id"`

	Name           string                `json:"name"`
	Desc           string                `json:"desc,omitempty"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus,omitempty"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`
	StartTime      *time.Time            `json:"startTime,omitempty"`
	EndTime        *time.Time            `json:"endTime,omitempty"`

	ParentId uint `json:"parentId"`
	ReportId uint `json:"reportId"`
	UseID    uint `json:"useId,omitempty"`

	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType"`

	// for interface
	InterfaceId           uint                `json:"interfaceId,omitempty"`
	ReqContent            string              `json:"reqContent,omitempty"`
	RespContent           string              `json:"respContent,omitempty"`
	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode,omitempty"`
	HttpRespStatusContent string              `json:"httpStatusContent,omitempty"`

	ExtractorsResult  []Extractor  `json:"extractorsResult,omitempty"`
	CheckpointsResult []Checkpoint `json:"checkpointsResult,omitempty"`

	// for processor
	ProcessorId      uint   `json:"processorId,omitempty"`
	ProcessorContent string `json:"processorContent,omitempty"`
	ProcessorResult  string `json:"processorResult,omitempty"`

	// for loop processor
	Iterator ExecIterator `json:"iterator,omitempty"`
	// for loop break processor
	WillBreak bool `json:"break,omitempty"`

	Summary string `json:"summary,omitempty"`

	Children []*Result `json:"logs,omitempty"`
}
