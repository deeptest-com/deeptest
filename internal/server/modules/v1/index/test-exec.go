package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type TestExecModule struct {
	TestExecCtrl *controller.TestExecCtrl `inject:""`
}

func NewTestExecModule() *TestExecModule {
	return &TestExecModule{}
}

// Party 脚本
func (m *TestExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/test", m.TestExecCtrl.Test).Name = "测试接口"
	}
	return module.NewModule("/exec", handler)
}
