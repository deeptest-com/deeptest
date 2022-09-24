package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type InvocationRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *InvocationRepo) List(interfaceId int) (pos []model.Invocation, err error) {
	err = r.DB.
		Select("id", "name").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *InvocationRepo) Get(id uint) (invocation model.Invocation, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&invocation).Error
	return
}

func (r *InvocationRepo) GetLast(interfaceId int) (invocation model.Invocation, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		First(&invocation).Error
	return
}

func (r *InvocationRepo) Save(invocation *model.Invocation) (err error) {
	err = r.DB.Save(invocation).Error
	return
}

func (r *InvocationRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Invocation{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}
