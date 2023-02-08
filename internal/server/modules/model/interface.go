package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type Interface struct {
	BaseModel

	InterfaceBase

	//ReqBodySpec InterfaceReqBodySpec `gorm:"-" json:"basicAuth"`

	Children []*Interface `gorm:"-" json:"children"`

	Params      []InterfaceParam     `gorm:"-" json:"params"`
	Headers     []InterfaceHeader    `gorm:"-" json:"headers"`
	Cookies     []InterfaceCookie    `gorm:"-" json:"cookies"`
	RequestBody InterfaceRequestBody `gorm:"-" json:"requestBody"`

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

type InterfaceCookie struct {
	BaseModel
	InterfaceCookieBase
}

func (InterfaceCookie) TableName() string {
	return "biz_interface_cookie"
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

type InterfaceRequestBodyItem struct {
	BaseModel
	InterfaceRequestBodyItemBase
}

func (InterfaceRequestBodyItem) TableName() string {
	return "biz_interface_request_body_item"
}

type InterfaceRequestBody struct {
	BaseModel
	InterfaceRequestBodyBase
	SchemaItem InterfaceRequestBodyItem `gorm:"-" json:"schemaItem"`
}

func (InterfaceRequestBody) TableName() string {
	return "biz_interface_request_body"
}

type InterfaceExtractor struct {
	BaseModel

	UsedBy consts.UsedBy `json:"usedBy"`

	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `json:"expression"`
	//NodeProp       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string                `json:"variable"`
	Scope    consts.ExtractorScope `json:"scope" gorm:"default:private"`
	//DisableShare bool                  `json:"disableShare"`

	Result      string `json:"result" gorm:"type:text"`
	InterfaceId uint   `json:"interfaceId"`
	ProcessorId uint   `json:"processorId"`
	ScenarioId  uint   `json:"scenarioId"`

	ProjectId uint `json:"projectId"`
}

func (InterfaceExtractor) TableName() string {
	return "biz_interface_extractor"
}

type InterfaceCheckpoint struct {
	BaseModel

	UsedBy consts.UsedBy         `json:"usedBy"`
	Type   consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	ActualResult string              `json:"actualResult"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	InterfaceId  uint                `json:"interfaceId"`
	ScenarioId   uint                `json:"scenarioId"`
}

func (InterfaceCheckpoint) TableName() string {
	return "biz_interface_checkpoint"
}
