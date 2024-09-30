package domain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
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

	Payload interface{} `json:"payload"`

	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	Logs []ExecLog `gorm:"-" json:"logs"`
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
