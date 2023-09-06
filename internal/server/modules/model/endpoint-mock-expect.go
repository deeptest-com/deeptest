package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type EndpointMockExpect struct {
	BaseModel
	Name       string `json:"name"`
	EndpointId uint   `json:"endpointId"`
	Ordr       int    `json:"ordr"`
	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (EndpointMockExpect) TableName() string {
	return "biz_endpoint_mock_expect"
}

type EndpointMockExpectRequest struct {
	BaseModel
	EndpointMockExpectId uint                  `json:"endpointMockExpectId"`
	CompareWay           consts.MockCompareWay `json:"compareWay"`
	Name                 string                `json:"name"`
	Value                string                `json:"value"`
	Source               consts.ParamIn        `json:"source"`
}

func (EndpointMockExpectRequest) TableName() string {
	return "biz_endpoint_mock_expect_request"
}

type EndpointMockExpectResponse struct {
	BaseModel
	EndpointMockExpectId uint   `json:"endpointMockExpectId"`
	Code                 string `json:"code"`
	DelayTime            uint   `json:"delayTime"`
	Value                string `json:"value"`
}

func (EndpointMockExpectResponse) TableName() string {
	return "biz_endpoint_mock_expect_response"
}

type EndpointMockExpectResponseHeader struct {
	BaseModel
	EndpointMockExpectId uint   `json:"endpointMockExpectId"`
	Name                 string `json:"name"`
	Value                string `json:"value"`
}

func (EndpointMockExpectResponseHeader) TableName() string {
	return "biz_endpoint_mock_expect_response_header"
}
