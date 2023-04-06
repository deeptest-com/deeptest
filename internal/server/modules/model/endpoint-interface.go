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
type SchemaParam struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Desc        string `json:"desc"`
	InterfaceId uint   `json:"interfaceId"`
	Format      string `json:"format"`
	Example     string `json:"example"`
	Pattern     string `json:"pattern"`
	MinLength   int64  `json:"minLength"`
	MaxLength   int64  `json:"maxLength"`
	Default     string `json:"default"`
	Required    bool   `json:"required"`
	MultipleOf  int64  `json:"multipleOf"`
	MinItems    int64  `json:"minItems"`
	MaxItems    int64  `json:"maxItems"`
	UniqueItems bool   `json:"uniqueItems"`
	Ref         string `json:"ref"`
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
