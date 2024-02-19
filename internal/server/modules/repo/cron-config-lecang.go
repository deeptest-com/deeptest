package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type CronConfigLecangRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *CronConfigLecangRepo) Create(config model.CronConfigLecang) (id uint, err error) {
	err = r.DB.Model(&model.CronConfigLecang{}).Create(&config).Error
	if err != nil {
		return
	}

	id = config.ID
	return
}

func (r *CronConfigLecangRepo) Update(config model.CronConfigLecang) (err error) {
	err = r.DB.Save(&config).Error

	return
}

func (r *CronConfigLecangRepo) Save(config model.CronConfigLecang) (id uint, err error) {
	err = r.DB.Save(&config).Error
	if err != nil {
		return
	}

	id = config.ID

	return
}

func (r *CronConfigLecangRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.CronConfigLecang{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error

	return
}

func (r *CronConfigLecangRepo) GetById(id uint) (config model.CronConfigLecang, err error) {
	err = r.DB.Model(&model.CronConfigLecang{}).
		Where("id = ?", id).
		Find(&config).Error

	return
}
