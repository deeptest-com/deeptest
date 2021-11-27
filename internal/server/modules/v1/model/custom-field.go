package model

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
)

type CustomField struct {
	BaseModel

	Name       string             `json:"name" gorm:"comment:'字段名称'"`
	Desc       string             `json:"desc" gorm:"column:descr;comment:'描述'"`
	Source     consts.FieldSource `json:"source" gorm:"comment:'字段来源'"`
	WidgetType consts.FieldSource `json:"widgetType" gorm:"comment:'控件类型'"`

	Default    string `json:"default" gorm:"comment:'默认值'"`
	IsRequired bool   `json:"isRequired" gorm:"comment:'是否必填'"`

	OrgId uint `json:"orgId" gorm:"comment:'所属组织'"`
}

func (CustomField) TableName() string {
	return "custom_field"
}
