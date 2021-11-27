package model

import ()

type TimeLog struct {
	BaseModel

	Content string `json:"content"`
	Hours   uint   `json:"hours"`

	EntityId   uint `json:"entityId"`
	WorkitemId uint `json:"workitemId"`
}

func (TimeLog) TableName() string {
	return "biz_timelog"
}
