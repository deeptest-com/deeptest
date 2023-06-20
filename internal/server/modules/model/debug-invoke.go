package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DebugInvoke struct {
	BaseModel
	InvocationBase

	ServeId uint `json:"serveId"`

	ProcessorId uint `json:"processorId"`
	ScenarioId  uint `json:"scenarioId"`
}

func (DebugInvoke) TableName() string {
	return "biz_debug_invoke"
}

type InvocationBase struct {
	Name string `json:"name"`
	Desc string `json:"desc,omitempty" gorm:"type:text"`

	ReqContent  string `json:"reqContent,omitempty" gorm:"type:mediumtext"`
	RespContent string `json:"respContent,omitempty" gorm:"type:mediumtext"`

	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	ResultStatus     consts.ResultStatus `json:"resultStatus" gorm:"default:pass"`
	CheckpointStatus consts.ResultStatus `json:"checkpointStatus"`

	EndpointInterfaceId uint `json:"endpointInterfaceId,omitempty"`
	DebugInterfaceId    uint `json:"debugInterfaceId,omitempty"`

	ProjectId uint `json:"projectId,omitempty"`
}
