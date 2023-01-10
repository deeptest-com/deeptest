package model

type ProcessorInterface struct {
	BaseModel
	InterfaceBase

	Children []*ProcessorInterface `gorm:"-" json:"children"`

	Params  []ProcessorInterfaceParam  `gorm:"-" json:"params"`
	Headers []ProcessorInterfaceHeader `gorm:"-" json:"headers"`

	BodyFormData       []ProcessorInterfaceBodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []ProcessorInterfaceBodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`

	BasicAuth   ProcessorInterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken ProcessorInterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     ProcessorInterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      ProcessorInterfaceApiKey      `gorm:"-" json:"apiKey"`

	InterfaceExtractors  []ProcessorInterfaceExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []ProcessorInterfaceCheckpoint `gorm:"-" json:"interfaceCheckpoints"`
}

func (ProcessorInterface) TableName() string {
	return "biz_processor_interface"
}

type ProcessorInterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (ProcessorInterfaceParam) TableName() string {
	return "biz_processor_interface_param"
}

type ProcessorInterfaceBodyFormDataItem struct {
	BaseModel
	InterfaceBodyFormDataItemBase
}

func (ProcessorInterfaceBodyFormDataItem) TableName() string {
	return "biz_processor_interface_form_data_item"
}

type ProcessorInterfaceBodyFormUrlEncodedItem struct {
	BaseModel
	InterfaceBodyFormUrlEncodedItemBase
}

func (ProcessorInterfaceBodyFormUrlEncodedItem) TableName() string {
	return "biz_processor_interface_form_urlencoded_item"
}

type ProcessorInterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (ProcessorInterfaceHeader) TableName() string {
	return "biz_processor_interface_header"
}

type ProcessorInterfaceBasicAuth struct {
	BaseModel
	InterfaceBasicAuthBase
}

func (ProcessorInterfaceBasicAuth) TableName() string {
	return "biz_processor_interface_basic_auth"
}

type ProcessorInterfaceBearerToken struct {
	BaseModel
	InterfaceBearerTokenBase
}

func (ProcessorInterfaceBearerToken) TableName() string {
	return "biz_processor_interface_bearer_token"
}

type ProcessorInterfaceOAuth20 struct {
	BaseModel
	InterfaceOAuth20Base
}

func (ProcessorInterfaceOAuth20) TableName() string {
	return "biz_processor_interface_oauth20"
}

type ProcessorInterfaceApiKey struct {
	BaseModel
	InterfaceApiKeyBase
}

func (ProcessorInterfaceApiKey) TableName() string {
	return "biz_processor_interface_apikey"
}

type ProcessorInterfaceExtractor struct {
	BaseModel
	InterfaceExtractorBase
}

func (ProcessorInterfaceExtractor) TableName() string {
	return "biz_processor_interface_extractor"
}

type ProcessorInterfaceCheckpoint struct {
	BaseModel
	InterfaceCheckpointBase
}

func (ProcessorInterfaceCheckpoint) TableName() string {
	return "biz_processor_interface_checkpoint"
}
