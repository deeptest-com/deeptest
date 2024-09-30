package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type CategoryModule struct {
	CategoryCtrl *handler.CategoryCtrl `inject:""`
}

// Party 场景
func (m *CategoryModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/load", m.CategoryCtrl.LoadTree).Name = "分类树状数据"
		index.Get("/{id:int}", m.CategoryCtrl.Get).Name = "分类详情"
		index.Post("/", m.CategoryCtrl.Create).Name = "新建分类"
		index.Put("/", m.CategoryCtrl.Update).Name = "更新分类"
		index.Put("/{id:uint}/updateName", m.CategoryCtrl.UpdateName).Name = "更新节点名称"
		index.Delete("/{id:uint}", m.CategoryCtrl.Delete).Name = "删除节点"
		index.Post("/move", m.CategoryCtrl.Move).Name = "移动节点"
		index.Post("/batchAddSchemaRoot", m.CategoryCtrl.BatchAddSchemaRoot).Name = "批量初始化schema的根分类"
		index.Get("/copy/{id:int}", m.CategoryCtrl.Copy).Name = "复制节点"
		index.Get("/loadChildren", m.CategoryCtrl.LoadChildren).Name = "分类树状数据"
	}

	return module.NewModule("/categories", handler)
}
