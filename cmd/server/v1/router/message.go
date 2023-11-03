package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/aaronchen2k/deeptest/internal/server/middleware"
	"github.com/kataras/iris/v12"
)

type MessageModule struct {
	MessageCtrl *handler.MessageCtrl `inject:""`
}

// Party 消息
func (m *MessageModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/receiveMcsApprovalData", m.MessageCtrl.ReceiveMcsApprovalData).Name = "接收mcs审批数据"

		index.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())

		index.Get("/", m.MessageCtrl.List).Name = "消息列表"
		index.Get("/unreadCount", m.MessageCtrl.UnreadCount).Name = "未读消息数"
		index.Post("/operateRead", m.MessageCtrl.OperateRead).Name = "已读操作"
	}

	m.MessageCtrl.InitThirdPartySyncCron()
	return module.NewModule("/message", handler)
}
