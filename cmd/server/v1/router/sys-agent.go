package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type SysAgentModule struct {
	SysAgentCtrl *handler.SysAgentCtrl `inject:""`
}

// Party 脚本
func (m *SysAgentModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.SysAgentCtrl.List).Name = "列出执行代理"
		index.Get("/{id:uint}", m.SysAgentCtrl.Get).Name = "获取执行代理详情"
		index.Post("/", m.SysAgentCtrl.Save).Name = "保存执行代理"

		index.Put("/updateName", m.SysAgentCtrl.UpdateName).Name = "修改名称"
		index.Put("/{id:uint}/disable", m.SysAgentCtrl.Disable).Name = "禁用"
		index.Delete("/{id:uint}", m.SysAgentCtrl.Delete).Name = "删除"
	}
	return module.NewModule("/agents", handler)
}
