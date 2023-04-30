package domain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type DebugReq struct {
	EndpointInterfaceId uint `json:"endpointInterfaceId"` // load by endpoint designer
	ScenarioProcessorId uint `json:"scenarioProcessorId"` // load by scenario designer

	UsedBy consts.UsedBy `json:"usedBy"`
}

type DebugData struct {
	EndpointInterfaceId uint          `gorm:"-" json:"endpointInterfaceId"`
	ScenarioProcessorId uint          `gorm:"-" json:"scenarioProcessorId"`
	UsedBy              consts.UsedBy `json:"usedBy"`

	BaseUrl         string                   `gorm:"-" json:"baseUrl"`
	ShareVars       []domain.VarKeyValuePair `gorm:"-" json:"shareVars"`
	EnvVars         []domain.VarKeyValuePair `gorm:"-" json:"envVars"`
	GlobalEnvVars   []domain.GlobalVar       `gorm:"-" json:"globalEnvVars"`
	GlobalParamVars []domain.GlobalParam     `gorm:"-" json:"globalParamVars"`

	//Datapools   domain.Datapools   `gorm:"-" json:"datapools"`

	Name string `gorm:"-" json:"name"`
	BaseRequest
}

type DebugResponse struct {
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

type SubmitDebugResultRequest struct {
	Request  DebugData     `json:"request"`
	Response DebugResponse `json:"response"`
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
