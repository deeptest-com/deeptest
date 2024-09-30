package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type MockJsModule struct {
	MockJsCtrl *handler.MockJsCtrl `inject:""`
}

// Party mockjs表达式
func (m *MockJsModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.PartyFunc("/expressions", func(party iris.Party) {
			party.Get("/", m.MockJsCtrl.ListExpressions).Name = "表达式列表"
			party.Post("/evaluate", m.MockJsCtrl.EvaluateExpression).Name = "评估表达式"
		})
	}

	return module.NewModule("/mockjs", handler)
}
