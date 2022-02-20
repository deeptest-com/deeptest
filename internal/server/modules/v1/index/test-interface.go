package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type TestInterfaceModule struct {
	TestInterfaceCtrl *controller.TestInterfaceCtrl `inject:""`
}

func NewTestScriptModule() *TestInterfaceModule {
	return &TestInterfaceModule{}
}

// Party 脚本
func (m *TestInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.TestInterfaceCtrl.Load).Name = "接口数据"
		index.Get("/{id:uint}", m.TestInterfaceCtrl.Get).Name = "接口详情"
		index.Post("/", m.TestInterfaceCtrl.Create).Name = "新建接口"
		index.Put("/{id:uint}", m.TestInterfaceCtrl.Update).Name = "更新接口"
		index.Delete("/{id:uint}", m.TestInterfaceCtrl.Delete).Name = "删除接口"
		index.Post("/move", m.TestInterfaceCtrl.Move).Name = "移动接口"
	}
	return module.NewModule("/interfaces", handler)
}
