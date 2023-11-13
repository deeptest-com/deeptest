package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type EndpointCase struct {
	BaseModel

	Name   string            `json:"name"`
	Desc   string            `json:"desc"`
	Method consts.HttpMethod `json:"method"`

	EndpointId uint `json:"endpointId"`
	ServeId    uint `json:"serveId"`
	ProjectId  uint `json:"projectId"`

	DebugInterfaceId uint            `gorm:"default:0" json:"debugInterfaceId"`
	DebugData        *DebugInterface `gorm:"-" json:"debugData"`

	CaseType consts.CaseType `gorm:"default:'default'" json:"caseType"`
	BaseCase uint            `gorm:"default:0" json:"baseCase"`

	SrcId uint `json:"srcId"`

	SerialNumber   string         `json:"serialNumber"`
	CreateUserId   uint           `json:"createUserId"`
	CreateUserName string         `json:"createUserName"`
	Children       []EndpointCase `gorm:"-" json:"children"`
}

func (EndpointCase) TableName() string {
	return "biz_endpoint_case"
}

type EndpointCaseAlternative struct {
	BaseModel

	BaseId uint `json:"baseId"`

	Type consts.AlternativeCaseType `json:"type"`
	Path string                     `json:"path"`
}

func (EndpointCaseAlternative) TableName() string {
	return "biz_endpoint_case_alternative"
}

type EndpointCaseAlternativeFactor struct {
	BaseModel

	CaseId uint `json:"caseId"`

	Value string `json:"value"`
	Path  string `json:"path"`
}

func (EndpointCaseAlternativeFactor) TableName() string {
	return "biz_endpoint_case_alternative_factor"
}
