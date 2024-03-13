package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type SummaryModule struct {
	SummaryCtrl *handler.SummaryCtrl `inject:""`
	Cron        *cron.ServerCron     `inject:""`
}

// Party 用户
func (m *SummaryModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin(), middleware.ProjectPerm())
		index.Get("/card/{projectId:uint}", m.SummaryCtrl.Card).Name = "汇总卡片位信息"
		index.Get("/bugs/{projectId:uint}", m.SummaryCtrl.Bugs).Name = "汇总bug信息"
		index.Get("/details", m.SummaryCtrl.Details).Name = "汇总项目详情"
		index.Get("/projectUserRanking/{cycle:uint}/{projectId:uint}", m.SummaryCtrl.ProjectUserRanking).Name = "汇总项目用户排行数据"
		index.Get("/collection/{store:string}", m.SummaryCtrl.Collection).Name = "汇总数据"
	}

	//SAAS
	m.Cron.AddTask("summary", consts.SummaryDataCheckInterval, func() {
		m.SummaryCtrl.Summary()
	})
	return module.NewModule("/summary", handler)
}
