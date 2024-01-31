package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DebugInterfaceModule struct {
	DebugInterfaceCtrl *handler.DebugInterfaceCtrl `inject:""`
	DebugInvokeCtrl    *handler.DebugInvokeCtrl    `inject:""`
}

// Party 脚本
func (m *DebugInterfaceModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.PartyFunc("/interface", func(party iris.Party) {
			party.Post("/load", m.DebugInterfaceCtrl.Load).Name = "获取调试接口请求"
			party.Post("/loadForExec", m.DebugInterfaceCtrl.LoadForExec).Name = "获取调试接口用于执行"

			party.Post("/save", m.DebugInterfaceCtrl.Save).Name = "保存调试接口"
			party.Post("/saveAsCase", m.DebugInterfaceCtrl.SaveAsCase).Name = "另存为接口用例"

			party.Post("/loadCurl", m.DebugInterfaceCtrl.LoadCurl).Name = "获取接口cURL命令"
		})
	}
	return module.NewModule("/debugs", handler)
}
