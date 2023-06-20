package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DebugReq struct {
	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"` // EndpointInterface without DebugInterface init
	ScenarioProcessorId uint `json:"scenarioProcessorId"` // used to load vars by scenario processor
	TestInterfaceId     uint `json:"testInterfaceId"`     // load by interface testing

	UsedBy consts.UsedBy `json:"usedBy"`
}

type SubmitDebugResultRequest struct {
	Request  DebugData     `json:"request"`
	Response DebugResponse `json:"response"`
}

type DebugData struct {
	BaseRequest

	DebugInterfaceId uint `json:"debugInterfaceId"`

	EndpointInterfaceId uint          `json:"endpointInterfaceId"`
	ScenarioProcessorId uint          `json:"scenarioProcessorId"`
	TestInterfaceId     uint          `json:"testInterfaceId"`
	UsedBy              consts.UsedBy `json:"usedBy"`

	ServeId  uint `json:"serveId"`
	ServerId uint `json:"serverId"`

	BaseUrl   string      `json:"baseUrl"`
	ShareVars []GlobalVar `json:"shareVars"` // used to show in right environment tab
	EnvVars   []GlobalVar `json:"envVars"`   // used to show in right environment tab

	Name string `json:"name"`
}
