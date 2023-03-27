package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type Param struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Disabled bool   `json:"disabled"`
	Desc     string `json:"desc"`
	Required bool   `json:"required"`
}

type Header struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Disabled bool   `json:"disabled"`
	Desc     string `json:"desc"`
	Required bool   `json:"required"`
}

type Cookie struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Disabled bool   `json:"disabled"`
	Desc     string `json:"desc"`
	Required bool   `json:"required"`
}

type RequestBody struct {
	ID          int64      `json:"id"`
	MediaType   string     `json:"mediaType"`
	Description string     `json:"description"`
	SchemaRefId int64      `json:"schemaRefId"`
	SchemaItem  SchemaItem `json:"schemaItem"`
	Examples    string     `json:"examples"`
}

type ResponseBody struct {
	ID          int64      `json:"id"`
	MediaType   string     `json:"mediaType"`
	Code        string     `json:"code"`
	SchemaRefId int64      `json:"schemaRefId"`
	SchemaItem  SchemaItem `json:"schemaItem"`
	Headers     []Header   `json:"headers"`
	Examples    string     `json:"examples"`
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type BearerToken struct {
	Token string `json:"token"`
}
type OAuth20 struct {
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
}
type ApiKey struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	TransferMode string `json:"transferMode"`
}
type SchemaItem struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}
