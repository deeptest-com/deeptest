package router

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type SpecModule struct {
	SpecCtrl *handler.SpecCtrl `inject:""`
}

func NewSpecModule() *SpecModule {
	return &SpecModule{}
}

// Party 上传文件模块
func (m *SpecModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/loadSpec", m.SpecCtrl.Load).Name = "解析导入文件内容"
	}
	return module.NewModule("/spec", handler)
}
