package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type InvocationRequest struct {
	Id   uint   `gorm:"-" json:"id"`
	Name string `json:"name"`

	BaseRequest

	//ProjectId uint `json:"projectId"`
}

type InvocationResponse struct {
	Id uint `json:"id"`

	StatusCode    consts.HttpRespCode `json:"statusCode"`
	StatusContent string              `json:"statusContent"`

	Headers     []domain.Header        `gorm:"-" json:"headers"`
	Content     string                 `gorm:"default:''" json:"content"`
	ContentType consts.HttpContentType `json:"contentType"`

	ContentLang    consts.HttpRespLangType `json:"contentLang"`
	ContentCharset consts.HttpRespCharset  `json:"contentCharset"`
	ContentLength  int                     `json:"contentLength"`

	Time int64 `json:"time"`
}

type SubmitInvocationResultRequest struct {
	UsedBy   string             `json:"usedBy"`
	Request  InvocationRequest  `json:"request"`
	Response InvocationResponse `json:"response"`
}

type BaseRequest struct {
	Method  consts.HttpMethod `gorm:"default:GET" json:"method"`
	Url     string            `json:"url"`
	Params  []Param           ` json:"params"`
	Headers []Header          ` json:"headers"`

	Body               string                   `gorm:"default:{}" json:"body"`
	BodyFormData       []BodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []BodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`
	BodyType           consts.HttpContentType   `gorm:"default:json" json:"bodyType"`
	BodyLang           consts.HttpRespLangType  `gorm:"default:json" json:"bodyLang"`

	AuthorizationType consts.AuthorType `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string            `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string            `gorm:"default:''" json:"validationScript"`

	BasicAuth   BasicAuth   ` json:"basicAuth"`
	BearerToken BearerToken ` json:"bearerToken"`
	OAuth20     OAuth20     ` json:"oauth20"`
	ApiKey      ApiKey      ` json:"apiKey"`
}

type Header struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type Param struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type BodyFormDataItem struct {
	Name        string              `json:"name"`
	Value       string              `json:"value"`
	Type        consts.FormDataType `json:"type"`
	Desc        string              `json:"desc"`
	InterfaceId uint                `json:"interfaceId"`
}

type BodyFormUrlEncodedItem struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Desc        string `json:"desc"`
	InterfaceId uint   `json:"interfaceId"`
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
