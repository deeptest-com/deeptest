package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"

	"github.com/kataras/iris/v12"
)

type ScenarioNodeCtrl struct {
	ScenarioNodeService *service.ScenarioNodeService `inject:""`
	ScenarioService     *service.ScenarioService     `inject:""`
	BaseCtrl
}

// LoadTree
func (c *ScenarioNodeCtrl) LoadTree(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")

	scenario, err := c.ScenarioService.GetById(uint(scenarioId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.ScenarioNodeService.GetTree(scenario)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// AddInterfaces 添加
func (c *ScenarioNodeCtrl) AddInterfaces(ctx iris.Context) {
	req := v1.ScenarioAddInterfacesReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	nodePo, bizErr := c.ScenarioNodeService.AddInterfaces(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// AddProcessor 添加
func (c *ScenarioNodeCtrl) AddProcessor(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req := v1.ScenarioAddScenarioReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.ProjectId = uint(projectId)

	nodePo, bizErr := c.ScenarioNodeService.AddProcessor(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nodePo})
}

// UpdateName 更新
func (c *ScenarioNodeCtrl) UpdateName(ctx iris.Context) {
	var req v1.ScenarioNodeReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioNodeService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Delete 删除
func (c *ScenarioNodeCtrl) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	err = c.ScenarioNodeService.Delete(uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Mode 移动
func (c *ScenarioNodeCtrl) Move(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req v1.ScenarioNodeMoveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	_, err = c.ScenarioNodeService.Move(uint(req.DragKey), uint(req.DropKey), req.DropPos, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
