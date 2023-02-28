package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ScenarioCategoryModule struct {
	ScenarioCategoryCtrl *handler.ScenarioCategoryCtrl `inject:""`
}

// Party 场景
func (m *ScenarioCategoryModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/load", m.ScenarioCategoryCtrl.LoadTree).Name = "分类树状数据"
		index.Get("/{id:uint}", m.ScenarioCategoryCtrl.Get).Name = "分类详情"
		index.Post("/", m.ScenarioCategoryCtrl.Create).Name = "新建分类"
		index.Put("/", m.ScenarioCategoryCtrl.Update).Name = "更新分类"
		index.Put("/{id:uint}/updateName", m.ScenarioCategoryCtrl.UpdateName).Name = "更新节点名称"
		index.Delete("/{id:uint}", m.ScenarioCategoryCtrl.Delete).Name = "删除节点"
		index.Post("/move", m.ScenarioCategoryCtrl.Move).Name = "移动节点"
	}

	return module.NewModule("/scenarios/categories", handler)
}
