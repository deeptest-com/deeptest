package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ConfigRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ConfigRepo) Get(key string) (config model.SysConfig, err error) {
	err = r.DB.
		Where("k = ?", key).
		First(&config).Error
	return
}

func (r *ConfigRepo) Save(req model.SysConfig) (err error) {
	config, err := r.Get(req.Key)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if config.Key == "" || err == gorm.ErrRecordNotFound {
		if err = r.DB.Model(&req).Create(req).Error; err != nil {
			return err
		}
	}

	err = r.DB.Model(&model.SysConfig{}).
		Where("k = ?", req.Key).
		Update("v", req.Value).Error
	return
}
