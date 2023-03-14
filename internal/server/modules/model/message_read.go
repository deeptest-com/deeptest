package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type MessageRead struct {
	BaseModel
	v1.MessageReadBase
}

func (MessageRead) TableName() string {
	return "biz_message_read"
}
