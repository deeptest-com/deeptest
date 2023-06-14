package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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
	req := serverDomain.TestInterfaceLoadReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	data, err := c.TestInterfaceService.Load(req.ProjectId, req.ServeId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// Save
func (c *TestInterfaceCtrl) Save(ctx iris.Context) {
	req := serverDomain.TestInterfaceCreateReq{}
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
