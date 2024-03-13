package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"time"
)

type Report struct {
	ID uint `json:"id"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	ProgressStatus consts.ProgressStatus `json:"progressStatus"`
	ResultStatus   consts.ResultStatus   `json:"resultStatus"`

	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Duration  int64      `json:"duration"` // sec

	TotalInterfaceNum int `json:"totalInterfaceNum"`
	PassInterfaceNum  int `json:"passInterfaceNum"`
	FailInterfaceNum  int `json:"failInterfaceNum"`

	TotalRequestNum int `json:"totalRequestNum"`
	PassRequestNum  int `json:"passRequestNum"`
	FailRequestNum  int `json:"failRequestNum"`

	TotalAssertionNum int `json:"totalAssertionNum"`
	PassAssertionNum  int `json:"passAssertionNum"`
	FailAssertionNum  int `json:"failAssertionNum"`

	InterfaceStatusMap map[uint]map[consts.ResultStatus]int `json:"interfaceStatusMap"`

	Payload string `json:"payload"`

	ScenarioId uint `json:"scenarioId"`
	ProjectId  uint `json:"projectId"`

	Logs []*agentExecDomain.ScenarioExecResult `json:"logs"`
}
