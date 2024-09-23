package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type WsReq struct {
	Act consts.ExecType `json:"act"`

	InterfaceExecReq InterfaceExecReq `json:"interfaceExecReq"`
	ScenarioExecReq  ScenarioExecReq  `json:"scenarioExecReq"`
	PlanExecReq      PlanExecReq      `json:"planExecReq"`
	CasesExecReq     CasesExecReq     `json:"casesExecReq"`

	WebsocketExecReq WebsocketExecReq `json:"websocketExecReq"`

	MessageReq     MessageExecReq  `json:"messageReq"`
	LocalVarsCache iris.Map        `json:"localVarsCache"`
	TenantId       consts.TenantId `json:"tenantId"`
}
