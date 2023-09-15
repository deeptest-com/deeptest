package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type InterfaceBase struct {
	Name        string `json:"name"`
	OperationId string `json:"operationId"`
	Description string `json:"description"`
	Desc        string `json:"desc"`
	Security    string `json:"security"`
	//IsDir       bool   `json:"isDir"`
	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr int `json:"ordr"`

	Slots iris.Map `gorm:"-" json:"slots"`

	InterfaceConfigBase
}

type InterfaceConfigBase struct {
	Url      string                 `gorm:"default:''" json:"url"`
	Method   consts.HttpMethod      `gorm:"default:GET" json:"method"`
	Body     string                 `gorm:"type:text" json:"body"`
	BodyType consts.HttpContentType `gorm:"default:'application/json'" json:"bodyType"`

	AuthorizationType string `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string `gorm:"default:''" json:"validationScript"`
	Version           string `gorm:"default:''" json:"Version"`
}

type InterfaceParamBase struct {
	Name        string         `json:"name"`
	Value       string         `json:"value" gorm:"type:text"`
	Type        string         `json:"type"`
	ParamIn     consts.ParamIn `json:"paramIn"`
	Desc        string         `json:"desc"`
	InterfaceId uint           `json:"interfaceId"`
}

type InterfaceBodyFormDataItemBase struct {
	Name        string              `json:"name"`
	Value       string              `json:"value"`
	Type        consts.FormDataType `json:"type"`
	Desc        string              `json:"desc"`
	InterfaceId uint                `json:"interfaceId"`
}

type InterfaceBodyFormUrlEncodedItemBase struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Desc        string `json:"desc"`
	InterfaceId uint   `json:"interfaceId"`
}

type InterfaceHeaderBase struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Value       string `gorm:"type:text" json:"value"`
	Type        string `json:"type"`
	InterfaceId uint   `json:"interfaceId"`
}

type InterfaceCookieBase struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Value       string `json:"value" gorm:"type:text"`
	Type        string `json:"type"`
	InterfaceId uint   `json:"interfaceId"`
}

type InterfaceBasicAuthBase struct {
	Username string `json:"username"`
	Password string `json:"password"`

	InterfaceId uint `json:"interfaceId"`
}

type InterfaceBearerTokenBase struct {
	Token       string `json:"token"`
	InterfaceId uint   `json:"interfaceId"`
}

type InterfaceOAuth20Base struct {
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

type InterfaceApiKeyBase struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	TransferMode string `json:"transferMode"`

	InterfaceId uint `json:"interfaceId"`
}

type InterfaceRequestBodyBase struct {
	InterfaceId uint   `json:"interfaceId" gorm:"index"`
	MediaType   string `json:"mediaType"`
	Description string `json:"description"`
	SchemaRefId int64  `json:"schemaRefId"`
	Examples    string `gorm:"type:text" json:"examples"`
}

type InterfaceResponseBodyBase struct {
	Code        string `json:"code"`
	InterfaceId uint   `json:"interfaceId" gorm:"index"`
	MediaType   string `json:"mediaType"`
	Description string `json:"description"`
	SchemaRefId int64  `json:"schemaRefId"`
	Examples    string `gorm:"type:text" json:"examples"`
}

type InterfaceRequestBodyItemBase struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Content       string `gorm:"type:longtext" json:"content"`
	RequestBodyId uint   `json:"requestBodyId" gorm:"index"`
}

type InterfaceResponseBodyItemBase struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Content        string `gorm:"type:longtext" json:"content"`
	ResponseBodyId uint   `json:"ResponseBodyId" gorm:"index"`
}

type InterfaceResponseBodyHeaderBase struct {
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Value          string `json:"value"`
	ResponseBodyId uint   `json:"responseBodyId" gorm:"index"`
}
