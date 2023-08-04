package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ScenarioExecResult struct {
	ID int `json:"id" yaml:"id"`

	ScenarioId       uint                  `json:"scenarioId"`
	ScenarioReportId uint                  `json:"scenarioReportId"`
	Name             string                `json:"name"`
	Desc             string                `json:"desc,omitempty"`
	ProgressStatus   consts.ProgressStatus `json:"progressStatus,omitempty"`
	ResultStatus     consts.ResultStatus   `json:"resultStatus"`
	StartTime        *time.Time            `json:"startTime,omitempty"`
	EndTime          *time.Time            `json:"endTime,omitempty"`

	ParentId int `json:"parentId"`
	//ReportId uint `json:"reportId"`
	UseID uint `json:"useId,omitempty"`

	ProcessorCategory consts.ProcessorCategory `json:"processorCategory"`
	ProcessorType     consts.ProcessorType     `json:"processorType"`

	// for interface
	EndpointInterfaceId uint `json:"endpointInterfaceId,omitempty"`
	DebugInterfaceId    uint `json:"debugInterfaceId,omitempty"`

	ReqContent            string              `json:"reqContent,omitempty"`
	RespContent           string              `json:"respContent,omitempty"`
	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode,omitempty"`
	HttpRespStatusContent string              `json:"httpStatusContent,omitempty"`

	ExtractorsResult  []domain.ExtractorBase  `json:"extractorsResult,omitempty"`
	CheckpointsResult []domain.CheckpointBase `json:"checkpointsResult,omitempty"`
	ScriptsResult     []domain.ScriptBase     `json:"scriptsResult,omitempty"`

	// for processor
	ProcessorId      uint   `json:"processorId,omitempty"`
	ProcessorContent string `json:"processorContent,omitempty"`
	ProcessorResult  string `json:"processorResult,omitempty"`

	// for loop processor
	Iterator ExecIterator `json:"iterator,omitempty"`
	// for loop break processor
	WillBreak bool `json:"break,omitempty"`

	Summary string `json:"summary,omitempty"`

	Children []*ScenarioExecResult `json:"logs,omitempty"`

	EnvironmentId int `json:"environmentId,omitempty"`

	LogId       uuid.UUID `json:"logId,omitempty"`
	ParentLogId uuid.UUID `json:"parentLogId,omitempty"`

	Cost int64 `json:"cost,omitempty"`

	Detail string `json:"detail,omitempty"`
}

type PlanExecResult struct {
	ID int `json:"id" yaml:"id"`

	Name          string `json:"name"`
	Desc          string `json:"desc,omitempty"`
	EnvironmentId int    `json:"environmentId"`

	Scenarios []*ScenarioExecResult `json:"scenarios"`
}

type MessageExecResult struct {
	UserId uint   `json:"userId"`
	Name   string `json:"name"`
}
