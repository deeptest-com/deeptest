package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	service "github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type TestInterfaceCtrl struct {
	TestInterfaceService *service.TestInterfaceService `inject:""`
	ExtractorService     *service.ExtractorService     `inject:""`
	CheckpointService    *service.CheckpointService    `inject:""`
	BaseCtrl
}

// Load
func (c *TestInterfaceCtrl) Load(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("projectId")
	serveId, _ := ctx.URLParamInt("serveId")

	data, err := c.TestInterfaceService.Load(projectId, serveId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
func (c *TestInterfaceCtrl) Save(ctx iris.Context) {
	req := serverDomain.TestInterfaceSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.TestInterfaceService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.TestInterfaceService.Load(int(po.ProjectId), int(po.ServeId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

// Update
func (c *TestInterfaceCtrl) Update(ctx iris.Context) {
	req := serverDomain.TestInterfaceSaveReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	po, err := c.TestInterfaceService.Save(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.TestInterfaceService.Load(int(po.ProjectId), int(po.ServeId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data})
}

func (c *TestInterfaceCtrl) Delete(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	typ := ctx.URLParam("type")

	err := c.TestInterfaceService.Remove(id, serverConsts.TestInterfaceType(typ))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}
