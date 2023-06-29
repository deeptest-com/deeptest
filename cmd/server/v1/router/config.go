package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type ConfigModule struct {
	ConfigCtrl *handler.ConfigCtrl `inject:""`
}

// Party 脚本
func (m *ConfigModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Get("/", m.ConfigCtrl.Get).Name = "获取服务端配置"
	}
	return module.NewModule("/configs", handler)
}
