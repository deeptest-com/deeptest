package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type InvocationModule struct {
	InvocationCtrl *handler.InvocationCtrl `inject:""`
}

// Party 脚本
func (m *InvocationModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/invoke", m.InvocationCtrl.Invoke).Name = "模拟测试接口"

		index.Get("/", m.InvocationCtrl.List).Name = "调用列表"
		index.Get("/{id:uint}", m.InvocationCtrl.GetAsInterface).Name = "调用详情"
		index.Delete("/{id:uint}", m.InvocationCtrl.Delete).Name = "删除调用"

		index.Get("/getLastResp", m.InvocationCtrl.GetLastResp).Name = "最后一次调用详情"
	}
	return module.NewModule("/invocations", handler)
}
