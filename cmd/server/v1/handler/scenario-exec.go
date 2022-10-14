package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ScenarioExecCtrl struct {
	ScenarioExecService *service.ExecScenarioService `inject:""`
	BaseCtrl
}

// LoadExecData
func (c *ScenarioExecCtrl) LoadExecData(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")

	data, err := c.ScenarioExecService.LoadExecData(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// loadExecData
func (c *ScenarioExecCtrl) LoadExecResult(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")

	data, err := c.ScenarioExecService.Load(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}
