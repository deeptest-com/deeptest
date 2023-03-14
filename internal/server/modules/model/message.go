package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type Message struct {
	BaseModel
	v1.MessageBase
}

func (Message) TableName() string {
	return "biz_message"
}
