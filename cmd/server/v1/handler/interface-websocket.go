package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type WebsocketInterfaceCtrl struct {
	WebsocketInterfaceService *service.WebsocketInterfaceService `inject:""`

	BaseCtrl
}

func (c *WebsocketInterfaceCtrl) GetDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	diagnoseInterfaceId, _ := ctx.URLParamInt("diagnoseInterfaceId")

	data, err := c.WebsocketInterfaceService.GetDebugData(tenantId, diagnoseInterfaceId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *WebsocketInterfaceCtrl) SaveDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)

	req := domain.WebsocketDebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.WebsocketInterfaceService.SaveDebugData(req, tenantId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
