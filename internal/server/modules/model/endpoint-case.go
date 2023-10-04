package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type EndpointCase struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	EndpointId uint `json:"endpointId"`
	ServeId    uint `json:"serveId"`
	ProjectId  uint `json:"projectId"`

	DebugInterfaceId uint            `gorm:"default:0" json:"debugInterfaceId"`
	DebugData        *DebugInterface `gorm:"-" json:"debugData"`

	SrcId uint `json:"srcId"`

	SerialNumber   string `json:"serialNumber"`
	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
}

func (EndpointCase) TableName() string {
	return "biz_endpoint_case"
}

type EndpointCaseAlternative struct {
	BaseModel

	BaseId uint `json:"baseId"`

	Type      consts.AlternativeCaseType `json:"type"`
	ParamName string                     `json:"paramName"`
	FormField string                     `json:"formField"`
	BodyPath  string                     `json:"bodyPath"`
}

func (EndpointCaseAlternative) TableName() string {
	return "biz_endpoint_case_alternative"
}
