package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type FileModule struct {
	FileCtrl *handler.FileCtrl `inject:""`
}

// Party 上传文件模块
func (m *FileModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.PartyFunc("/upload", func(party iris.Party) {
			party.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

			party.Post("/", iris.LimitRequestBodySize(config.CONFIG.MaxSize*iris.MB),
				m.FileCtrl.Upload).Name = "上传文件"
			party.Post("/do", iris.LimitRequestBodySize(config.CONFIG.MaxSize*iris.MB),
				m.FileCtrl.Do).Name = "上传文件"
		})

		index.PartyFunc("/download", func(party iris.Party) {
			party.Get("/{path:path}", m.FileCtrl.Download).Name = "下载打包的资源文件"
		})
	}
	return module.NewModule("/", handler)
}
