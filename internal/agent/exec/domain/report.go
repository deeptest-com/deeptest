package agentExecDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Report struct {
	ID   int    `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Desc string `json:"desc" yaml:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus" yaml:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" yaml:"resultStatus"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Duration  int64      `json:"duration"` // 毫秒

	TotalScenarioNum int `json:"totalScenarioNum"`
	PassScenarioNum  int `json:"passScenarioNum"`
	FailScenarioNum  int `json:"failScenarioNum" yaml:"failScenarioNum"`

	TotalInterfaceNum int `json:"totalInterfaceNum"`
	PassInterfaceNum  int `json:"passInterfaceNum"`
	FailInterfaceNum  int `json:"failInterfaceNum" yaml:"failInterfaceNum"`

	TotalRequestNum int `json:"totalRequestNum"`
	PassRequestNum  int `json:"passRequestNum"`
	FailRequestNum  int `json:"failRequestNum"`

	TotalAssertionNum int `json:"totalAssertionNum"`
	PassAssertionNum  int `json:"passAssertionNum"`
	FailAssertionNum  int `json:"failAssertionNum"`

	TotalProcessorNum  int `json:"totalProcessorNum"`
	FinishProcessorNum int `json:"finishProcessorNum"`

	InterfaceStatusMap map[uint]map[consts.ResultStatus]int `gorm:"-"`

	Payload interface{} `json:"payload"`

	ScenarioId uint   `json:"scenarioId"`
	ProjectId  uint   `json:"projectId"`
	PlanId     uint   `json:"planId"`
	PlanName   string `json:"planName"`

	ExecEnv  string `json:"execEnv"`
	Priority string `json:"priority"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

type ExecLogProcessor struct {
	ID             int                   `json:"id" yaml:"id"`
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
	ReqContent            string              `json:"reqContent,omitempty" gorm:"type:mediumtext"`
	RespContent           string              `json:"respContent,omitempty" gorm:"type:mediumtext"`
	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	InterfaceExtractorsResult  []ExecLogExtractor  `gorm:"-" json:"interfaceExtractorsResult,omitempty"`
	InterfaceCheckpointsResult []ExecLogCheckpoint `gorm:"-" json:"interfaceCheckpointsResult,omitempty"`

	// for processor
	ProcessorType consts.ProcessorType `json:"processorType" yaml:"processorType"`
	ProcessorId   uint                 `json:"processorId,omitempty"`
	//ProcessorContent string               `json:"processorContent,omitempty"`
	//ProcessorResult  string               `json:"processorResult,omitempty"`

	Summary string `json:"summary,omitempty"`
	Output  string `json:"output,omitempty"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

type ExecLogExtractor struct {
	ID   int                  `json:"id" yaml:"id"`
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

type ExecLogCheckpoint struct {
	ID   int                   `json:"id" yaml:"id"`
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

type ReportSimple struct {
	ID   int    `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Desc string `json:"desc" yaml:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus" yaml:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" yaml:"resultStatus"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Duration  int64      `json:"duration"` // sec

	TotalInterfaceNum int `json:"totalInterfaceNum"`
	PassInterfaceNum  int `json:"passInterfaceNum"`
	FailInterfaceNum  int `json:"failInterfaceNum" yaml:"failInterfaceNum"`

	TotalRequestNum int `json:"totalRequestNum"`
	PassRequestNum  int `json:"passRequestNum"`
	FailRequestNum  int `json:"failRequestNum"`

	TotalAssertionNum int `json:"totalAssertionNum"`
	PassAssertionNum  int `json:"passAssertionNum"`
	FailAssertionNum  int `json:"failAssertionNum"`

	InterfaceStatusMap map[uint]map[consts.ResultStatus]int `gorm:"-"`
}

type PlanNormalData struct {
	PlanId             uint   `json:"planId"`
	PlanName           string `json:"planName"`
	ExecEnv            string `json:"execEnv"`
	TotalScenarioNum   int    `json:"totalScenarioNum"`
	TotalInterfaceNum  int    `json:"totalInterfaceNum"`
	TotalAssertionNum  int    `json:"totalAssertionNum"`
	PassAssertionNum   int    `json:"passAssertionNum"`
	FailAssertionNum   int    `json:"failAssertionNum"`
	FailScenarioNum    int    `json:"failAssertionNum"`
	PassScenarioNum    int    `json:"passScenarioNum"`
	TotalProcessorNum  int    `json:"totalProcessorNum"`
	FinishProcessorNum int    `json:"finishProcessorNum"`
}
