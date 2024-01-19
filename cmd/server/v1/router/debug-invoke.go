package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DebugInvokeModule struct {
	DebugInterfaceCtrl *handler.DebugInterfaceCtrl `inject:""`
	DebugInvokeCtrl    *handler.DebugInvokeCtrl    `inject:""`
}

// Party 脚本
func (m *DebugInvokeModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.PartyFunc("/invoke", func(party iris.Party) {
			party.Get("/", m.DebugInvokeCtrl.List).Name = "调试记录列表"
			party.Get("/{id:uint}", m.DebugInvokeCtrl.GetAsInterface).Name = "调试记录详情"
			party.Get("/getLastResp", m.DebugInvokeCtrl.GetLastResp).Name = "获取调试结果"
			party.Get("/getResult", m.DebugInvokeCtrl.GetResult).Name = "获取调试结果"
			party.Get("/getConsoleLog", m.DebugInvokeCtrl.GetLog).Name = "获取调试日志"
			party.Delete("/{id:uint}", m.DebugInvokeCtrl.Delete).Name = "删除调试记录"

			party.Post("/submitResult", m.DebugInvokeCtrl.SubmitResult).Name = "Agent提交接口执行结果"
		})
	}
	return module.NewModule("/debugs", handler)
}
