package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DebugRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *DebugRepo) Save(invocation *model.Debug) (err error) {
	err = r.DB.Save(invocation).Error
	return
}

func (r *DebugRepo) List(interfaceId int) (pos []model.Debug, err error) {
	err = r.DB.
		Select("id", "name").
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *DebugRepo) Get(id uint) (invocation model.Debug, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&invocation).Error
	return
}

func (r *DebugRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.Debug{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *DebugRepo) GetLast(interfaceId int) (debug model.Debug, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at DESC").
		First(&debug).Error
	return
}

func (r *DebugRepo) Tested(interfaceId uint) (res bool, err error) {
	var count int64
	err = r.DB.Model(&model.Debug{}).Where("interface_id=?", interfaceId).Count(&count).Error
	if err != nil {
		return
	}
	res = count > 0
	return
}
