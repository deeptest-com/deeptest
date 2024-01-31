package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type SaasCtrl struct {
	SassService *service.SaasService `inject:""`
	BaseCtrl
}

func (c *SaasCtrl) GetUserList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	res, err := c.SassService.GetUserList(tenantId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

}
