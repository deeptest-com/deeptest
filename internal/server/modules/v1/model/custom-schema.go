package model

import ()

type CustomSchema struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	OrgId uint `json:"orgId" gorm:"comment:'所属组织'"`
}

func (CustomSchema) TableName() string {
	return "custom_schema"
}
