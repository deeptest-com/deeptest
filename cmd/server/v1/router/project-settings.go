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
		public.Get("/swaggerSyncDetail", m.ProjectSettingsCtrl.SwaggerSyncDetail).Name = "保存同步信息"

		public.Post("/saveMock", m.ProjectSettingsCtrl.SaveMock).Name = "保存同步信息"
		public.Get("/getMock", m.ProjectSettingsCtrl.GetMock).Name = "保存同步信息"
	}
	//SAAS 增加定义初始化定时任务，防止租户服务请求失败，导致初始化失败
	m.ProjectSettingsCtrl.InitSwaggerCron()
	m.ProjectSettingsCtrl.InitThirdPartySyncCron()
	return module.NewModule("/projectSettings", handler)
}
