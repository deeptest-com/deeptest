package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PostConditionModule struct {
	PostConditionCtrl *handler.PostConditionCtrl `inject:""`
}

// Party 后置条件
func (m *PostConditionModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("", m.PostConditionCtrl.List).Name = "前置条件列表"
		index.Post("/", m.PostConditionCtrl.Create).Name = "新建后置条件"

		index.Delete("/{id:uint}", m.PostConditionCtrl.Delete).Name = "删除后置条件"
		index.Post("/{id:uint}/disable", m.PostConditionCtrl.Disable).Name = "禁用后置条件"
		index.Post("/disable", m.PostConditionCtrl.Disable).Name = "移动后置条件"
		index.Post("/move", m.PostConditionCtrl.Move).Name = "移动后置条件"
		index.Get("/resetForCase", m.PostConditionCtrl.ResetForCase).Name = "重制用例的后置条件"

	}

	return module.NewModule("/postConditions", handler)
}
