package model

type Interface struct {
	BaseModel

	InterfaceBase

	Children []*Interface `gorm:"-" json:"children"`

	Params  []InterfaceParam  `gorm:"-" json:"params"`
	Headers []InterfaceHeader `gorm:"-" json:"headers"`

	BodyFormData       []InterfaceBodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []InterfaceBodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`

	BasicAuth   InterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken InterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     InterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      InterfaceApiKey      `gorm:"-" json:"apiKey"`

	InterfaceExtractors  []InterfaceExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []InterfaceCheckpoint `gorm:"-" json:"interfaceCheckpoints"`
}

func (Interface) TableName() string {
	return "biz_interface"
}

type InterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (InterfaceParam) TableName() string {
	return "biz_interface_param"
}

type InterfaceBodyFormDataItem struct {
	BaseModel
	InterfaceBodyFormDataItemBase
}

func (InterfaceBodyFormDataItem) TableName() string {
	return "biz_interface_form_data_item"
}

type InterfaceBodyFormUrlEncodedItem struct {
	BaseModel
	InterfaceBodyFormUrlEncodedItemBase
}

func (InterfaceBodyFormUrlEncodedItem) TableName() string {
	return "biz_interface_form_urlencoded_item"
}

type InterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (InterfaceHeader) TableName() string {
	return "biz_interface_header"
}

type InterfaceBasicAuth struct {
	BaseModel
	InterfaceBasicAuthBase
}

func (InterfaceBasicAuth) TableName() string {
	return "biz_interface_basic_auth"
}

type InterfaceBearerToken struct {
	BaseModel
	InterfaceBearerTokenBase
}

func (InterfaceBearerToken) TableName() string {
	return "biz_interface_bearer_token"
}

type InterfaceOAuth20 struct {
	BaseModel

	InterfaceOAuth20Base
}

func (InterfaceOAuth20) TableName() string {
	return "biz_interface_oauth20"
}

type InterfaceApiKey struct {
	BaseModel
	InterfaceApiKeyBase
}

func (InterfaceApiKey) TableName() string {
	return "biz_interface_apikey"
}

type InterfaceExtractor struct {
	BaseModel

	InterfaceExtractorBase
}

func (InterfaceExtractor) TableName() string {
	return "biz_interface_extractor"
}

type InterfaceCheckpoint struct {
	BaseModel

	InterfaceCheckpointBase
}

func (InterfaceCheckpoint) TableName() string {
	return "biz_interface_checkpoint"
}
