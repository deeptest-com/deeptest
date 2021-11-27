package model

import ()

type Comment struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc" gorm:"column:descr"`

	CreatedBy uint `json:"createdBy"`

	EntityId   uint `json:"entityId"`
	WorkitemId uint `json:"workitemId"`
}

func (Comment) TableName() string {
	return "biz_comment"
}
