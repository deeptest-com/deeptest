package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProcessorDataModule struct {
	ProcessorDataCtrl *handler.ProcessorDataCtrl `inject:""`
}

// Party 场景
func (m *ProcessorDataModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/upload", m.ProcessorDataCtrl.Upload).Name = "上传数据文件"
	}

	return module.NewModule("/processors/data", handler)
}
