package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type EndpointMockScript struct {
	BaseModel
	Content    string `json:"content" gorm:"type:longtext"`
	EndpointId uint   `json:"endpointId"`
	//EndpointInterfaceId uint              `json:"endpointInterfaceId"`
	Method     consts.HttpMethod `json:"method"`
	CreateUser string            `json:"createUser"`
	UpdateUser string            `json:"updateUser"`
}

func (EndpointMockScript) TableName() string {
	return "biz_endpoint_mock_script"
}
