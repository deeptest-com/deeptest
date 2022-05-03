package model

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
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

	BasicAuth   domain.BasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken domain.BearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     domain.OAuth20     `gorm:"-" json:"oAuth20"`
	ApiKey      domain.ApiKey      `gorm:"-" json:"apiKey"`

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

type InterfaceExtractor struct {
	BaseModel
	Src  serverConsts.ExtractorSrc  `json:"src"`
	Type serverConsts.ExtractorType `json:"type"`
	Key  string                     `json:"key"`

	Expression string `json:"expression"`
	Prop       string `json:"prop"`

	Variable string `json:"variable"`

	Result      string `json:"result"`
	InterfaceId uint   `json:"interfaceId"`
}

func (InterfaceExtractor) TableName() string {
	return "biz_interface_extractor"
}

type InterfaceCheckpoint struct {
	BaseModel
	Type serverConsts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator serverConsts.CheckpointOperator `json:"operator"`
	Value    string                          `json:"value"`

	Result      serverConsts.CheckpointResult `json:"result"`
	InterfaceId uint                          `json:"interfaceId"`
}

func (InterfaceCheckpoint) TableName() string {
	return "biz_interface_checkpoint"
}
