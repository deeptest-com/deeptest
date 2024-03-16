package agentDomain

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	ptdomain "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type WsReq struct {
	Act consts.ExecType `json:"act"`

	ScenarioExecReq agentExec.ScenarioExecReq `json:"scenarioExecReq"`
	PlanExecReq     agentExec.PlanExecReq     `json:"planExecReq"`
	CasesExecReq    agentExec.CasesExecReq    `json:"casesExecReq"`

	MessageReq     agentExec.MessageExecReq `json:"messageReq"`
	LocalVarsCache iris.Map                 `json:"localVarsCache"`

	PerformanceTestExecReq ptdomain.PerformanceTestReq `json:"performanceTestExecReq"`
	PerformanceLogExecReq  ptdomain.PerformanceLogReq  `json:"performanceLogExecReq"`
}
