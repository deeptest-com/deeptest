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

func (r *ProcessorInvocationRepo) List(interfaceId int) (pos []model.ProcessorInvocation, err error) {
	err = r.DB.
		Select("id", "name").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *ProcessorInvocationRepo) Get(id uint) (invocation model.ProcessorInvocation, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&invocation).Error
	return
}

func (r *ProcessorInvocationRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.ProcessorInvocation{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

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
