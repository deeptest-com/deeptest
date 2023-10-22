package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DebugInvoke struct {
	BaseModel
	InvocationBase

	ServeId uint `json:"serveId"`

	ScenarioProcessorId uint `gorm:"default:0" json:"scenarioProcessorId"`
	ScenarioId          uint `gorm:"default:0" json:"scenarioId"`
}

func (DebugInvoke) TableName() string {
	return "biz_debug_invoke"
}

type InvocationBase struct {
	Name string `json:"name"`
	Desc string `json:"desc,omitempty" gorm:"type:text"`

	ReqContent  string `json:"reqContent,omitempty" gorm:"type:longtext"`
	RespContent string `json:"respContent,omitempty" gorm:"type:longtext"`

	PreConditionsContent  string `json:"preConditionsContent,omitempty" gorm:"type:longtext"`
	PostConditionsContent string `json:"postConditionsContent,omitempty" gorm:"type:longtext"`

	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	ResultStatus     consts.ResultStatus `json:"resultStatus" gorm:"default:pass"`
	CheckpointStatus consts.ResultStatus `json:"checkpointStatus"`

	EndpointInterfaceId uint `gorm:"index:,default:0" json:"endpointInterfaceId,omitempty"`
	DebugInterfaceId    uint `gorm:"index:,default:0" json:"debugInterfaceId,omitempty"`

	ProjectId uint `json:"projectId,omitempty"`
}
