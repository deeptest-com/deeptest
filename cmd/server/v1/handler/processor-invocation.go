package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ProcessorInvocationCtrl struct {
	ScenarioInvocationService *service.ProcessorInvocationService `inject:""`
	InterfaceService          *service.InterfaceService           `inject:""`
	ExtractorService          *service.ExtractorService           `inject:""`
	CheckpointService         *service.CheckpointService          `inject:""`
	BaseCtrl
}

// LoadInterfaceExecData
func (c *ProcessorInvocationCtrl) LoadInterfaceExecData(ctx iris.Context) {
	req := domain.InvocationRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.ScenarioInvocationService.LoadInterfaceExecData(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitInterfaceInvokeResult
func (c *ProcessorInvocationCtrl) SubmitInterfaceInvokeResult(ctx iris.Context) {
	req := domain.SubmitInvocationResultRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.ScenarioInvocationService.SubmitInterfaceInvokeResult(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}
