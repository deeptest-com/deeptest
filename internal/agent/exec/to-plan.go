package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type PlanExecReq struct {
	ExecUuid      string          `json:"execUuid"`
	ServerUrl     string          `json:"serverUrl"`
	Token         string          `json:"token"`
	PlanId        uint            `json:"planId"`
	EnvironmentId uint            `json:"environmentId"`
	TenantId      consts.TenantId `json:"tenantId"`
}

type PlanExecObj struct {
	Name      string            `json:"name"`
	Scenarios []ScenarioExecObj `json:"scenarios"`

	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
