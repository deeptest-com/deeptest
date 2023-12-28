package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ComponentSchema struct {
	BaseModel
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Content     string            `json:"content" gorm:"type:longtext"`
	ServeId     int64             `json:"serveId"`
	Examples    string            `json:"examples" gorm:"type:longtext"`
	Tags        string            `json:"tags"`
	Description string            `gorm:"type:text" json:"description"`
	Ref         string            `json:"ref"`
	SourceType  consts.SourceType `json:"sourceType" gorm:"default:0"`
}

func (ComponentSchema) TableName() string {
	return "biz_project_serve_component_schema"
}
