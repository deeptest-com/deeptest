package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type EndpointInterface struct {
	BaseModel
	InterfaceBase
	EndpointId uint `json:"endpointId" gorm:"index"`

	Params         []EndpointInterfaceParam        `gorm:"-" json:"params"`
	Headers        []EndpointInterfaceHeader       `gorm:"-" json:"headers"`
	Cookies        []EndpointInterfaceCookie       `gorm:"-" json:"cookies"`
	RequestBody    EndpointInterfaceRequestBody    `gorm:"-" json:"requestBody"`
	ResponseBodies []EndpointInterfaceResponseBody `gorm:"-" json:"responseBodies"`
	ResponseCodes  string                          `json:"responseCodes"`
	Tags           []string                        `gorm:"-" json:"tags"`
	PathParams     []EndpointPathParam             `gorm:"-" json:"pathParams"`

	DebugInterfaceId uint                           `gorm:"default:0" json:"debugInterfaceId"`
	SourceType       consts.SourceType              `json:"sourceType" gorm:"default:0"`
	Creator          string                         `gorm:"-" json:"creator"`
	GlobalParams     []EndpointInterfaceGlobalParam `gorm:"-" json:"globalParams"`
}
type SchemaParam struct {
	Name        string  `json:"name"`
	Value       string  `gorm:"type:text" json:"value"`
	Type        string  `json:"type"`
	Desc        string  `json:"desc"`
	InterfaceId uint    `json:"interfaceId" gorm:"index"`
	Format      string  `json:"format"`
	Example     string  `gorm:"type:text" json:"example"`
	Pattern     string  `json:"pattern"`
	MinLength   uint64  `json:"minLength"`
	MaxLength   uint64  `json:"maxLength"`
	Default     string  `gorm:"type:text" json:"default"`
	Required    bool    `json:"required"`
	MultipleOf  float64 `json:"multipleOf"`
	MinItems    uint64  `json:"minItems"`
	MaxItems    uint64  `json:"maxItems"`
	UniqueItems bool    `json:"uniqueItems"`
	Ref         string  `json:"ref"`
	Description string  `gorm:"type:text" json:"description"`
	Minimum     float64 `json:"minimum"`
	Maximum     float64 `json:"maximum"`
	IsGlobal    bool    `gorm:"-" json:"isGlobal"`
}

func (EndpointInterface) TableName() string {
	return "biz_endpoint_interface"
}

type EndpointInterfaceParam struct {
	BaseModel
	SchemaParam
}

func (EndpointInterfaceParam) TableName() string {
	return "biz_endpoint_interface_param"
}

type EndpointInterfaceHeader EndpointInterfaceParam

func (EndpointInterfaceHeader) TableName() string {
	return "biz_endpoint_interface_header"
}

type EndpointInterfaceCookie EndpointInterfaceParam

func (EndpointInterfaceCookie) TableName() string {
	return "biz_endpoint_interface_cookie"
}

type EndpointInterfaceRequestBodyItem struct {
	BaseModel
	InterfaceRequestBodyItemBase
}

func (EndpointInterfaceRequestBodyItem) TableName() string {
	return "biz_endpoint_interface_request_body_item"
}

type EndpointInterfaceResponseBodyItem struct {
	BaseModel
	InterfaceResponseBodyItemBase
}

func (EndpointInterfaceResponseBodyItem) TableName() string {
	return "biz_endpoint_interface_response_body_item"
}

type EndpointInterfaceResponseBodyHeader struct {
	BaseModel
	SchemaParam
	ResponseBodyId uint `json:"responseBodyId"`
}

func (EndpointInterfaceResponseBodyHeader) TableName() string {
	return "biz_endpoint_interface_response_body_header"
}

type EndpointInterfaceRequestBody struct {
	BaseModel
	InterfaceRequestBodyBase

	SchemaItem EndpointInterfaceRequestBodyItem `gorm:"-" json:"schemaItem"`
}

func (EndpointInterfaceRequestBody) TableName() string {
	return "biz_endpoint_interface_request_body"
}

type EndpointInterfaceResponseBody struct {
	BaseModel
	InterfaceResponseBodyBase
	SchemaItem EndpointInterfaceResponseBodyItem     `gorm:"-" json:"schemaItem"`
	Headers    []EndpointInterfaceResponseBodyHeader `gorm:"-" json:"headers"`
}

func (EndpointInterfaceResponseBody) TableName() string {
	return "biz_endpoint_interface_response_body"
}

type EndpointInterfaceGlobalParam struct {
	domain.GlobalParam
	InterfaceId uint `json:"InterfaceId"`
}

func (EndpointInterfaceGlobalParam) TableName() string {
	return "biz_endpoint_interface_global_param"
}
