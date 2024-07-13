package router

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/handler"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/module"
	"github.com/kataras/iris/v12"
)

type AiChatModule struct {
	AiChatCtrl *handler.AiChatCtrl `inject:""`
}

// Party chatchat
func (m *AiChatModule) Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Post("/knowledge_base_chat", m.AiChatCtrl.KnowledgeBaseChat).Name = "与知识库对话"

		index.Post("/list_valid_models", m.AiChatCtrl.ListValidModel).Name = "列出大模型"
		index.Get("/list_knowledge_bases", m.AiChatCtrl.ListKnowledgeBase).Name = "列出知识库"
	}
	return module.NewModule("/aichat", handler)
}
