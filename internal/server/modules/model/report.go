package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Report struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" gorm:"default:pass"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	ExecTime  int        `json:"execTime"`
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

	Payload string `json:"payload"`

	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	Logs []*ExecLogProcessor `gorm:"-" json:"logs"`
}

func (Report) TableName() string {
	return "biz_report"
}
