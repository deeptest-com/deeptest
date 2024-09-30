package handler

import (
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type ScenarioInterfaceCtrl struct {
	DebugInterfaceService    *service.DebugInterfaceService    `inject:""`
	ScenarioInterfaceService *service.ScenarioInterfaceService `inject:""`
	ExtractorService         *service.ExtractorService         `inject:""`
	CheckpointService        *service.CheckpointService        `inject:""`
	BaseCtrl
}

// SaveDebugData
// @Tags	场景模块/场景调试
// @summary	保存场景调试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string				true	"Authentication header"
// @Param 	currProjectId	query	int					true	"当前项目ID"
// @Param 	DebugData		body	domain.DebugData	true	"保存场景调试接口的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/scenarios/interface/saveDebugData	[post]
func (c *ScenarioInterfaceCtrl) SaveDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	_, err = c.ScenarioInterfaceService.SaveDebugData(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// ResetDebugData
// @Tags	场景模块/场景调试
// @summary	重置场景调试接口
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	scenarioProcessorId	query	int		true	"scenarioProcessorId"
// @success	200	{object}	_domain.Response{data=model.Processor}
// @Router	/api/v1/scenarios/interface/resetDebugData	[post]
func (c *ScenarioInterfaceCtrl) ResetDebugData(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	createBy := multi.GetUserId(ctx)
	scenarioProcessorId, err := ctx.URLParamInt("scenarioProcessorId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	newProcessor, err := c.ScenarioInterfaceService.ResetDebugData(tenantId, scenarioProcessorId, createBy)
	if err != nil {
		if err.Error() == "interface is deleted" {
			ctx.JSON(_domain.Response{Code: _domain.ErrImportSourceDeleted.Code, Msg: _domain.ErrImportSourceDeleted.Msg})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: newProcessor})
}
