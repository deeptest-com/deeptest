package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type ScenarioCtrl struct {
	ScenarioService *service.ScenarioService `inject:""`
	BaseCtrl
}

func (c *ScenarioCtrl) ListByProject(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	res, err := c.ScenarioService.ListByProject(projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

func (c *ScenarioCtrl) List(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.ScenarioReqPaginate
	err = ctx.ReadQuery(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.ScenarioService.Paginate(req, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *ScenarioCtrl) Get(ctx iris.Context) {
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	scenario, err := c.ScenarioService.GetById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: scenario, Msg: _domain.NoErr.Msg})
}

func (c *ScenarioCtrl) Create(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := model.Scenario{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)
	req.CreateUserName = multi.GetUsername(ctx)
	req.CreateUserId = multi.GetUserId(ctx)
	req.Status = consts.Draft
	po, err := c.ScenarioService.Create(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po, Msg: _domain.NoErr.Msg})
}

func (c *ScenarioCtrl) Update(ctx iris.Context) {
	var req model.Scenario
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	req.CreateUserName = multi.GetUsername(ctx)
	req.CreateUserId = multi.GetUserId(ctx)
	err = c.ScenarioService.Update(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *ScenarioCtrl) Delete(ctx iris.Context) {
	var req _domain.ReqId
	err := ctx.ReadParams(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ScenarioService.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *ScenarioCtrl) AddPlans(ctx iris.Context) {
	scenarioId, _ := ctx.Params().GetInt("id")

	planIds := make([]int, 0)
	err := ctx.ReadJSON(&planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "ids"})
		return
	}

	err = c.ScenarioService.AddPlans(scenarioId, planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *ScenarioCtrl) Plans(ctx iris.Context) {

	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	scenarioId, _ := ctx.Params().GetInt("id")

	var req serverDomain.ScenarioPlanReqPaginate
	err = ctx.ReadJSON(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	req.ProjectId = uint(projectId)
	data, err := c.ScenarioService.PlanPaginate(req, scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})

}

func (c *ScenarioCtrl) UpdateStatus(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	status := ctx.URLParamDefault("status", "")
	if status == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
	}
	err := c.ScenarioService.UpdateStatus(uint(id), consts.TestStatus(status))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ScenarioCtrl) UpdatePriority(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	priority := ctx.URLParamDefault("priority", "")
	if priority == "" {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
	}
	err := c.ScenarioService.UpdatePriority(uint(id), priority)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ScenarioCtrl) RemovePlans(ctx iris.Context) {
	scenarioId, _ := ctx.Params().GetInt("id")

	planIds := make([]int, 0)
	err := ctx.ReadJSON(&planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "ids"})
		return
	}

	err = c.ScenarioService.RemovePlans(scenarioId, planIds)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}
