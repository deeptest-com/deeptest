package handler

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type PlanExecCtrl struct {
	PlanExecService *service.PlanExecService `inject:""`

	BaseCtrl
}

// LoadExecData
func (c *PlanExecCtrl) LoadExecData(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")

	data, err := c.PlanExecService.LoadExecData(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// LoadExecResult
func (c *PlanExecCtrl) LoadExecResult(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("planId")

	data, err := c.PlanExecService.LoadExecResult(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitResult
func (c *PlanExecCtrl) SubmitResult(ctx iris.Context) {
	//scenarioId, err := ctx.URLParamInt("id")
	planId, err := ctx.Params().GetInt("id")

	result := map[uint]*agentDomain.PlanExecResult{}
	err = ctx.ReadJSON(&result)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	report, err := c.PlanExecService.SaveReport(planId, result)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report})
}
