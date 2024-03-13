package handler

import (
	agentExec "github.com/aaronchen2k/deeptest/internal/agent/exec"
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type ScenarioExecCtrl struct {
	ScenarioExecService *service.ScenarioExecService `inject:""`

	BaseCtrl
}

// LoadExecData
// @Tags	场景模块/场景执行
// @summary	加载执行场景
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				query	int		true	"场景ID"
// @Param 	environmentId	query	int		true	"环境ID"
// @success	200	{object}	_domain.Response{data=agentExec.ScenarioExecObjMsg}
// @Router	/api/v1/scenarios/exec/loadExecScenario	[get]
func (c *ScenarioExecCtrl) LoadExecData(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	environmentId, err := ctx.URLParamInt("environmentId")

	data, err := c.ScenarioExecService.LoadExecData(uint(id), uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := agentExec.ScenarioExecObjMsg{}
	copier.CopyWithOption(&ret, data, copier.Option{DeepCopy: true})

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})
}

// LoadExecResult
// @Tags	场景模块/场景执行
// @summary	加载执行结果
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	scenarioId		query	int		true	"场景ID"
// @success	200	{object}	_domain.Response{data=domain.Report}
// @Router	/api/v1/scenarios/exec/loadExecResult	[get]
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
// @Tags	场景模块/场景执行
// @summary	提交测试结果
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	id					path	int								true	"场景ID"
// @Param 	ScenarioExecResult	body	agentDomain.ScenarioExecResult	true	"场景执行结果"
// @success	200	{object}	_domain.Response{data=model.ScenarioReport}
// @Router	/api/v1/scenarios/exec/submitResult/{id}	[post]
func (c *ScenarioExecCtrl) SubmitResult(ctx iris.Context) {
	scenarioId, err := ctx.Params().GetInt("id")

	result := agentExecDomain.ScenarioExecResult{}
	err = ctx.ReadJSON(&result)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	userId := multi.GetUserId(ctx)

	report, err := c.ScenarioExecService.SaveReport(scenarioId, userId, result)

	report.Logs = nil // otherwise will cause an json parse err on agent size
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: report})
}

// GetScenarioNormalData
// @Tags	场景模块/场景执行
// @summary	获取场景执行初始化信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string	true	"Authentication header"
// @Param 	currProjectId		query	int		true	"当前项目ID"
// @Param 	id					query	int		true	"场景ID"
// @Param 	environmentId		query	int		true	"环境ID"
// @success	200	{object}	_domain.Response{data=agentDomain.Report}
// @Router	/api/v1/scenarios/exec/getScenarioNormalData	[get]
func (c *ScenarioExecCtrl) GetScenarioNormalData(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	environmentId, err := ctx.URLParamInt("environmentId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.ScenarioExecService.GetScenarioNormalData(uint(id), uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}
