package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type TestInterfaceModule struct {
	TestInterfaceCtrl *handler.TestInterfaceCtrl `inject:""`
}

// Party 脚本
func (m *TestInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/load", m.TestInterfaceCtrl.Load).Name = "获取测试接口"
		index.Post("/save", m.TestInterfaceCtrl.Save).Name = "保存调试接口"
	}
	return module.NewModule("/testInterfaces", handler)
}
