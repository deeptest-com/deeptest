package router

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type ExecModule struct {
	ExecInterfaceCtrl *handler.ExecInterfaceCtrl `inject:""`
}

// Party 脚本
func (m *ExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/call", m.ExecInterfaceCtrl.Call).Name = "调用接口测试"
	}
	return module.NewModule("/exec", handler)
}
