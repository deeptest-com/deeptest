package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type DebugRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *DebugRepo) Save(invocation *model.DebugInvoke) (err error) {
	err = r.DB.Save(invocation).Error
	return
}

func (r *DebugRepo) List(endpointInterfaceId, debugInterfaceId int) (pos []model.DebugInvoke, err error) {
	db := r.DB.Select("id", "name")

	if endpointInterfaceId > 0 {
		db.Where("endpoint_interface_id=?", endpointInterfaceId)
	} else if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	}

	err = db.Where("NOT deleted").
		Order("created_at DESC").
		Find(&pos).Error
	return
}

func (r *DebugRepo) Get(id uint) (invocation model.DebugInvoke, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&invocation).Error
	return
}

func (r *DebugRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugInvoke{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *DebugRepo) GetLast(endpointInterfaceId, debugInterfaceId int) (debug model.DebugInvoke, err error) {
	db := r.DB

	if endpointInterfaceId > 0 {
		db.Where("endpoint_interface_id=?", endpointInterfaceId)
	} else if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	}

	err = db.Where("NOT deleted").
		Order("created_at DESC").
		First(&debug).Error

	return
}
