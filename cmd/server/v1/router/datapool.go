package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type DatapoolModule struct {
	DatapoolCtrl *handler.DatapoolCtrl `inject:""`
}

func NewDatapoolModule() *DatapoolModule {
	return &DatapoolModule{}
}

// Party 项目
func (m *DatapoolModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.DatapoolCtrl.List).Name = "数据池列表"
		index.Get("/{id:uint}", m.DatapoolCtrl.Get).Name = "数据池详情"
		index.Post("/", m.DatapoolCtrl.Create).Name = "保存数据池"
		index.Put("/", m.DatapoolCtrl.SaveData).Name = "保存数据池数据"
		index.Delete("/{id:uint}", m.DatapoolCtrl.Delete).Name = "删除数据池"
		index.Post("/{id:uint}/upload", m.DatapoolCtrl.Upload).Name = "上传数据池文件"
	}
	return module.NewModule("/datapools", handler)
}
