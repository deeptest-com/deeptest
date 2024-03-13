package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ResponseDefineCtrl struct {
	BaseCtrl
	ResponseDefineService *service.ResponseDefineService `inject:""`
}

func (c *ResponseDefineCtrl) Update(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req model.DebugConditionResponseDefine
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ResponseDefineService.Update(tenantId, req.ID, req.Disabled, req.Code)
	if err != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})

	/*

		err = c.PostConditionService.CreateExpression(&condition)
		if err != nil {
			ctx.JSON(_domain.Response{
				Code: _domain.SystemErr.Code,
			})
			return
		}

		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: condition, Msg: _domain.NoErr.Msg})

	*/
}
