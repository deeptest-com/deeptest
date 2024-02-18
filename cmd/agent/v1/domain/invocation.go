package agentDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/kataras/iris/v12"
)

type InterfaceCall struct {
	ExecUuid  string `json:"execUuid"`
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Data           domain.DebugData `json:"data"`
	LocalVarsCache iris.Map         `json:"localVarsCache"`
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
