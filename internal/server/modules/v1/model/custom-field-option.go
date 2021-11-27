package model

import ()

type CustomFieldOption struct {
	BaseModel

	Label string `json:"label"`
	Value string `json:"value"`
	Desc  string `json:"desc" gorm:"column:descr"`
	Order int    `json:"order" gorm:"column:ordr"`

	CustomFieldId uint `json:"customFieldId"`
}

func (CustomFieldOption) TableName() string {
	return "custom_field_option"
}
