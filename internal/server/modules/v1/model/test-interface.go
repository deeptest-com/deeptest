package model

import "github.com/kataras/iris/v12"

type TestInterface struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	IsDir     bool `json:"isDir"`
	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int              `json:"ordr"`
	Children []*TestInterface `gorm:"-" json:"children"`

	Slots iris.Map `gorm:"-" json:"slots"`

	// config
	Url               string   `json:"url"`
	Method            string   `gorm:"default:GET" json:"method"`
	Params            []Param  `gorm:"-" json:"params"`
	Headers           []Header `gorm:"-" json:"headers"`
	Body              string   `gorm:"default:{}" json:"body"`
	BodyType          string   `gorm:"default:json" json:"bodyType"`
	AuthorizationType string   `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string   `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string   `gorm:"default:''" json:"validationScript"`

	BasicAuth   BasicAuth   `gorm:" -" json:"basicAuth"`
	BearerToken BearerToken `gorm:" -" json:"bearerToken"`
	OAuth20     OAuth20     `gorm:" -" json:"oAuth20"`
	ApiKey      ApiKey      `gorm:" -" json:"apiKey"`
}

type Param struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type Header struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type BearerToken struct {
	Username string `json:"username"`
}
type OAuth20 struct {
	Key              string `json:"key"`              // key
	OidcDiscoveryURL string `json:"oidcDiscoveryURL"` // OpenID Connect Discovery URL
	AuthURL          string `json:"authURL"`          // Authentication URL
	AccessTokenURL   string `json:"accessTokenURL"`   // Access Token URL
	ClientID         string `json:"clientID"`         // Client ID
	Scope            string `json:"scope"`            // Scope
}
type ApiKey struct {
	Username     string `json:"username"`
	Value        string `json:"value"`
	TransferMode string `json:"transferMode"`
}

func (TestInterface) TableName() string {
	return "biz_test_interface"
}
