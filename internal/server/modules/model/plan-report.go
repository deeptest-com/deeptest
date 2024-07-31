package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type PlanReport struct {
	BaseModel

	Name string `json:"name"`
	Desc string `gorm:"type:text" json:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" gorm:"default:pass"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Duration  int64      `json:"duration"` // sec

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

	Payload string `json:"payload"`

	PlanId    uint `json:"planId"`
	ProjectId uint `json:"projectId"`
	//ReportId  uint `json:"reportId"`

	CreateUserId uint   `json:"createUserId"`
	SerialNumber string `json:"serialNumber"`
	ExecEnvId    uint   `json:"execEnvId"` //执行环境Id

	StatRaw string `json:"stat"`

	//Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (PlanReport) TableName() string {
	return "biz_plan_report"
}

type PlanReportDetail struct {
	PlanReport
	CreateUserName  string                 `json:"createUserName"`
	ExecUserName    string                 `json:"execUserName"`
	ExecEnv         string                 `json:"execEnv"` //执行环境
	ScenarioReports []ScenarioReportDetail `json:"scenarioReports"`
	TestRate        uint                   `json:"testRate"`
	PlanName        string                 `json:"planName"`
}
