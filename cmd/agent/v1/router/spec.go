package router

import (
	"github.com/deeptest-com/deeptest/cmd/agent/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type SpecModule struct {
	SpecCtrl *handler.SpecCtrl `inject:""`
}

// Party 上传文件模块
func (m *SpecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/parseSpec", m.SpecCtrl.Parse).Name = "解析导入文件内容"
	}
	return module.NewModule("/spec", handler)
}
