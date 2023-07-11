package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DebugReq struct {
	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"` // EndpointInterface without DebugInterface init
	CaseInterfaceId     uint `json:"caseInterfaceId"`     // load by endpoint case
	ScenarioProcessorId uint `json:"scenarioProcessorId"` // used to load vars by scenario processor
	DiagnoseInterfaceId uint `json:"diagnoseInterfaceId"` // load by interface diagnose

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
	DiagnoseInterfaceId uint          `json:"diagnoseInterfaceId"`
	CaseInterfaceId     uint          `json:"caseInterfaceId"`
	UsedBy              consts.UsedBy `json:"usedBy"`

	ServeId   uint `json:"serveId"`
	ServerId  uint `json:"serverId"`
	ProjectId uint `json:"projectId"`

	BaseUrl string `json:"baseUrl"`

	// used for selection and show in right environment tab
	ShareVars    []GlobalVar   `json:"shareVars"`
	EnvVars      []GlobalVar   `json:"envVars"`
	GlobalVars   []GlobalVar   `json:"globalVars"`
	GlobalParams []GlobalParam `json:"globalParams"`

	Name string `json:"name"`
}
