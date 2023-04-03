package handler


type MessageCtrl struct {
	MessageService   *service.MessageService   `inject:""`
	WebSocketService *service.WebSocketService `inject:""`
	BaseCtrl
}

func (c *MessageCtrl) List(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	var req v1.MessageReqPaginate
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

	c.WebSocketService.SendMsg(
		consts.WsDefaultNameSpace,
		consts.WsMessageEvent,
		ret)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *MessageCtrl) OperateRead(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := v1.MessageReadReq{}
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
