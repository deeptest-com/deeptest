package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type ExecLogProcessor struct {
	BaseModel

	Name           string                `json:"name"`
	Desc           string                `json:"desc"`
	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" gorm:"default:pass"`
	StartTime      *time.Time            `json:"startTime"`
	EndTime        *time.Time            `json:"endTime"`

	ParentId uint `json:"parentId"`
	InvokeId uint `json:"invokeId"`
	ReportId uint `json:"reportId"`
	UseID    uint `json:"useId"`

	// type
	ProcessorCategory consts.ProcessorCategory `json:"processorCategory" yaml:"processorCategory"`

	// for interface
	EndpointInterfaceId   uint                `json:"endpointInterfaceId"`
	DebugInterfaceId      uint                `json:"debugInterfaceId"`
	ReqContent            string              `json:"reqContent,omitempty" gorm:"type:mediumtext"`
	RespContent           string              `json:"respContent,omitempty" gorm:"type:mediumtext"`
	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	InterfaceExtractorsResult  []ExecLogExtractor  `gorm:"-" json:"interfaceExtractorsResult,omitempty"`
	InterfaceCheckpointsResult []ExecLogCheckpoint `gorm:"-" json:"interfaceCheckpointsResult,omitempty"`

	// for processor
	ProcessorType       consts.ProcessorType `json:"processorType" yaml:"processorType"`
	ScenarioProcessorId uint                 `gorm:"default:0" json:"scenarioProcessorId,omitempty"`
	ScenarioId          uint                 `gorm:"default:0" json:"scenarioId,omitempty"`
	//ProcessorContent string               `json:"processorContent,omitempty"`
	//ProcessorResult  string               `json:"processorResult,omitempty"`

	Summary string `json:"summary,omitempty"`
	Detail  string `gorm:"type:text" json:"Detail,omitempty"`
	Output  string `json:"output,omitempty"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (ExecLogProcessor) TableName() string {
	return "biz_exec_log_processor"
}

type ExecLogExtractor struct {
	DebugConditionExtractor
	InvokeId uint `json:"invokeId"`
}

func (ExecLogExtractor) TableName() string {
	return "biz_exec_log_extractor"
}

type ExecLogCookie struct {
	DebugConditionCookie
	InvokeId uint `json:"invokeId"`
}

func (ExecLogCookie) TableName() string {
	return "biz_exec_log_cookie"
}

type ExecLogCheckpoint struct {
	DebugConditionCheckpoint
	InvokeId uint `json:"invokeId"`
}

func (ExecLogCheckpoint) TableName() string {
	return "biz_exec_log_checkpoint"
}

type ExecLogScript struct {
	DebugConditionScript
	InvokeId uint `json:"invokeId"`
}

func (ExecLogScript) TableName() string {
	return "biz_exec_log_script"
}

type ExecLogResponseDefine struct {
	DebugConditionResponseDefine
	InvokeId uint `json:"invokeId"`
}

func (ExecLogResponseDefine) TableName() string {
	return "biz_exec_log_response_define"
}
