package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type ScenarioInterface struct {
	BaseModel
	InterfaceBase

	EndpointInterfaceId uint `json:"endpointInterfaceId"`
	ServerId            uint `json:"serverId"`

	//ReqBodySpec InterfaceReqBodySpec `gorm:"-" json:"basicAuth"`

	Children []*ScenarioInterface `gorm:"-" json:"children"`

	QueryParams []ScenarioInterfaceParam  `gorm:"-" json:"queryParams"`
	PathParams  []ScenarioInterfaceParam  `gorm:"-" json:"pathParams"`
	Headers     []ScenarioInterfaceHeader `gorm:"-" json:"headers"`
	Cookies     []ScenarioInterfaceCookie `gorm:"-" json:"cookies"`

	BodyFormData       []ScenarioInterfaceBodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []ScenarioInterfaceBodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`

	BasicAuth   ScenarioInterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken ScenarioInterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     ScenarioInterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      ScenarioInterfaceApiKey      `gorm:"-" json:"apiKey"`

	InterfaceExtractors  []ScenarioInterfaceExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []ScenarioInterfaceCheckpoint `gorm:"-" json:"interfaceCheckpoints"`
}

func (ScenarioInterface) TableName() string {
	return "biz_scenario_interface"
}

type ScenarioInterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (ScenarioInterfaceParam) TableName() string {
	return "biz_scenario_interface_param"
}

type ScenarioInterfaceBodyFormDataItem struct {
	BaseModel
	InterfaceBodyFormDataItemBase
}

func (ScenarioInterfaceBodyFormDataItem) TableName() string {
	return "biz_scenario_interface_form_data_item"
}

type ScenarioInterfaceBodyFormUrlEncodedItem struct {
	BaseModel
	InterfaceBodyFormUrlEncodedItemBase
}

func (ScenarioInterfaceBodyFormUrlEncodedItem) TableName() string {
	return "biz_scenario_interface_form_urlencoded_item"
}

type ScenarioInterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (ScenarioInterfaceHeader) TableName() string {
	return "biz_scenario_interface_header"
}

type ScenarioInterfaceCookie struct {
	BaseModel
	InterfaceCookieBase
}

func (ScenarioInterfaceCookie) TableName() string {
	return "biz_scenario_interface_cookie"
}

type ScenarioInterfaceBasicAuth struct {
	BaseModel
	InterfaceBasicAuthBase
}

func (ScenarioInterfaceBasicAuth) TableName() string {
	return "biz_scenario_interface_basic_auth"
}

type ScenarioInterfaceBearerToken struct {
	BaseModel
	InterfaceBearerTokenBase
}

func (ScenarioInterfaceBearerToken) TableName() string {
	return "biz_scenario_interface_bearer_token"
}

type ScenarioInterfaceOAuth20 struct {
	BaseModel
	InterfaceOAuth20Base
}

func (ScenarioInterfaceOAuth20) TableName() string {
	return "biz_scenario_interface_oauth20"
}

type ScenarioInterfaceApiKey struct {
	BaseModel
	InterfaceApiKeyBase
}

func (ScenarioInterfaceApiKey) TableName() string {
	return "biz_scenario_interface_apikey"
}

type ScenarioInterfaceExtractor struct {
	BaseModel

	domain.ExtractorBase

	Scope consts.ExtractorScope `json:"scope" gorm:"default:private"`

	EndpointInterfaceId uint `json:"endpointInterfaceId"`

	ProcessorId uint `json:"processorId"`
	ScenarioId  uint `json:"scenarioId"`

	ProjectId uint `json:"projectId"`
}

func (ScenarioInterfaceExtractor) TableName() string {
	return "biz_scenario_interface_extractor"
}

type ScenarioInterfaceCheckpoint struct {
	BaseModel

	UsedBy consts.UsedBy         `json:"usedBy"`
	Type   consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	ActualResult string              `json:"actualResult"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`

	EndpointInterfaceId uint `json:"endpointInterfaceId"`
	ScenarioId          uint `json:"scenarioId"`
}

func (ScenarioInterfaceCheckpoint) TableName() string {
	return "biz_scenario_interface_checkpoint"
}
