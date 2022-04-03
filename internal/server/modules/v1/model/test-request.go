package model

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
)

type TestRequest struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	serverDomain.TestRequest

	InterfaceId uint `json:"interfaceId"`
	ProjectId   uint `json:"projectId"`
}

func (TestRequest) TableName() string {
	return "biz_test_request"
}
