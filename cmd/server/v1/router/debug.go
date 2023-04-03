package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type DebugModule struct {
	DebugCtrl *handler.DebugCtrl `inject:""`
}

// Party 脚本
func (m *DebugModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		//index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Post("/loadData", m.DebugCtrl.LoadData).Name = "获取接口执行数据"
		index.Post("/submitResult", m.DebugCtrl.SubmitResult).Name = "提交接口执行结果"

		index.Get("/", m.DebugCtrl.List).Name = "调试列表"
		index.Get("/{id:uint}", m.DebugCtrl.Get).Name = "调试详情"
		index.Delete("/{id:uint}", m.DebugCtrl.Delete).Name = "删除调试记录"

		index.Get("/getLastResp", m.DebugCtrl.GetLastResp).Name = "最后一次调试响应"
	}
	return module.NewModule("/debugs", handler)
}
