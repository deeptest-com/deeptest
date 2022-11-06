package router

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type FileModule struct {
	FileCtrl *handler.FileCtrl `inject:""`
}

func NewFileModule() *FileModule {
	return &FileModule{}
}

// Party 上传文件模块
func (m *FileModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/", iris.LimitRequestBodySize(config.CONFIG.MaxSize+1<<20), m.FileCtrl.Upload).Name = "上传文件"
	}
	return module.NewModule("/upload", handler)
}
