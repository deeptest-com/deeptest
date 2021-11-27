package model

import ()

type Iteration struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	ProjectId uint `json:"projectId"`
}

func (Iteration) TableName() string {
	return "biz_iteration"
}
