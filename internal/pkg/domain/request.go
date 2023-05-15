package domain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type DebugResponse struct {
	Id uint `json:"id"`

	StatusCode    consts.HttpRespCode `json:"statusCode"`
	StatusContent string              `json:"statusContent"`

	Headers []Header     `gorm:"-" json:"headers"`
	Cookies []ExecCookie `gorm:"-" json:"cookies"`

	Content     string                 `gorm:"default:''" json:"content"`
	ContentType consts.HttpContentType `json:"contentType"`

	ContentLang    consts.HttpRespLangType `json:"contentLang"`
	ContentCharset consts.HttpRespCharset  `json:"contentCharset"`
	ContentLength  int                     `json:"contentLength"`

	Time int64 `json:"time"`
}

type BaseRequest struct {
	Method  consts.HttpMethod `gorm:"default:GET" json:"method"`
	Url     string            `json:"url"`
	Params  []Param           ` json:"params"`
	Headers []Header          ` json:"headers"`
	Cookies []ExecCookie      ` json:"cookies"` // from cookie processor in scenario

	Body               string                   `gorm:"default:{}" json:"body"`
	BodyFormData       []BodyFormDataItem       `gorm:"-" json:"bodyFormData"`
	BodyFormUrlencoded []BodyFormUrlEncodedItem `gorm:"-" json:"bodyFormUrlencoded"`
	BodyType           consts.HttpContentType   `gorm:"default:json" json:"bodyType"`
	BodyLang           consts.HttpRespLangType  `gorm:"default:json" json:"bodyLang"`

	AuthorizationType consts.AuthorType `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string            `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string            `gorm:"default:''" json:"validationScript"`

	BasicAuth   BasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken BearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     OAuth20     `gorm:"-" json:"oauth20"`
	ApiKey      ApiKey      `gorm:"-" json:"apiKey"`
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

type SchemaItem struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}
