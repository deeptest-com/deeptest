package model

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
)

type CustomPageField struct {
	BaseModel

	Type       consts.FieldType `json:"type"`
	Property   string           `json:"property"` // default is the field name
	ValueType  consts.ValueType `json:"valueType"`
	Default    string           `json:"property"`
	Format     consts.FieldType `json:"format"`
	IsRequired bool             `json:"isRequired"`
	Order      int              `json:"order" gorm:"column:ordr"`

	Code  string `json:"code"`
	Label string `json:"label"`
	Tips  string `json:"tips"`
	Desc  string `json:"desc" gorm:"column:descr"`

	Column int `json:"column"` // 1-24
	Row    int `json:"row"`    // >= 1

	FieldId uint `json:"fieldId"`
	PageId  uint `json:"pageId"`

	WorkitemId uint `json:"workitemId"`
}

func (CustomPageField) TableName() string {
	return "custom_page_field"
}
