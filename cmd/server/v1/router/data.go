package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/kataras/iris/v12"
)

type DataModule struct {
	DataCtrl    *handler.DataCtrl    `inject:""`
	DataService *service.DataService `inject:""`
}

// Party 初始化模块
func (m *DataModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/initdb", m.DataCtrl.Init)
		index.Get("/checkdb", m.DataCtrl.Check)
	}

	//m.DataService.InitDB(serverDomain.DataReq{ClearData: true})

	return module.NewModule("/init", handler)
}
