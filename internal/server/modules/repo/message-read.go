package repo

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MessageReadRepo struct {
	DB        *gorm.DB `inject:""`
	*BaseRepo `inject:""`
}

func (r *MessageReadRepo) Create(tenantId consts.TenantId, req v1.MessageReadReq) (id uint, err error) {
	messageRead := model.MessageRead{MessageReadBase: req.MessageReadBase}
	err = r.GetDB(tenantId).Model(&model.MessageRead{}).Create(&messageRead).Error
	if err != nil {
		logUtils.Errorf("add message-read error", zap.String("error:", err.Error()))
		return
	}

	id = messageRead.ID
	return
}

func (r *MessageReadRepo) ListByMessageIds(tenantId consts.TenantId, messageIds []uint) (messages []model.MessageRead, err error) {
	err = r.GetDB(tenantId).Model(&model.MessageRead{}).
		Where("message_id IN (?)", messageIds).Find(&messages).Error

	return
}
