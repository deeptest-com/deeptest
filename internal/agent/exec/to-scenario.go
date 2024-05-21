package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type ScenarioExecReq struct {
	ExecUuid   string          `json:"execUuid"`
	ServerUrl  string          `json:"serverUrl"`
	Token      string          `json:"token"`
	TenantId   consts.TenantId `json:"tenantId"`
	ScenarioId uint            `json:"scenarioId"`

	EnvironmentId uint `json:"environmentId"`
}

type ScenarioExecObj struct {
	ScenarioExecObjBase
	RootProcessor *Processor `json:"rootProcessor"`
}

type ScenarioExecObjMsg struct {
	ScenarioExecObjBase
	RootProcessor *ProcessorMsg `json:"rootProcessor"`
}

type ScenarioExecObjBase struct {
	ScenarioId uint   `json:"scenarioId"`
	Name       string `json:"name"`

	BaseUrl string `json:"baseUrl"`

	ExecScene domain.ExecScene `json:"execScene"`

	ExecUuid  string          `json:"execUuid"`
	ServerUrl string          `json:"serverUrl"`
	Token     string          `json:"token"`
	TenantId  consts.TenantId `json:"tenantId"`
}

type WebsocketExecReq struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Room string                    `json:"room"`
	Data domain.WebsocketDebugData `json:"data"`

	WebsocketInterfaceId int             `json:"websocketInterfaceId"`
	EnvironmentId        int             `json:"environmentId"`
	TenantId             consts.TenantId `json:"tenantId"`
}
