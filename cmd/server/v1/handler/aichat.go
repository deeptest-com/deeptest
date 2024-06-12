package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type AiChatCtrl struct {
	BaseCtrl
	AiChatService *service.AiChatService `inject:""`
}

func (c *AiChatCtrl) ListValidModel(ctx iris.Context) {
	req := v1.ChatChatModelReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.AiChatService.ListValidModel(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *AiChatCtrl) ListKnowledgeBase(ctx iris.Context) {
	data, err := c.AiChatService.ListKnowledgeBase()
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *AiChatCtrl) KnowledgeBaseChat(ctx iris.Context) {
	flusher, ok := ctx.ResponseWriter().Flusher()
	if !ok {
		ctx.StopWithText(iris.StatusHTTPVersionNotSupported, "Streaming unsupported!")
		return
	}

	ctx.ContentType("text/event-stream")
	//ctx.Header("content-type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")

	req := v1.KnowledgeBaseChatReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	c.AiChatService.KnowledgeBaseChat(req, flusher, ctx)
}
