package model

import ()

type CustomStatus struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	WorkitemId uint `json:"workitemId"`
}

func (CustomStatus) TableName() string {
	return "custom_status"
}
