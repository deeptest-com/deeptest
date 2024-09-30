package router

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/handler"
	"github.com/deeptest-com/deeptest/internal/pkg/core/module"
	"github.com/deeptest-com/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DatapoolModule struct {
	DatapoolCtrl *handler.DatapoolCtrl `inject:""`
}

// Party 项目
func (m *DatapoolModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Post("/index", m.DatapoolCtrl.Index).Name = "数据池列表"
		index.Get("/{id:uint}", m.DatapoolCtrl.Get).Name = "数据池详情"
		index.Post("/save", m.DatapoolCtrl.Save).Name = "保存数据池"
		index.Delete("/{id:uint}", m.DatapoolCtrl.Delete).Name = "删除数据池"
		index.Put("/{id:uint}/disable", m.DatapoolCtrl.Disable).Name = "禁用数据池"
	}
	return module.NewModule("/datapools", handler)
}
