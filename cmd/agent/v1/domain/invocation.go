package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/kataras/iris/v12"
)

type InterfaceCall struct {
	ExecUuid  string          `json:"execUuid"`
	ServerUrl string          `json:"serverUrl"`
	Token     string          `json:"token"`
	TenantId  consts.TenantId `json:"tenantId"`

	LocalVarsCache iris.Map `json:"localVarsCache"`

	Data      domain.DebugData `json:"data"`
	ExecScene domain.ExecScene `json:"execScene"`
}

type InvokeRequest struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Data domain.DebugData `json:"data"`
}

type InvokeObject struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
	Id        int    `json:"id"`
	UsedBy    string `json:"usedBy"`
}
