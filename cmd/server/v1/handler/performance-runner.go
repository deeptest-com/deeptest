package handler

import (
	agentDomain "github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type PerformanceRunnerCtrl struct {
	PerformanceRunnerService *service.PerformanceRunnerService `inject:""`
	BaseCtrl
}

func (c *PerformanceRunnerCtrl) List(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")
	if scenarioId <= 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.PerformanceRunnerService.List(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *PerformanceRunnerCtrl) Get(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	performanceTestPlan, err := c.PerformanceRunnerService.GetById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: performanceTestPlan, Msg: _domain.NoErr.Msg})
}

func (c *PerformanceRunnerCtrl) Select(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := agentDomain.PerformanceRunnerSelectionReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = projectId

	err = c.PerformanceRunnerService.Select(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: req, Msg: _domain.NoErr.Msg})
}

func (c *PerformanceRunnerCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetUint("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.PerformanceRunnerService.DeleteById(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
