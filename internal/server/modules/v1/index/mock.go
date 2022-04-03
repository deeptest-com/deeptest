package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type TestExecModule struct {
	MockCtrl *controller.MockCtrl `inject:""`
}

func NewTestExecModule() *TestExecModule {
	return &TestExecModule{}
}

// Party 脚本
func (m *TestExecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/get", m.MockCtrl.Get).Name = "用于接口GET测试"
		index.Post("/post", m.MockCtrl.Post).Name = "用于接口POST测试"
	}
	return module.NewModule("/mock", handler)
}
