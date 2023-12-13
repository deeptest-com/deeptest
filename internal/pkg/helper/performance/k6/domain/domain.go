package keDomain

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	k6Comm "github.com/aaronchen2k/deeptest/internal/pkg/helper/performance/k6/comm"
)

type PerfPlan struct {
	Name      string         `json:"name"`
	Scenarios []PerfScenario `json:"scenarios"`
}

type PerfScenario struct {
	Name     string               `json:"name"`
	Func     string               `json:"func"`
	Executor k6Comm.ExecutorsType `json:"executor"`
	Vues     int                  `json:"vues"`
	Duration int                  `json:"duration"`

	CodesGenerated string `json:"codesGenerated"`

	RootProcessor agentExec.Processor `json:"rootProcessor"`
}
