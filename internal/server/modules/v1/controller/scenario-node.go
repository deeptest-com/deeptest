package controller

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"

	"github.com/kataras/iris/v12"
)

type ScenarioNodeCtrl struct {
	ScenarioNodeService *service.ScenarioNodeService `inject:""`
	BaseCtrl
}

// LoadTree
func (c *ScenarioNodeCtrl) LoadTree(ctx iris.Context) {
	scenarioId, err := ctx.URLParamInt("scenarioId")

	data, err := c.ScenarioNodeService.GetTree(scenarioId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// AddInterfaces 添加
func (c *ScenarioNodeCtrl) AddInterfaces(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req := serverDomain.ScenarioAddInterfacesReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = projectId

	bizErr := c.ScenarioNodeService.AddInterfaces(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.ErrComm.Code,
			Data: nil,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// AddProcessor 添加
func (c *ScenarioNodeCtrl) AddProcessor(ctx iris.Context) {
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req := serverDomain.ScenarioAddScenarioReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = projectId

	po, bizErr := c.ScenarioNodeService.AddProcessor(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.ErrComm.Code,
			Data: nil,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: po})
}

// Get 详情
func (c *ScenarioNodeCtrl) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: _domain.ParamErr.Msg})
		return
	}

	interf, err := c.ScenarioNodeService.Get(id)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: interf})
}

//// Create 添加
//func (c *ScenarioNodeCtrl) Create(ctx iris.Context) {
//	projectId, err := ctx.URLParamInt("currProjectId")
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
//		return
//	}
//
//	req := serverDomain.ScenarioReq{}
//	err = ctx.ReadJSON(&req)
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
//		return
//	}
//
//	req.ProjectId = projectId
//	intf, err := c.ScenarioNodeService.Create(req)
//	if err != nil {
//		ctx.JSON(_domain.Response{
//			Code: c.ErrCode(err),
//			Data: nil,
//		})
//		return
//	}
//
//	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: intf, Msg: _domain.NoErr.Msg})
//}
//
//// Update 更新
//func (c *ScenarioNodeCtrl) Update(ctx iris.Context) {
//	id, err := ctx.URLParamInt("id")
//
//	var req serverDomain.ScenarioReq
//	if err := ctx.ReadJSON(&req); err != nil {
//		errs := validate.ValidRequest(err)
//		if len(errs) > 0 {
//			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
//			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
//			return
//		}
//	}
//
//	err = c.ScenarioNodeService.Update(id, req)
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
//		return
//	}
//	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
//}

// UpdateName 更新
func (c *ScenarioNodeCtrl) UpdateName(ctx iris.Context) {
	var req serverDomain.ScenarioNodeReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		logUtils.Errorf("参数验证失败", err.Error())
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	err = c.ScenarioNodeService.UpdateName(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: nil, Msg: _domain.NoErr.Msg})
}

//// Delete 删除
//func (c *ScenarioNodeCtrl) Delete(ctx iris.Context) {
//	projectId, _ := ctx.URLParamInt("currProjectId")
//
//	id, err := ctx.Params().GetInt("id")
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
//		return
//	}
//
//	err = c.ScenarioNodeService.Delete(uint(projectId), uint(id))
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
//		return
//	}
//
//	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
//}
//
//// Mode 移动
//func (c *ScenarioNodeCtrl) Move(ctx iris.Context) {
//	projectId, _ := ctx.URLParamInt("currProjectId")
//
//	var req serverDomain.ScenarioMoveReq
//	err := ctx.ReadJSON(&req)
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
//		return
//	}
//
//	_, err = c.ScenarioNodeService.Move(uint(req.DragKey), uint(req.DropKey), req.DropPos, uint(projectId))
//	if err != nil {
//		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
//		return
//	}
//
//	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
//}
