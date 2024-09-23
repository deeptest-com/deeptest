package agentExec

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type InterfaceExecReq struct {
	ExecUuid  string          `json:"execUuid"`
	ServerUrl string          `json:"serverUrl"`
	Token     string          `json:"token"`
	TenantId  consts.TenantId `json:"tenantId"`

	Data      domain.DebugData `json:"data"`
	ExecScene domain.ExecScene `json:"execScene"`
}

type InterfaceExecObj struct {
	DebugData domain.DebugData `json:"debugData"`

	PreConditions  []domain.InterfaceExecCondition `json:"preConditions"`
	PostConditions []domain.InterfaceExecCondition `json:"postConditions"`

	ExecScene domain.ExecScene `json:"execScene"`
	TenantId  consts.TenantId  `json:"tenantId"`
}
