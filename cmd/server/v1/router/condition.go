package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ConditionModule struct {
	ConditionCtrl *handler.ConditionCtrl `inject:""`
}

// Party 后置条件
func (m *ConditionModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("", m.ConditionCtrl.List).Name = "前/后置条件列表"
		index.Post("/", m.ConditionCtrl.Create).Name = "新建前/后置条件"

		index.Delete("/{id:uint}", m.ConditionCtrl.Delete).Name = "删除前/后置条件"
		index.Post("/{id:uint}/disable", m.ConditionCtrl.Disable).Name = "禁用前/后置条件"
		index.Post("/disable", m.ConditionCtrl.Disable).Name = "移动前/后置条件"
		index.Post("/move", m.ConditionCtrl.Move).Name = "移动前/后置条件"
		index.Get("/resetForCase", m.ConditionCtrl.ResetForCase).Name = "重制用例的前/后置条件"
	}

	return module.NewModule("/conditions", handler)
}
