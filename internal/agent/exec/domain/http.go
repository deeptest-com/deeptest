package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type Request struct {
	Method            consts.HttpMethod      `gorm:"default:GET" json:"method"`
	Url               string                 `json:"url"`
	Params            []Param                ` json:"params"`
	Headers           []Header               ` json:"headers"`
	Body              string                 `gorm:"default:{}" json:"body"`
	BodyType          consts.HttpContentType `gorm:"default:json" json:"bodyType"`
	AuthorizationType consts.AuthorType      `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	BasicAuth   BasicAuth   ` json:"basicAuth"`
	BearerToken BearerToken ` json:"bearerToken"`
	OAuth20     OAuth20     ` json:"oauth20"`
	ApiKey      ApiKey      ` json:"apiKey"`

	InterfaceID uint `json:"interfaceID"`
	ProjectID   uint `json:"projectID"`
}

type Response struct {
	StatusCode    consts.HttpRespCode `json:"statusCode"`
	StatusContent string              `json:"statusContent"`

	Headers     []Header               ` json:"respHeaders"`
	Content     string                 `json:"content"`
	ContentType consts.HttpContentType `json:"contentType"`

	ContentLang    consts.HttpRespLangType `json:"contentLang"`
	ContentCharset consts.HttpRespCharset  `json:"contentCharset"`
	ContentLength  int                     `json:"contentLength"`

	Time int64 `json:"time"`
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
