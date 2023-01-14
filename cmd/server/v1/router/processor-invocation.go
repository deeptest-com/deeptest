package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProcessorInvocationModule struct {
	ProcessorInvocationCtrl *handler.ProcessorInvocationCtrl `inject:""`
}

// Party 脚本
func (m *ProcessorInvocationModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/loadInterfaceExecData", m.ProcessorInvocationCtrl.LoadInterfaceExecData).Name = "获取场景接口"
		index.Post("/submitInterfaceInvokeResult", m.ProcessorInvocationCtrl.SubmitInterfaceInvokeResult).Name = "提交场景接口测试结果"

		index.Get("/getLastResp", m.ProcessorInvocationCtrl.GetLastResp).Name = "最后一次调用详情"
	}
	return module.NewModule("/processors/invocations", handler)
}
