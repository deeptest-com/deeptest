package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProjectCronModule struct {
	ProjectCronCtrl *handler.ProjectCronCtrl `inject:""`
}

func (m *ProjectCronModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		index.Get("/", m.ProjectCronCtrl.List).Name = "定时任务列表"
		index.Get("/{id:uint}", m.ProjectCronCtrl.Get).Name = "定时任务详情"
		index.Post("/", m.ProjectCronCtrl.Save).Name = "新建定时任务"
		index.Put("/", m.ProjectCronCtrl.Update).Name = "更新定时任务"
		index.Delete("/{id:uint}", m.ProjectCronCtrl.Delete).Name = "删除定时任务"
		index.Get("/{id:uint}/clone", m.ProjectCronCtrl.Clone).Name = "克隆定时任务"
		index.Get("/engineeringOptions", m.ProjectCronCtrl.EngineeringOptions).Name = "获取工程下拉选项"
		index.Get("/serviceOptions", m.ProjectCronCtrl.ServiceOptions).Name = "获取服务下拉选项"
		index.Get("/allServiceList", m.ProjectCronCtrl.AllServiceList).Name = "获取所有服务列表"

	}
	return module.NewModule("/project/cron", handler)

}
