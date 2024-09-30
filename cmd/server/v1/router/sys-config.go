package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type ConfigModule struct {
	ConfigCtrl *handler.ConfigCtrl `inject:""`
}

// Party 脚本
func (m *ConfigModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/", m.ConfigCtrl.Get).Name = "获取服务端配置"
		index.Get("/getValue", m.ConfigCtrl.GetValue).Name = "获取配置"
		index.Post("/", m.ConfigCtrl.Save).Name = "存储配置"
	}
	return module.NewModule("/configs", handler)
}
