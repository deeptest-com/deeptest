package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Report struct {
	BaseModel

	Name string `json:"name" yaml:"name"`
	Desc string `json:"desc" yaml:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus" yaml:"progressStatus"`
	ReportStatus   consts.ResultStatus   `json:"resultStatus" yaml:"resultStatus"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Duration  int        `json:"duration" yaml:"duration"` // sec

	TotalNum  int `json:"totalNum" yaml:"totalNum"`
	PassNum   int `json:"passNum" yaml:"passNum"`
	FailNum   int `json:"failNum" yaml:"failNum"`
	MissedNum int `json:"missedNum" yaml:"missedNum"`

	Payload string `json:"payload"`

	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	Logs []*Log `gorm:"-" json:"logs"`
}

func (Report) TableName() string {
	return "biz_report"
}
