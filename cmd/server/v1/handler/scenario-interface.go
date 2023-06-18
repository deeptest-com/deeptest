package handler

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ScenarioInterfaceCtrl struct {
	DebugInterfaceService    *service.DebugInterfaceService    `inject:""`
	ScenarioInterfaceService *service.ScenarioInterfaceService `inject:""`
	ExtractorService         *service.ExtractorService         `inject:""`
	CheckpointService        *service.CheckpointService        `inject:""`
	BaseCtrl
}

func (c *ScenarioInterfaceCtrl) SaveDebugData(ctx iris.Context) {
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.ScenarioInterfaceService.SaveDebugData(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	loadReq := domain.DebugReq{
		EndpointInterfaceId: po.EndpointInterfaceId,
		ScenarioProcessorId: req.ScenarioProcessorId,
		UsedBy:              consts.ScenarioDebug,
	}

	data, err := c.DebugInterfaceService.Load(loadReq)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}
