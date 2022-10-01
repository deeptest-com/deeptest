package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type Invocation struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`

	ReqContent  string `json:"reqContent,omitempty"`
	RespContent string `json:"respContent,omitempty"`

	HttpRespStatusCode    consts.HttpRespCode `json:"httpStatusCode"`
	HttpRespStatusContent string              `json:"httpStatusContent"`

	ResultStatus     consts.ResultStatus `json:"resultStatus" gorm:"default:pass"`
	CheckpointStatus consts.ResultStatus `json:"checkpointStatus"`

	InterfaceId uint `json:"interfaceId,omitempty"`
	ProjectId   uint `json:"projectId,omitempty"`
}

func (Invocation) TableName() string {
	return "biz_invocation"
}
