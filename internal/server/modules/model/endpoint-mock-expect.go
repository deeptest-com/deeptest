package model

import "github.com/deeptest-com/deeptest/internal/pkg/consts"

type EndpointMockExpect struct {
	BaseModel
	Disabled            bool                               `json:"disabled" gorm:"default:false"`
	Name                string                             `json:"name"`
	EndpointId          uint                               `json:"endpointId"`
	EndpointInterfaceId uint                               `json:"endpointInterfaceId"`
	Method              consts.HttpMethod                  `json:"method"`
	Ordr                int                                `json:"ordr"`
	CreateUser          string                             `json:"createUser"`
	UpdateUser          string                             `json:"updateUser"`
	RequestHeaders      []EndpointMockExpectRequest        `gorm:"-" json:"requestHeaders"`
	RequestBodies       []EndpointMockExpectRequest        `gorm:"-" json:"requestBodies"`
	RequestQueryParams  []EndpointMockExpectRequest        `gorm:"-" json:"requestQueryParams"`
	RequestPathParams   []EndpointMockExpectRequest        `gorm:"-" json:"requestPathParams"`
	ResponseBody        EndpointMockExpectResponse         `gorm:"-" json:"responseBody"`
	ResponseHeaders     []EndpointMockExpectResponseHeader `gorm:"-" json:"responseHeaders"`
}

func (EndpointMockExpect) TableName() string {
	return "biz_endpoint_mock_expect"
}

type EndpointMockExpectRequest struct {
	BaseModel
	EndpointMockExpectId uint                           `json:"endpointMockExpectId"`
	CompareWay           consts.ComparisonOperator      `json:"compareWay"`
	Name                 string                         `json:"name"`
	Value                string                         `json:"value"`
	Source               consts.ParamIn                 `json:"source"`
	SelectType           consts.ExpectRequestSelectType `json:"selectType"`
}

func (EndpointMockExpectRequest) TableName() string {
	return "biz_endpoint_mock_expect_request"
}

type EndpointMockExpectResponse struct {
	BaseModel
	EndpointMockExpectId uint   `json:"endpointMockExpectId"`
	Code                 string `json:"code"`
	DelayTime            uint   `json:"delayTime"`
	Value                string `gorm:"type:text" json:"value"`
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
