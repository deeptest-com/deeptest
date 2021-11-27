package model

import ()

type Attachment struct {
	BaseModel
	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`
	Path string `json:"path"`
}

func (Attachment) TableName() string {
	return "biz_attachment"
}
