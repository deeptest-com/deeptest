package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type CheckpointModule struct {
	CheckpointCtrl *handler.CheckpointCtrl `inject:""`
}

// Party 检查点
func (m *CheckpointModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/{id:uint}", m.CheckpointCtrl.Get).Name = "检查点详情"
		index.Put("/", m.CheckpointCtrl.Update).Name = "更新检查点"
	}

	return module.NewModule("/checkpoints", handler)
}
