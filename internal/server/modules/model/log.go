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

func (ExecLogExtractor) TableName() string {
	return "biz_exec_log_extractor"
}

type ExecLogCheckpoint struct {
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

func (ExecLogCheckpoint) TableName() string {
	return "biz_exec_log_checkpoint"
}

type ExecLogScript struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	Output       string              `json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	InterfaceId  uint                `json:"interfaceId"`
	LogId        uint                `json:"logId"`
}

func (ExecLogScript) TableName() string {
	return "biz_exec_log_script"
}
