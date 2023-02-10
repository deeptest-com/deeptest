package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type InterfaceBase struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	EndpointId uint   `json:"endpoint_id"`
	Security   string `json:"security"`
	IsLeaf     bool   `json:"isLeaf"`
	ParentId   uint   `json:"parentId"`
	ProjectId  uint   `json:"projectId"`
	UseID      uint   `json:"useId"`

	Ordr int `json:"ordr"`

	Slots iris.Map `gorm:"-" json:"slots"`

	Url    string            `json:"url"`
	Method consts.HttpMethod `gorm:"default:GET" json:"method"`

	Body     string                 `gorm:"default:{}" json:"body"`
	BodyType consts.HttpContentType `gorm:"default:''" json:"bodyType"`

	AuthorizationType string `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string `gorm:"default:''" json:"validationScript"`
}

type InterfaceParamBase struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Desc        string `json:"desc"`
	InterfaceId uint   `json:"interfaceId"`
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
	Value       string `json:"value"`
	Type        string `json:"type"`
	InterfaceId uint   `json:"interfaceId"`
}

type InterfaceCookieBase struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Value       string `json:"value"`
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
	InterfaceId uint   `json:"interfaceId"`
	MediaType   string `json:"mediaType"`
	Description string `json:"description"`
	SchemaRefId int64  `json:"schemaRefId"`
	Examples    string `json:"examples"`
}

type InterfaceResponseBodyBase struct {
	Code        uint   `json:"code"`
	InterfaceId uint   `json:"interfaceId"`
	MediaType   string `json:"mediaType"`
	Description string `json:"description"`
	SchemaRefId int64  `json:"schemaRefId"`
	Examples    string `json:"examples"`
}

type InterfaceRequestBodyItemBase struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Content       string `json:"content"`
	RequestBodyId uint   `json:"requestBodyId"`
}

type InterfaceResponseBodyItemBase struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Content        string `json:"content"`
	ResponseBodyId uint   `json:"ResponseBodyId"`
}

type InterfaceResponseBodyHeaderBase struct {
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Value          string `json:"value"`
	Type           string `json:"type"`
	ResponseBodyId uint   `json:"responseBodyId"`
}
