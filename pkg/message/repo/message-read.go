package repo

import (
	"gorm.io/gorm"
	"message/domain"
	"message/model"
)

type MessageReadRepo struct {
	DB *gorm.DB `inject:""`
}

func NewMessageReadRepo() *MessageReadRepo {
	return &MessageReadRepo{}
}

func (r *MessageReadRepo) Create(message domain.MessageReadReq) (id uint, err error) {
	err = r.DB.Model(&model.MessageRead{}).Create(&message).Error
	if err != nil {
		return
	}

	id = message.Id
	return
}

func (r *MessageReadRepo) ListByMessageIds(messageIds []uint) (messages []model.MessageRead, err error) {
	err = r.DB.Model(&model.MessageRead{}).
		Where("message_id IN (?)", messageIds).Find(&messages).Error

	return
}
