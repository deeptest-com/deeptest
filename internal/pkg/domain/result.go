package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Result struct {
	ID   int    `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Desc string `json:"desc" yaml:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus" yaml:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus" yaml:"resultStatus"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`

	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	Logs []Log `gorm:"-" json:"logs"`
}
