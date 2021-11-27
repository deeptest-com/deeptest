package model

import ()

// CustomWorkitem 自定义工作项类型，包括系统内置的
type CustomWorkitem struct {
	BaseModel

	Name string `json:"name" gorm:"comment:'字段名称'"`
	Desc string `json:"desc" gorm:"column:descr;comment:'描述'"`

	OrgId uint `json:"orgId"`
}

func (CustomWorkitem) TableName() string {
	return "custom_workitem"
}
