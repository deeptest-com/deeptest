package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ConfigService struct {
	ConfigRepo *repo.ConfigRepo `inject:""`
}

func (s *ConfigService) Get(key string) (config model.SysConfig, err error) {
	config, err = s.ConfigRepo.Get(key)
	return
}

func (s *ConfigService) Save(req model.SysConfig) (err error) {
	err = s.ConfigRepo.Save(req)
	return
}
