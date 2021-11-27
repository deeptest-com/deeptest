package model

import ()

type ProjectPermission struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`
}

func (ProjectPermission) TableName() string {
	return "project_permission"
}
