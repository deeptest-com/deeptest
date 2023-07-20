package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DebugModule struct {
	DebugInterfaceCtrl *handler.DebugInterfaceCtrl `inject:""`
	DebugInvokeCtrl    *handler.DebugInvokeCtrl    `inject:""`
}

// Party 脚本
func (m *DebugModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.PartyFunc("/interface", func(party iris.Party) {
			party.Post("/load", m.DebugInterfaceCtrl.Load).Name = "获取调试接口请求"
			party.Post("/loadForExec", m.DebugInterfaceCtrl.LoadForExec).Name = "获取调试接口用于执行"

			party.Post("/save", m.DebugInterfaceCtrl.Save).Name = "保存调试接口"
			party.Post("/saveAsCase", m.DebugInterfaceCtrl.SaveAsCase).Name = "另存为接口用例"
		})

		index.PartyFunc("/invoke", func(party iris.Party) {
			party.Get("/", m.DebugInvokeCtrl.List).Name = "调试记录列表"
			party.Get("/{id:uint}", m.DebugInvokeCtrl.GetAsInterface).Name = "调试记录详情"
			party.Get("/getLastResp", m.DebugInvokeCtrl.GetLastResp).Name = "获取调试结果"
			party.Get("/getResult", m.DebugInvokeCtrl.GetResult).Name = "获取最后调试记录响应"
			party.Delete("/{id:uint}", m.DebugInvokeCtrl.Delete).Name = "删除调试记录"

			party.Post("/submitResult", m.DebugInvokeCtrl.SubmitResult).Name = "Agent提交接口执行结果"
		})
	}
	return module.NewModule("/debugs", handler)
}
