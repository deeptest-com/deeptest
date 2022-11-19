package service

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/fatih/color"
)

type InitService struct {
}

func NewInitService() *InitService {
	return &InitService{}
}

func (s *InitService) InitModels() {
	//if !_commonUtils.IsRelease() {
	err := dao.GetDB().AutoMigrate(
		model.Project{},
		model.Environment{},
		model.EnvironmentVar{},
		model.Invocation{},
	)
	if err != nil {
		color.Yellow(fmt.Sprintf("初始化数据表错误 ：%+v", err))
	}
	//}
}
