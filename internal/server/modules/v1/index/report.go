package index

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/controller"
	"github.com/kataras/iris/v12"
)

type ReportModule struct {
	ReportCtrl *controller.ReportCtrl `inject:""`
}

func NewReportModule() *ReportModule {
	return &ReportModule{}
}

// Party 场景
func (m *ReportModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Get("/", m.ReportCtrl.List).Name = "结果列表"
		index.Get("/{id:uint}", m.ReportCtrl.Get).Name = "结果详情"
		index.Delete("/{id:uint}", m.ReportCtrl.Delete).Name = "删除场景"
	}

	return module.NewModule("/reports", handler)
}
