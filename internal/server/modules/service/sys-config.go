package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"gorm.io/gorm"
)

type ConfigService struct {
	ConfigRepo *repo.ConfigRepo `inject:""`
}

func (s *ConfigService) Get(key string) (value string, err error) {
	config, err := s.ConfigRepo.Get(key)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}
	value = config.Value
	return
}

func (s *ConfigService) Save(req model.SysConfig) (err error) {
	err = s.ConfigRepo.Save(req)
	return
}
