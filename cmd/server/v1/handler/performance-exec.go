package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type PerformanceExecCtrl struct {
	PerformanceExecService *service.PerformanceExecService `inject:""`

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
// @success	200	{object}	_domain.Response{data=agentExec.PerformanceExecObjMsg}
// @Router	/api/v1/scenarios/exec/loadExecPerformance	[get]
func (c *PerformanceExecCtrl) LoadExecData(ctx iris.Context) {
	planId, err := ctx.URLParamInt("planId")
	environmentId, err := ctx.URLParamInt("environmentId")

	data, err := c.PerformanceExecService.LoadExecData(uint(planId), uint(environmentId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}
