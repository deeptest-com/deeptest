package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ProcessorInvocationRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ProcessorInvocationRepo) Save(invocation *model.ProcessorInvocation) (err error) {
	err = r.DB.Save(invocation).Error
	return
}

func (r *ProcessorInvocationRepo) GetLast(interfaceId int) (invocation model.ProcessorInvocation, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		First(&invocation).Error
	return
}
