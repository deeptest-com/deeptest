package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

type EndpointFavoriteCtrl struct {
	BaseCtrl
	EndpointService *service.EndpointService `inject:""`
}

func (c *EndpointFavoriteCtrl) Favorite(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.FavoriteReq
	if err := ctx.ReadJSON(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	userId := multi.GetUserId(ctx)
	data, err := c.EndpointService.Favorite(tenantId, req.Id, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: data})
}

func (c *EndpointFavoriteCtrl) Index(ctx iris.Context) {

	tenantId := c.getTenantId(ctx)
	var req serverDomain.FavoriteReq
	if err := ctx.ReadJSON(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	userId := multi.GetUserId(ctx)
	data, err := c.EndpointService.FavoriteList(tenantId, req.ProjectId, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrThirdPartyFunctions.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: data})
}
