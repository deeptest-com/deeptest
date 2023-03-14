package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/pkg/message"
	messageDomain "github.com/aaronchen2k/deeptest/pkg/message/Domain"
	messageService "github.com/aaronchen2k/deeptest/pkg/message/service"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type MessageCtrl struct {
	MessageService *service.MessageService `inject:""`
	BaseCtrl
}

var client *message.Client

func init() {
	client = &message.Client{
		Handler: &messageService.MessageServiceV1{},
		Db:      &gorm.DB{},
	}
}

func (c *MessageCtrl) List(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	var req messageDomain.MessageReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.Scope = c.MessageService.GetScope(userId)

	data, err := client.Paginate(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *MessageCtrl) UnreadCount(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	scope := c.MessageService.GetScope(userId)
	req := messageDomain.MessageScope{
		Scope: scope,
	}
	count, err := client.UnreadCount(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"count": count}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

func (c *MessageCtrl) OperateRead(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := messageDomain.MessageReadReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.UserId = userId

	id, err := client.OperateRead(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: 400, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}
