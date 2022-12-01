package router

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type InvocationModule struct {
	InvocationCtrl *handler.InvocationCtrl `inject:""`
}

// Party 脚本
func (m *InvocationModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/invoke", m.InvocationCtrl.Invoke).Name = "调用接口测试"
	}
	return module.NewModule("/invocations", handler)
}
