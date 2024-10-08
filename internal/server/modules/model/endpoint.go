package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"time"
)

type Endpoint struct {
	BaseModel
	Title        string              `json:"title"`
	ProjectId    uint                `json:"projectId" gorm:"index:idx_projectId_serveId"`
	ServeId      uint                `json:"serveId" gorm:"index:idx_projectId_serveId"`
	ServerId     uint                `json:"serverId"`
	Path         string              `json:"path"`
	Version      string              `json:"version"`
	CreateUser   string              `json:"createUser"`
	UpdateUser   string              `json:"updateUser"`
	Status       int64               `json:"status"`
	CategoryId   int64               `gorm:"index:idx_category_id" json:"categoryId"`
	PathParams   []EndpointPathParam `gorm:"-" json:"pathParams"`
	Interfaces   []EndpointInterface `gorm:"-" json:"interfaces"`
	Versions     []EndpointVersion   `gorm:"-" json:"versions"`
	Tags         []string            `gorm:"-" json:"tags"`
	ServeName    string              `gorm:"-" json:"serveName"`
	Description  string              `gorm:"type:text" json:"description"`
	SerialNumber string              `json:"serialNumber"`
	Curl         string              `gorm:"-" json:"curl"`
	SourceType   consts.SourceType   `json:"sourceType"`
	Maintainer   string              `gorm:"-" json:"maintainer"`
	Methods      []consts.HttpMethod `gorm:"-" json:"methods"`

	AdvancedMockDisabled bool `json:"advancedMockDisabled"`
	ScriptMockDisabled   bool `json:"scriptMockDisabled"`

	GlobalParams  []EnvironmentParam   `gorm:"-" json:"globalParams"`
	Snapshot      string               `gorm:"type:longtext" json:"snapshot"`
	ChangedTime   *time.Time           `json:"changedTime,omitempty"`
	ChangedStatus consts.ChangedStatus `gorm:"default:1" json:"changedStatus,omitempty"`

	IsFavorite bool `gorm:"-" json:"isFavorite"`
}

func (Endpoint) TableName() string {
	return "biz_endpoint"
}

type EndpointPathParam struct {
	EndpointInterfaceParam
	EndpointId uint `json:"endpointId"`
}

func (EndpointPathParam) TableName() string {
	return "biz_endpoint_path_param"
}

type EndpointVersion struct {
	BaseModel
	Version    string `json:"version"`
	EndpointId uint   `json:"endpointId"`
}

func (EndpointVersion) TableName() string {
	return "biz_endpoint_version"
}
