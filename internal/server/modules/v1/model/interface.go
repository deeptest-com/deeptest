package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type Interface struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	IsDir     bool `json:"isDir"`
	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int          `json:"ordr"`
	Children []*Interface `gorm:"-" json:"children"`

	Slots iris.Map `gorm:"-" json:"slots"`

	Url               string                 `json:"url"`
	Method            string                 `gorm:"default:GET" json:"method"`
	Params            []InterfaceParam       `gorm:"-" json:"params"`
	Headers           []InterfaceHeader      `gorm:"-" json:"headers"`
	Body              string                 `gorm:"default:{}" json:"body"`
	BodyType          consts.HttpContentType `gorm:"default:''" json:"bodyType"`
	AuthorizationType string                 `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	BasicAuth   InterfaceBasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken InterfaceBearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     InterfaceOAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      InterfaceApiKey      `gorm:"-" json:"apiKey"`

	EnvironmentId uint `json:"environmentId"`

	InterfaceExtractors  []InterfaceExtractor  `gorm:"-" json:"interfaceExtractors"`
	InterfaceCheckpoints []InterfaceCheckpoint `gorm:"-" json:"interfaceCheckpoints"`
}

func (Interface) TableName() string {
	return "biz_interface"
}

type InterfaceParam struct {
	BaseModel
	Name        string `json:"name"`
	Value       string `json:"value"`
	InterfaceId uint   `json:"interfaceId"`
}

func (InterfaceParam) TableName() string {
	return "biz_interface_param"
}

type InterfaceHeader struct {
	BaseModel
	Name        string `json:"name"`
	Value       string `json:"value"`
	InterfaceId uint   `json:"interfaceId"`
}

func (InterfaceHeader) TableName() string {
	return "biz_interface_header"
}

type InterfaceBasicAuth struct {
	BaseModel

	Username string `json:"username"`
	Password string `json:"password"`

	InterfaceId uint `json:"interfaceId"`
}

func (InterfaceBasicAuth) TableName() string {
	return "biz_interface_basic_auth"
}

type InterfaceBearerToken struct {
	BaseModel

	Token       string `json:"token"`
	InterfaceId uint   `json:"interfaceId"`
}

func (InterfaceBearerToken) TableName() string {
	return "biz_interface_bearer_token"
}

type InterfaceOAuth20 struct {
	BaseModel

	AccessToken  string `json:"accessToken"`
	HeaderPrefix string `json:"headerPrefix" gorm:"default:Bearer"`

	Name           string           `json:"name"`
	GrantType      consts.GrantType `json:"grantType" gorm:"default:authorizationCode"`
	CallbackUrl    string           `json:"callbackUrl"`
	AuthURL        string           `json:"authURL"`
	AccessTokenURL string           `json:"accessTokenURL"`
	ClientID       string           `json:"clientID"`
	ClientSecret   string           `json:"clientSecret"`
	Scope          string           `json:"scope"`
	State          string           `json:"state"`

	ClientAuthentication consts.ClientAuthenticationWay `json:"clientAuthentication" gorm:"default:sendAsBasicAuthHeader"`

	InterfaceId uint `json:"interfaceId"`
}

func (InterfaceOAuth20) TableName() string {
	return "biz_interface_oauth20"
}

type InterfaceApiKey struct {
	BaseModel

	Key          string `json:"key"`
	Value        string `json:"value"`
	TransferMode string `json:"transferMode"`

	InterfaceId uint `json:"interfaceId"`
}

func (InterfaceApiKey) TableName() string {
	return "biz_interface_apikey"
}

type InterfaceExtractor struct {
	BaseModel
	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `json:"expression"`
	Prop       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable string `json:"variable"`

	Result      string `json:"result"`
	InterfaceId uint   `json:"interfaceId"`
	LogId       uint   `json:"logId"`
}

func (InterfaceExtractor) TableName() string {
	return "biz_interface_extractor"
}

type InterfaceCheckpoint struct {
	BaseModel
	Type consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	Result      consts.ResultStatus `json:"result"`
	InterfaceId uint                `json:"interfaceId"`
}

func (InterfaceCheckpoint) TableName() string {
	return "biz_interface_checkpoint"
}
