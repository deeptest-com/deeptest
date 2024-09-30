package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type CronConfigLecangRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *CronConfigLecangRepo) Create(tenantId consts.TenantId, config model.CronConfigLecang) (id uint, err error) {
	err = r.GetDB(tenantId).Model(&model.CronConfigLecang{}).Create(&config).Error
	if err != nil {
		return
	}

	id = config.ID
	return
}

func (r *CronConfigLecangRepo) Update(tenantId consts.TenantId, config model.CronConfigLecang) (err error) {
	err = r.GetDB(tenantId).Save(&config).Error

	return
}

func (r *CronConfigLecangRepo) Save(tenantId consts.TenantId, config model.CronConfigLecang) (id uint, err error) {
	err = r.GetDB(tenantId).Save(&config).Error
	if err != nil {
		return
	}

	id = config.ID

	return
}

func (r *CronConfigLecangRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.CronConfigLecang{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *CronConfigLecangRepo) GetById(tenantId consts.TenantId, id uint) (config model.CronConfigLecang, err error) {
	err = r.GetDB(tenantId).Model(&model.CronConfigLecang{}).
		Where("id = ?", id).
		Find(&config).Error

	return
}
