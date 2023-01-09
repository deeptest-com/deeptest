package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type InvocationBase struct {
	Name string `json:"name"`
	Desc string `json:"desc,omitempty" gorm:"type:text"`

	ReqContent  string `json:"reqContent,omitempty" gorm:"type:mediumtext"`
	RespContent string `json:"respContent,omitempty" gorm:"type:mediumtext"`

	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	ResultStatus     consts.ResultStatus `json:"resultStatus" gorm:"default:pass"`
	CheckpointStatus consts.ResultStatus `json:"checkpointStatus"`

	InterfaceId uint `json:"interfaceId,omitempty"`
	ProjectId   uint `json:"projectId,omitempty"`
}

type Invocation struct {
	BaseModel
	InvocationBase
}

func (Invocation) TableName() string {
	return "biz_invocation"
}
