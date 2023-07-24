package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type DiagnoseInterfaceCtrl struct {
	DiagnoseInterfaceService *service.DiagnoseInterfaceService `inject:""`
	ExtractorService         *service.ExtractorService         `inject:""`
	CheckpointService        *service.CheckpointService        `inject:""`

	DebugInterfaceService *service.DebugInterfaceService `inject:""`

	BaseCtrl
}

// Load
func (c *DiagnoseInterfaceCtrl) Load(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("projectId")
	serveId, _ := ctx.URLParamInt("serveId")

	data, err := c.DiagnoseInterfaceService.Load(projectId, serveId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Get
func (c *DiagnoseInterfaceCtrl) Get(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")

	data, err := c.DiagnoseInterfaceService.Get(id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
func (c *DiagnoseInterfaceCtrl) Save(ctx iris.Context) {
	req := serverDomain.DiagnoseInterfaceSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.DiagnoseInterfaceService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.DiagnoseInterfaceService.Load(int(po.ProjectId), int(po.ServeId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// SaveDebugData
func (c *DiagnoseInterfaceCtrl) SaveDebugData(ctx iris.Context) {
	req := domain.DebugData{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	_, err = c.DiagnoseInterfaceService.SaveDebugData(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	loadReq := domain.DebugInfo{
		DiagnoseInterfaceId: req.DiagnoseInterfaceId,
		DebugInterfaceId:    req.DebugInterfaceId,
		UsedBy:              consts.DiagnoseDebug,
	}

	data, err := c.DebugInterfaceService.Load(loadReq)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Update
func (c *DiagnoseInterfaceCtrl) Update(ctx iris.Context) {
	req := serverDomain.DiagnoseInterfaceSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.DiagnoseInterfaceService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.DiagnoseInterfaceService.Load(int(po.ProjectId), int(po.ServeId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *DiagnoseInterfaceCtrl) Delete(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	typ := ctx.URLParam("type")

	err := c.DiagnoseInterfaceService.Remove(id, serverConsts.DiagnoseInterfaceType(typ))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

// Mode 移动
func (c *DiagnoseInterfaceCtrl) Move(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	var req serverDomain.DiagnoseInterfaceMoveReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	_, err = c.DiagnoseInterfaceService.Move(uint(req.DragKey), uint(req.DropKey), req.DropPos, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ImportInterfaces 导入接口
func (c *DiagnoseInterfaceCtrl) ImportInterfaces(ctx iris.Context) {
	req := serverDomain.DiagnoseInterfaceImportReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	newNode, bizErr := c.DiagnoseInterfaceService.ImportInterfaces(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: newNode})
}

// ImportCurl 导入cURL命令
func (c *DiagnoseInterfaceCtrl) ImportCurl(ctx iris.Context) {
	req := serverDomain.DiagnoseCurlImportReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	req.CreateBy = multi.GetUserId(ctx)
	newNode, bizErr := c.DiagnoseInterfaceService.ImportCurl(req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{
			Code: _domain.SystemErr.Code,
			Msg:  bizErr.Error(),
		})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: newNode})
}
