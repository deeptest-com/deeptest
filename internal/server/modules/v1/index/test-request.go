package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type TestRequestModule struct {
	TestRequestCtrl *controller.TestRequestCtrl `inject:""`
}

func NewTestRequestModule() *TestRequestModule {
	return &TestRequestModule{}
}

// Party 脚本
func (m *TestRequestModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.TestRequestCtrl.List).Name = "请求列表"
		index.Get("/{id:uint}", m.TestRequestCtrl.Get).Name = "请求详情"
		index.Delete("/{id:uint}", m.TestRequestCtrl.Delete).Name = "删除请求"

		index.Get("/loadHistory", m.TestRequestCtrl.LoadHistory).Name = "请求详情"
	}
	return module.NewModule("/requests", handler)
}
