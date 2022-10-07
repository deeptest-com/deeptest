package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Result struct {
	ID uint `json:"id" yaml:"id"`

	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	ParentId uint `json:"parentId"`
	ReportId uint `json:"reportId"`
	UseID    uint `json:"useId"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`

	// for interface
	InterfaceId           uint                `json:"interfaceId,omitempty"`
	ReqContent            string              `json:"reqContent,omitempty"`
	RespContent           string              `json:"respContent,omitempty"`
	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode,omitempty"`
	HttpRespStatusContent string              `json:"httpStatusContent,omitempty"`

	ExtractorsResult  []domain.Extractor  `json:"extractorsResult,omitempty"`
	CheckpointsResult []domain.Checkpoint `json:"checkpointsResult,omitempty"`

	// for processor
	ProcessorType    consts.ProcessorType `json:"processorType,omitempty"`
	ProcessorId      uint                 `json:"processorId,omitempty"`
	ProcessorContent string               `json:"processorContent,omitempty"`
	ProcessorResult  string               `json:"processorResult,omitempty"`

	// for loop processor
	Iterator domain.ExecIterator `json:"iterator,omitempty"`
	// for loop break processor
	WillBreak bool `json:"break,omitempty"`

	Summary string `json:"summary,omitempty"`
	Output  string `json:"output,omitempty"`

	Logs []*Result `json:"logs"`
}
