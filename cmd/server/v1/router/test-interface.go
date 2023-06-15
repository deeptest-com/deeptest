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

		index.Get("/", m.TestInterfaceCtrl.Load).Name = "获取测试接口"
		index.Post("/", m.TestInterfaceCtrl.Save).Name = "新建测试接口"
		index.Put("/", m.TestInterfaceCtrl.Update).Name = "更新测试接口"
		index.Delete("/{id:uint}", m.TestInterfaceCtrl.Delete).Name = "删除测试接口"
		index.Post("/move", m.TestInterfaceCtrl.Move).Name = "移动节点"
	}
	return module.NewModule("/testInterfaces", handler)
}
