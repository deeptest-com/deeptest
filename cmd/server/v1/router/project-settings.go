package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type ProjectSettingsModule struct {
	ProjectSettingsCtrl *handler.ProjectSettingsCtrl `inject:""`
}

// Party 注册模块
func (m *ProjectSettingsModule) Party() module.WebModule {
	handler := func(public iris.Party) {
		public.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())

		public.Post("/saveSwaggerSync", m.ProjectSettingsCtrl.SaveSwaggerSync).Name = "保存同步信息"
		public.Get("/swaggerSyncDetail", m.ProjectSettingsCtrl.SwaggerSyncDetail).Name = "获取同步信息"

		public.Post("/saveMock", m.ProjectSettingsCtrl.SaveMock).Name = "保存Mock设置"
		public.Get("/getMock", m.ProjectSettingsCtrl.GetMock).Name = "获取Mock设置"

		public.Post("/savePerformance", m.ProjectSettingsCtrl.SavePerformance).Name = "保存性能测试设置"
		public.Get("/getPerformance", m.ProjectSettingsCtrl.GetPerformance).Name = "获取性能测试设置"
	}

	m.ProjectSettingsCtrl.InitSwaggerCron()
	m.ProjectSettingsCtrl.InitThirdPartySyncCron()
	return module.NewModule("/projectSettings", handler)
}
