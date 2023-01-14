package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type InterfaceModule struct {
	InterfaceCtrl *handler.InterfaceCtrl `inject:""`
}

// Party 脚本
func (m *InterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/saveInterface", m.InterfaceCtrl.SaveInterface).Name = "保存接口"

		index.Get("/", m.InterfaceCtrl.Load).Name = "接口数据"
		index.Get("/{id:uint}", m.InterfaceCtrl.Get).Name = "接口详情"
		index.Post("/", m.InterfaceCtrl.Create).Name = "新建接口"
		index.Put("/{id:uint}", m.InterfaceCtrl.Update).Name = "更新接口"
		index.Put("/updateName", m.InterfaceCtrl.UpdateName).Name = "更新名称接口"

		index.Delete("/{id:uint}", m.InterfaceCtrl.Delete).Name = "删除接口"
		index.Post("/move", m.InterfaceCtrl.Move).Name = "移动接口"
	}
	return module.NewModule("/interfaces", handler)
}
