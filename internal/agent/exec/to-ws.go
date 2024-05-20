package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type WsReq struct {
	Act consts.ExecType `json:"act"`

	ScenarioExecReq ScenarioExecReq `json:"scenarioExecReq"`
	PlanExecReq     PlanExecReq     `json:"planExecReq"`
	CasesExecReq    CasesExecReq    `json:"casesExecReq"`

	MessageReq     MessageExecReq  `json:"messageReq"`
	LocalVarsCache iris.Map        `json:"localVarsCache"`
	TenantId       consts.TenantId `json:"tenantId"`
}
