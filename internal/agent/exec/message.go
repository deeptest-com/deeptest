package agentExec

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type MessageExecReq struct {
	UserId    uint            `json:"userId"`
	ServerUrl string          `json:"serverUrl"`
	Token     string          `json:"token"`
	TenantId  consts.TenantId `json:"tenantId"`
}

type MessageExecObj struct {
	Count uint `json:"count"`
}
