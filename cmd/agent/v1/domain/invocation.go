package domain

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type InvokeCall struct {
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`

	Data v1.DebugCall `json:"data"`
}

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
