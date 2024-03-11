package domain

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type DebugInfo struct {
	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"` // EndpointInterface without DebugInterface init
	ScenarioProcessorId uint `json:"scenarioProcessorId"` // used to load vars by scenario processor
	DiagnoseInterfaceId uint `json:"diagnoseInterfaceId"` // load by interface diagnose
	CaseInterfaceId     uint `json:"caseInterfaceId"`     // load by endpoint case
	IsForBenchmarkCase  bool `json:"isForBenchmarkCase"`  // load by endpoint case

	UsedBy        consts.UsedBy `json:"usedBy"`
	UserId        uint          `json:"userId"`        // used by loading debugData for display
	EnvironmentId uint          `json:"environmentId"` // used by loading debugData for exec
	FromDefine    bool          `json:"fromDefine"`
	ProjectId     int           `json:"projectId"`
}

type SubmitDebugResultRequest struct {
	ResultStatus consts.ResultStatus
	Request      DebugData     `json:"request"`
	Response     DebugResponse `json:"response"`

	PreConditions  []InterfaceExecCondition `json:"preConditions"`
	PostConditions []InterfaceExecCondition `json:"postConditions"`
}

type DebugData struct {
	BaseRequest

	Name string `json:"name"`

	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
	CaseInterfaceId     uint `json:"caseInterfaceId"`
	DiagnoseInterfaceId uint `json:"diagnoseInterfaceId"`

	ScenarioProcessorId uint `json:"scenarioProcessorId"`
	EnvironmentId       uint `json:"environmentId"`

	UsedBy consts.UsedBy `json:"usedBy"`

	ServeId   uint `json:"serveId"`
	ServerId  uint `json:"serverId"`
	ProjectId uint `json:"projectId"`

	BaseUrl string `json:"baseUrl"`

	// used for selection and show in right environment tab
	EnvDataToView *EnvDataToView `json:"envDataToView,omitempty"`

	ProcessorInterfaceSrc consts.ProcessorInterfaceSrc `json:"processorInterfaceSrc"`
	ResponseDefine        Condition                    `json:"responseDefine"`
}

type EnvDataToView struct {
	ShareVars  []GlobalVar `json:"shareVars"`
	EnvVars    []GlobalVar `json:"envVars"`
	GlobalVars []GlobalVar `json:"globalVars"`
	//GlobalParams []GlobalParam `json:"globalParams"`
}

type Condition struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`

	EntityType string      `json:"entityType"`
	EntityId   uint        `json:"entityId"`
	EntityData interface{} `json:"entityData"`

	Ordr     int  `json:"ordr"`
	Disabled bool `json:"disabled"`
}

type InterfaceExecCondition struct {
	Type consts.ConditionType `json:"type"`
	Desc string               `json:"desc"`
	Raw  json.RawMessage      `json:"raw"`
}
