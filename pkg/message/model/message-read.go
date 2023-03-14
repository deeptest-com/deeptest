package model

import "message/domain"

type MessageRead struct {
	Base
	domain.MessageReadBase
}

func (MessageRead) TableName() string {
	return "biz_message_read"
}
