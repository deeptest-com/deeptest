package model

import (
	"github.com/aaronchen2k/deeptest/pkg/message/domain"
	"time"
)

type Base struct {
	ID        uint       `gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL" json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

type Message struct {
	Base
	domain.MessageBase
}

func (Message) TableName() string {
	return "biz_message"
}
