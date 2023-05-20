package handler

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type PlanExecCtrl struct {
	PlanExecService *service.PlanExecService `inject:""`

	BaseCtrl
}

// LoadExecData
func (c *PlanExecCtrl) LoadExecData(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	environmentId, err := ctx.URLParamInt("environmentId")

	data, err := c.PlanExecService.LoadExecData(id, environmentId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// LoadExecResult
func (c *PlanExecCtrl) LoadExecResult(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("planId")

	data, err := c.PlanExecService.LoadExecResult(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitResult
func (c *PlanExecCtrl) SubmitResult(ctx iris.Context) {
	planId, err := ctx.Params().GetInt("id")

	result := agentDomain.PlanExecResult{}
	err = ctx.ReadJSON(&result)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	userId := multi.GetUserId(ctx)
	report, err := c.PlanExecService.SaveReport(planId, userId, result)

	// report.Logs = nil // otherwise will cause an json parse err on agent size
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report})
}

func (c *PlanExecCtrl) GetPlanReportNormalData(ctx iris.Context) {
	planId, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	environmentId, err := ctx.URLParamInt("environmentId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.PlanExecService.GetPlanReportNormalData(uint(planId), uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}
