package service

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	"gorm.io/gorm"
)

type ConfigService struct {
	ConfigRepo *repo.ConfigRepo `inject:""`
}

func (s *ConfigService) Get(tenantId consts.TenantId, key string) (value string, err error) {
	config, err := s.ConfigRepo.Get(tenantId, key)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}
	value = config.Value
	return
}

func (s *ConfigService) Save(tenantId consts.TenantId, req model.SysConfig) (err error) {
	err = s.ConfigRepo.Save(tenantId, req)
	return
}
