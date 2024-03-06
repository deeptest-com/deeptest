package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ConfigRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ConfigRepo) Get(tenantId consts.TenantId, key string) (config model.SysConfig, err error) {
	err = r.GetDB(tenantId).
		Where("k = ?", key).
		First(&config).Error
	return
}

func (r *ConfigRepo) Save(tenantId consts.TenantId, req model.SysConfig) (err error) {
	config, err := r.Get(tenantId, req.Key)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if config.Key == "" || err == gorm.ErrRecordNotFound {
		if err = r.GetDB(tenantId).Model(&req).Create(req).Error; err != nil {
			return err
		}
	}

	err = r.GetDB(tenantId).Model(&model.SysConfig{}).
		Where("k = ?", req.Key).
		Update("v", req.Value).Error

	return
}
