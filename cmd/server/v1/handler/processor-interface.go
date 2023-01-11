package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ProcessorInterfaceCtrl struct {
	ScenarioInterfaceService *service.ProcessorInterfaceService `inject:""`
	BaseCtrl
}

func (c *ProcessorInterfaceCtrl) GetInterface(ctx iris.Context) {
	interfaceId, err := ctx.URLParamInt("interfaceId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	interf, err := c.ScenarioInterfaceService.GetById(uint(interfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: interf})
}

func (c *ProcessorInterfaceCtrl) ListInvocation(ctx iris.Context) {
	interfaceId, err := ctx.URLParamInt("interfaceId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	data, err := c.ScenarioInterfaceService.ListInvocation(uint(interfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}
