package model

type EndpointInterface struct {
	BaseModel
	InterfaceBase
	Params         []EndpointInterfaceParam        `gorm:"-" json:"params"`
	Headers        []EndpointInterfaceHeader       `gorm:"-" json:"headers"`
	Cookies        []EndpointInterfaceCookie       `gorm:"-" json:"cookies"`
	RequestBody    EndpointInterfaceRequestBody    `gorm:"-" json:"requestBody"`
	ResponseBodies []EndpointInterfaceResponseBody `gorm:"-" json:"responseBodies"`
}

func (EndpointInterface) TableName() string {
	return "biz_endpoint_interface"
}

type EndpointInterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (EndpointInterfaceParam) TableName() string {
	return "biz_endpoint_interface_param"
}

type EndpointInterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (EndpointInterfaceHeader) TableName() string {
	return "biz_endpoint_interface_header"
}

type EndpointInterfaceCookie struct {
	BaseModel
	InterfaceCookieBase
}

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
	InterfaceResponseBodyHeaderBase
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
