package model

import ()

type CustomWorkitemField struct {
	BaseModel

	CustomWorkitemId uint `json:"customWorkitemId"`
	CustomFieldId    uint `json:"customFieldId"` // null mean all
}

func (CustomWorkitemField) TableName() string {
	return "r_custom_workitem_field"
}
