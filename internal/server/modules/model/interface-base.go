package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/kataras/iris/v12"
)

type InterfaceBase struct {
	Name string `json:"name"`
	Desc string `json:"desc"`

	IsDir     bool `json:"isDir"`
	ParentId  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

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

type InterfaceExtractorBase struct {
	Src  consts.ExtractorSrc  `json:"src"`
	Type consts.ExtractorType `json:"type"`
	Key  string               `json:"key"`

	Expression string `json:"expression"`
	//NodeProp       string `json:"prop"`

	BoundaryStart    string `json:"boundaryStart"`
	BoundaryEnd      string `json:"boundaryEnd"`
	BoundaryIndex    int    `json:"boundaryIndex"`
	BoundaryIncluded bool   `json:"boundaryIncluded"`

	Variable     string                `json:"variable"`
	Scope        consts.ExtractorScope `json:"scope" gorm:"default:private"`
	DisableShare bool                  `json:"disableShare"`

	Result      string `json:"result" gorm:"type:text"`
	InterfaceId uint   `json:"interfaceId"`

	ProjectId uint `json:"projectId"`
}

type InterfaceCheckpointBase struct {
	Type consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	ActualResult string              `json:"actualResult"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
	InterfaceId  uint                `json:"interfaceId"`
}
