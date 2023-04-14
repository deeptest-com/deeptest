package domain

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type InvocationReq struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Data v1.DebugRequest `json:"data"`
}

type InvocationObject struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
	Id        int    `json:"id"`
	UsedBy    string `json:"usedBy"`
}
