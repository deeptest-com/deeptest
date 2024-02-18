package handler

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationService "github.com/aaronchen2k/deeptest/integration/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
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
	MessageService            *service.MessageService            `inject:""`
	WebSocketService          *service.WebSocketService          `inject:""`
	IntegrationMessageService *integrationService.MessageService `inject:""`
	BaseCtrl
}

// List
// @Tags	消息管理
// @summary	消息列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	MessageReqPaginate	query	serverDomain.MessageReqPaginate	true	"获取消息列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.Message}}
// @Router	/api/v1/message	[get]
func (c *MessageCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
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

	data, err := c.MessageService.Paginate(tenantId, req, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// UnreadCount
// @Tags	消息管理
// @summary	未读消息数
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=object{count=int}}
// @Router	/api/v1/message/unreadCount	[get]
func (c *MessageCtrl) UnreadCount(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)

	count, err := c.MessageService.UnreadCount(tenantId, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"count": count}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

// OperateRead
// @Tags	消息管理
// @summary	已读操作
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string						true	"Authentication header"
// @Param 	currProjectId		query	int							true	"当前项目ID"
// @Param 	MessageReadReq		body	serverDomain.MessageReadReq	true	"已读操作的请求参数"
// @success	200	{object}	_domain.Response{data=object{id=int}}
// @Router	/api/v1/message/operateRead	[post]
func (c *MessageCtrl) OperateRead(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)

	req := serverDomain.MessageReadReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.UserId = userId
	id, err := c.MessageService.OperateRead(tenantId, req)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

func (c *MessageCtrl) ReceiveMcsApprovalData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.McsApprovalRes{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	reqData := serverDomain.McsApprovalResData{}
	_ = json.Unmarshal([]byte(req.Data), &reqData)

	logUtils.Infof("ReceiveMcsApprovalData Req:%+v", req)
	err = c.MessageService.ReceiveMcsApprovalResult(tenantId, reqData)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})

}

func (c *MessageCtrl) InitThirdPartySyncCron(ctx iris.Context) {
	//SAAS
	tenantId := c.getTenantId(ctx)
	if !config.CONFIG.Mcs.Switch {
		return
	}

	c.IntegrationMessageService.SendMessageCron(tenantId)
}
