package model

import "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type MessageRead struct {
	BaseModel
	serverDomain.MessageReadBase
}

func (MessageRead) TableName() string {
	return "biz_message_read"
}
