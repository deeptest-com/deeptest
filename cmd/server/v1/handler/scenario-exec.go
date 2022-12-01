package handler

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type ScenarioExecCtrl struct {
	ScenarioExecService *service.ExecScenarioService `inject:""`

	BaseCtrl
}

// LoadExecData
func (c *ScenarioExecCtrl) LoadExecData(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")

	data, err := c.ScenarioExecService.LoadExecData(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// loadExecScenario
func (c *ScenarioExecCtrl) LoadExecResult(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")

	data, err := c.ScenarioExecService.Load(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// SubmitResult
func (c *ScenarioExecCtrl) SubmitResult(ctx iris.Context) {
	//scenarioId, err := ctx.URLParamInt("id")
	scenarioId, err := ctx.Params().GetInt("id")

	result := domain.Result{}
	err = ctx.ReadJSON(&result)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	c.ScenarioExecService.SaveReport(scenarioId, result)

	logUtils.Infof("%v", scenarioId)

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
