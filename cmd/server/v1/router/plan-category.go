package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type PlanCategoryModule struct {
	PlanCategoryCtrl *handler.PlanCategoryCtrl `inject:""`
}

// Party 场景
func (m *PlanCategoryModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/load", m.PlanCategoryCtrl.LoadTree).Name = "分类树状数据"
		index.Get("/{id:uint}", m.PlanCategoryCtrl.Get).Name = "分类详情"
		index.Post("/", m.PlanCategoryCtrl.Create).Name = "新建分类"
		index.Put("/", m.PlanCategoryCtrl.Update).Name = "更新分类"
		index.Put("/{id:uint}/updateName", m.PlanCategoryCtrl.UpdateName).Name = "更新节点名称"
		index.Delete("/{id:uint}", m.PlanCategoryCtrl.Delete).Name = "删除节点"
		index.Post("/move", m.PlanCategoryCtrl.Move).Name = "移动节点"
	}

	return module.NewModule("/plans/categories", handler)
}
