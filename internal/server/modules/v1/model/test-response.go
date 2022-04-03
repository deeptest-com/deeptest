package model

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
)

type TestResponse struct {
	BaseModel

	serverDomain.TestRequestResp

	InterfaceId uint `json:"interfaceId"`
	ProjectId   uint `json:"projectId"`
}

func (TestResponse) TableName() string {
	return "biz_test_response"
}
