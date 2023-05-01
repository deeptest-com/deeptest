package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type MessageCtrl struct {
	MessageService   *service.MessageService   `inject:""`
	WebSocketService *service.WebSocketService `inject:""`
	BaseCtrl
}

func (c *MessageCtrl) List(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	var req serverDomain.MessageReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.MessageService.Paginate(req, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *MessageCtrl) UnreadCount(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	count, err := c.MessageService.UnreadCount(userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"count": count}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

func (c *MessageCtrl) OperateRead(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := serverDomain.MessageReadReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.UserId = userId
	id, err := c.MessageService.OperateRead(req)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}
