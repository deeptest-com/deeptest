package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PreConditionModule struct {
	PreConditionCtrl *handler.PreConditionCtrl `inject:""`
}

// Party 前置条件
func (m *PreConditionModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/getScript", m.PreConditionCtrl.GetScript).Name = "前置条件列表"
		index.Post("/", m.PreConditionCtrl.Create).Name = "新建前置条件"
		index.Delete("/{id:uint}", m.PreConditionCtrl.Delete).Name = "删除前置条件"
		index.Post("/{id:uint}/disable", m.PreConditionCtrl.Disable).Name = "禁用前置条件"
		index.Post("/move", m.PreConditionCtrl.Move).Name = "移动后置条件"
		index.Get("/resetForCase", m.PreConditionCtrl.ResetForCase).Name = "重制用例的前置条件"
	}

	return module.NewModule("/preConditions", handler)
}
