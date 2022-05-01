package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type CheckpointModule struct {
	CheckpointCtrl *controller.CheckpointCtrl `inject:""`
}

// Party 检查点
func (m *CheckpointModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.CheckpointCtrl.List).Name = "检查点列表"
		index.Get("/{id:uint}", m.CheckpointCtrl.Get).Name = "检查点详情"
		index.Post("/", m.CheckpointCtrl.Create).Name = "新建检查点"
		index.Put("/", m.CheckpointCtrl.Update).Name = "更新检查点"
		index.Delete("/{id:uint}", m.CheckpointCtrl.Delete).Name = "删除检查点"
	}

	return module.NewModule("/checkpoints", handler)
}
