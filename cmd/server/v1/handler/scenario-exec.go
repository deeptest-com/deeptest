package handler

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ScenarioExecCtrl struct {
	ScenarioExecService *service.ScenarioExecService `inject:""`

	BaseCtrl
}

// LoadExecData
func (c *ScenarioExecCtrl) LoadExecData(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")

	data, err := c.ScenarioExecService.LoadExecData(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// LoadExecResult
func (c *ScenarioExecCtrl) LoadExecResult(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")

	data, err := c.ScenarioExecService.LoadExecResult(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitResult
func (c *ScenarioExecCtrl) SubmitResult(ctx iris.Context) {
	scenarioId, err := ctx.Params().GetInt("id")

	result := agentDomain.ScenarioExecResult{}
	err = ctx.ReadJSON(&result)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	report, err := c.ScenarioExecService.SaveReport(scenarioId, result)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report})
}
