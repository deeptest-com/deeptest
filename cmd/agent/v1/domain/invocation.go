package agentDomain

import (
	v1 "github.com/deeptest-com/deeptest/internal/pkg/domain"
)

type InvokeRequest struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Data v1.DebugData `json:"data"`
}

type InvokeObject struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
	Id        int    `json:"id"`
	UsedBy    string `json:"usedBy"`
}
