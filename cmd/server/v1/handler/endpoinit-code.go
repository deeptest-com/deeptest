package handler

import (
	domain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointCodeCtrl struct {
	BaseCtrl
	EndpointCodeService *service.EndpointCodeService `inject:""`
}

func (c *EndpointCodeCtrl) Index(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req domain.EndpointCodeReq

	if err := ctx.ReadJSON(&req); err == nil {

		res := c.EndpointCodeService.Generate(tenantId, req.LangType, req.NameRule, req.ProjectId, req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}

	return
}
