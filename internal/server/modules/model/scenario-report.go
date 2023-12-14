package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type ScenarioReport struct {
	BaseModel

	Name string `json:"name"`
	Desc string `gorm:"type:text" json:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" gorm:"default:pass"`

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

	TotalProcessorNum  int `json:"totalProcessorNum"`
	FinishProcessorNum int `json:"finishProcessorNum"`

	InterfaceStatusMap map[uint]map[consts.ResultStatus]int `gorm:"-"`

	Payload string `json:"payload"`

	ScenarioId   uint `json:"scenarioId"`
	ProjectId    uint `json:"projectId"`
	PlanReportId uint `json:"planReportId"`

	CreateUserId   uint                `json:"createUserId"`
	CreateUserName string              `gorm:"-" json:"createUserName"`
	ExecUserName   string              `gorm:"-" json:"execUserName"`
	SerialNumber   string              `json:"serialNumber"`
	Logs           []*ExecLogProcessor `gorm:"-" json:"logs"`

	ExecEnv   string `gorm:"-" json:"execEnv"`
	ExecEnvId int    `json:"execEnvId"`
	Priority  string `gorm:"-" json:"priority"`

	StatRaw string `json:"stat"`
}

func (ScenarioReport) TableName() string {
	return "biz_scenario_report"
}

type ScenarioReportDetail struct {
	ScenarioReport
	Priority string `json:"priority"`
}
