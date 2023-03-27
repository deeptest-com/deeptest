package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ProcessorInterfaceCtrl struct {
	ProcessorInterfaceService *service.ProcessorInterfaceService `inject:""`
	BaseCtrl
}

func (c *ProcessorInterfaceCtrl) GetInterface(ctx iris.Context) {
	interfaceId, err := ctx.URLParamInt("interfaceId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	interf, err := c.ProcessorInterfaceService.GetById(uint(interfaceId))
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

	data, err := c.ProcessorInterfaceService.ListInvocation(uint(interfaceId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// SaveInterface
func (c *ProcessorInterfaceCtrl) SaveInterface(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	req := domain.DebugRequest{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	//req.ProjectId = uint(projectId)
	err = c.ProcessorInterfaceService.UpdateByConfig(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *ProcessorInterfaceCtrl) Index(ctx iris.Context) {
	currProjectId, err := ctx.URLParamInt("currProjectId")
	scenarioId := ctx.URLParamIntDefault("scenarioId", 0)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	res, err := c.ProcessorInterfaceService.GetList(uint(currProjectId), uint(scenarioId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
}
