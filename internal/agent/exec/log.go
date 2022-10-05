package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"time"
)

type Log struct {
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

	InterfaceExtractorsResult  []ExtractLog    `json:"interfaceExtractorsResult,omitempty"`
	InterfaceCheckpointsResult []CheckpointLog `json:"interfaceCheckpointsResult,omitempty"`

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

	Logs []*Log `json:"logs"`
}

type ExtractLog struct {
	ID   uint                 `json:"id" yaml:"id"`
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

type CheckpointLog struct {
	ID   uint                  `json:"id" yaml:"id"`
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
