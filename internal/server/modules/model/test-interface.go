package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
)

type TestInterface struct {
	BaseModel

	Title  string                         `json:"title"`
	Desc   string                         `json:"desc"`
	IsLeaf bool                           `json:"isLeaf"`
	Type   serverConsts.TestInterfaceType `json:"type"`

	ParentId  uint `json:"parentId"`
	ServerId  uint `json:"serverId"`
	ServeId   uint `json:"serveId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int              `json:"ordr"`
	Children []*TestInterface `gorm:"-" json:"children"`
	Slots    iris.Map         `gorm:"-" json:"slots"`

	// debug data
	BaseUrl string `json:"baseUrl"`
	InterfaceConfigBase

	QueryParams []TestInterfaceParam  `gorm:"-" json:"queryParams"`
	PathParams  []TestInterfaceParam  `gorm:"-" json:"pathParams"`
	Headers     []TestInterfaceHeader `gorm:"-" json:"headers"`
	Cookies     []TestInterfaceCookie `gorm:"-" json:"cookies"`

	BodyFormData       []TestInterfaceBodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []TestInterfaceBodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`

	BasicAuth   TestInterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken TestInterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     TestInterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      TestInterfaceApiKey      `gorm:"-" json:"apiKey"`

	InterfaceExtractors  []TestInterfaceExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []TestInterfaceCheckpoint `gorm:"-" json:"interfaceCheckpoints"`
}

func (TestInterface) TableName() string {
	return "biz_test_interface"
}

type TestInterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (TestInterfaceParam) TableName() string {
	return "biz_test_interface_param"
}

type TestInterfaceBodyFormDataItem struct {
	BaseModel
	InterfaceBodyFormDataItemBase
}

func (TestInterfaceBodyFormDataItem) TableName() string {
	return "biz_test_interface_form_data_item"
}

type TestInterfaceBodyFormUrlEncodedItem struct {
	BaseModel
	InterfaceBodyFormUrlEncodedItemBase
}

func (TestInterfaceBodyFormUrlEncodedItem) TableName() string {
	return "biz_test_interface_form_urlencoded_item"
}

type TestInterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (TestInterfaceHeader) TableName() string {
	return "biz_test_interface_header"
}

type TestInterfaceCookie struct {
	BaseModel
	InterfaceCookieBase
}

func (TestInterfaceCookie) TableName() string {
	return "biz_test_interface_cookie"
}

type TestInterfaceBasicAuth struct {
	BaseModel
	InterfaceBasicAuthBase
}

func (TestInterfaceBasicAuth) TableName() string {
	return "biz_test_interface_basic_auth"
}

type TestInterfaceBearerToken struct {
	BaseModel
	InterfaceBearerTokenBase
}

func (TestInterfaceBearerToken) TableName() string {
	return "biz_test_interface_bearer_token"
}

type TestInterfaceOAuth20 struct {
	BaseModel
	InterfaceOAuth20Base
}

func (TestInterfaceOAuth20) TableName() string {
	return "biz_test_interface_oauth20"
}

type TestInterfaceApiKey struct {
	BaseModel
	InterfaceApiKeyBase
}

func (TestInterfaceApiKey) TableName() string {
	return "biz_test_interface_apikey"
}

type TestInterfaceExtractor struct {
	BaseModel

	domain.ExtractorBase

	Scope consts.ExtractorScope `json:"scope" gorm:"default:private"`

	EndpointInterfaceId uint `json:"endpointInterfaceId"`

	ProcessorId uint `json:"processorId"`
	ScenarioId  uint `json:"scenarioId"`

	ProjectId uint `json:"projectId"`
}

func (TestInterfaceExtractor) TableName() string {
	return "biz_test_interface_extractor"
}

type TestInterfaceCheckpoint struct {
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

func (TestInterfaceCheckpoint) TableName() string {
	return "biz_test_interface_checkpoint"
}
