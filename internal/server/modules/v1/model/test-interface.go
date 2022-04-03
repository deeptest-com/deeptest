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
	Headers           []Header `gorm:"-" json:"heads"`
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
	name     string
	value    string
	disabled bool
}

type Header struct {
	name     string
	value    string
	disabled bool
}

type BasicAuth struct {
	username string
	password string
}
type BearerToken struct {
	username string
}
type OAuth20 struct {
	key              string // key
	oidcDiscoveryURL string // OpenID Connect Discovery URL
	authURL          string // Authentication URL
	accessTokenURL   string // Access Token URL
	clientID         string // Client ID
	scope            string // Scope
}
type ApiKey struct {
	username     string
	value        string
	transferMode string
}

func (TestInterface) TableName() string {
	return "biz_test_interface"
}
