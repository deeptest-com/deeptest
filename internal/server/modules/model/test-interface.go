package model

import (
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

	DebugInterfaceId uint            `json:"debugInterfaceId"`
	DebugInterface   *DebugInterface `gorm:"-" json:"debugInterface"`

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

func (TestInterface) TableName() string {
	return "biz_test_interface"
}
