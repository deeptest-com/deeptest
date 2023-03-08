package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type EnvironmentModule struct {
	EnvironmentCtrl *handler.EnvironmentCtrl `inject:""`
}

// Party 脚本
func (m *EnvironmentModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		//index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Post("/save", m.EnvironmentCtrl.Save).Name = "保存环境"
		index.Get("/list", m.EnvironmentCtrl.ListAll).Name = "环境列表"

		index.Get("/", m.EnvironmentCtrl.List).Name = "环境列表"
		index.Get("/{id:uint}", m.EnvironmentCtrl.Get).Name = "环境详情"
		index.Post("/", m.EnvironmentCtrl.Create).Name = "新建环境"
		index.Put("/", m.EnvironmentCtrl.Update).Name = "更新环境"
		index.Post("/copyEnvironment", m.EnvironmentCtrl.Copy).Name = "复制环境"
		index.Delete("/{id:uint}", m.EnvironmentCtrl.Delete).Name = "删除环境"
		index.Post("/changeEnvironment", m.EnvironmentCtrl.Change).Name = "修改环境"

		index.PartyFunc("/vars", func(party iris.Party) {
			party.Post("/", m.EnvironmentCtrl.CreateVar).Name = "新建环境变量"
			party.Put("/", m.EnvironmentCtrl.UpdateVar).Name = "更新环境变量"
			party.Delete("/{id:uint}", m.EnvironmentCtrl.DeleteVar).Name = "删除环境变量"
			party.Post("/clear", m.EnvironmentCtrl.ClearVar).Name = "清空环境变量"
			party.Post("/global", m.EnvironmentCtrl.SaveGlobal).Name = "保存全局变量"
			party.Get("/global", m.EnvironmentCtrl.ListGlobal).Name = "保存全局变量"
		})

		index.PartyFunc("/shareVars", func(party iris.Party) {
			party.Delete("/{id:uint}", m.EnvironmentCtrl.DeleteShareVar).Name = "删除共享变量"
			party.Post("/clear", m.EnvironmentCtrl.ClearShareVar).Name = "清空共享变量"
		})
	}

	return module.NewModule("/environments", handler)
}
