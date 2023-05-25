package model

import "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type Message struct {
	BaseModel
	serverDomain.MessageBase
}

func (Message) TableName() string {
	return "biz_message"
}
