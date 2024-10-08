package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type FileModule struct {
	FileCtrl *handler.FileCtrl `inject:""`
}

// Party 上传文件模块
func (m *FileModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/", iris.LimitRequestBodySize(1024*iris.MB),
			m.FileCtrl.Upload).Name = "上传文件"
		index.Post("/do", iris.LimitRequestBodySize(1024*iris.MB),
			m.FileCtrl.Do).Name = "上传文件"
	}

	return module.NewModule("/upload", handler)
}
