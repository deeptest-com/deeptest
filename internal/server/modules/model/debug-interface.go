package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type DebugInterface struct {
	BaseModel
	InterfaceBase

	EndpointInterfaceId uint `gorm:"default:0" json:"endpointInterfaceId"`
	CaseInterfaceId     uint `gorm:"default:0" json:"caseInterfaceId"`
	DiagnoseInterfaceId uint `gorm:"default:0" json:"diagnoseInterfaceId"`

	ScenarioProcessorId   uint                         `gorm:"default:0" json:"scenarioProcessorId"`
	ProcessorInterfaceSrc consts.ProcessorInterfaceSrc `json:"processorInterfaceSrc"`

	ServeId uint `json:"serveId"`
	// used by DiagnoseInterface
	ServerId uint   `json:"serverId"`
	BaseUrl  string `json:"baseUrl"`

	// debug data
	QueryParams []DebugInterfaceParam  `gorm:"-" json:"queryParams"`
	PathParams  []DebugInterfaceParam  `gorm:"-" json:"pathParams"`
	Headers     []DebugInterfaceHeader `gorm:"-" json:"headers"`
	Cookies     []DebugInterfaceCookie `gorm:"-" json:"cookies"`

	BodyFormData       []DebugInterfaceBodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []DebugInterfaceBodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`

	BasicAuth   DebugInterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken DebugInterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     DebugInterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      DebugInterfaceApiKey      `gorm:"-" json:"apiKey"`

	InterfaceExtractors  []DebugConditionExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []DebugConditionCheckpoint `gorm:"-" json:"interfaceCheckpoints"`

	GlobalParams []DebugInterfaceGlobalParam `gorm:"-" json:"globalParams"`
}

func (DebugInterface) TableName() string {
	return "biz_debug_interface"
}

type DebugInterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (DebugInterfaceParam) TableName() string {
	return "biz_debug_interface_param"
}

type DebugInterfaceBodyFormDataItem struct {
	BaseModel
	InterfaceBodyFormDataItemBase
}

func (DebugInterfaceBodyFormDataItem) TableName() string {
	return "biz_debug_interface_form_data_item"
}

type DebugInterfaceBodyFormUrlEncodedItem struct {
	BaseModel
	InterfaceBodyFormUrlEncodedItemBase
}

func (DebugInterfaceBodyFormUrlEncodedItem) TableName() string {
	return "biz_debug_interface_form_urlencoded_item"
}

type DebugInterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (DebugInterfaceHeader) TableName() string {
	return "biz_debug_interface_header"
}

type DebugInterfaceCookie struct {
	BaseModel
	InterfaceCookieBase
}

func (DebugInterfaceCookie) TableName() string {
	return "biz_debug_interface_cookie"
}

type DebugInterfaceBasicAuth struct {
	BaseModel
	InterfaceBasicAuthBase
}

func (DebugInterfaceBasicAuth) TableName() string {
	return "biz_debug_interface_basic_auth"
}

type DebugInterfaceBearerToken struct {
	BaseModel
	InterfaceBearerTokenBase
}

func (DebugInterfaceBearerToken) TableName() string {
	return "biz_debug_interface_bearer_token"
}

type DebugInterfaceOAuth20 struct {
	BaseModel

	InterfaceOAuth20Base
}

func (DebugInterfaceOAuth20) TableName() string {
	return "biz_debug_interface_oauth20"
}

type DebugInterfaceApiKey struct {
	BaseModel
	InterfaceApiKeyBase
}

func (DebugInterfaceApiKey) TableName() string {
	return "biz_debug_interface_apikey"
}

type DebugInterfaceGlobalParam struct {
	domain.GlobalParam
	//BaseModel
	InterfaceId uint `json:"interfaceId"`
}

func (DebugInterfaceGlobalParam) TableName() string {
	return "biz_debug_interface_global_param"
}
