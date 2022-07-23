package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type TestResult struct {
	BaseModel

	Name string `json:"name" yaml:"name"`
	Desc string `json:"desc" yaml:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus" yaml:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" yaml:"resultStatus"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`

	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	Logs []TestLog `gorm:"-" json:"logs"`
}

func (TestResult) TableName() string {
	return "biz_test_result"
}
