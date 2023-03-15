package model

import "github.com/aaronchen2k/deeptest/pkg/message/domain"

type MessageRead struct {
	Base
	domain.MessageReadBase
}

func (MessageRead) TableName() string {
	return "biz_message_read"
}
