package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type LlmToolModule struct {
	LlmToolCtrl *handler.LlmToolCtrl `inject:""`
}

// Party 项目
func (m *LlmToolModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.LlmToolCtrl.List).Name = "工具大模型列表"
		index.Get("/{id:uint}", m.LlmToolCtrl.Get).Name = "工具大模型详情"
		index.Post("/", m.LlmToolCtrl.Save).Name = "保存工具大模型"
		index.Put("/updateName", m.LlmToolCtrl.UpdateName).Name = "修改工具大模型名称"
		index.Delete("/{id:uint}", m.LlmToolCtrl.Delete).Name = "删除工具大模型"

		index.Put("/{id:uint}/setDefault", m.LlmToolCtrl.SetDefault).Name = "激活默认工具大模型"
		index.Put("/{id:uint}/disable", m.LlmToolCtrl.Disable).Name = "禁用工具大模型"
	}
	return module.NewModule("/llms", handler)
}
