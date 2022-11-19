package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/agent/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/agent/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type SpecCtrl struct {
	SpecService        *service.SpecService        `inject:""`
	ProjectService     *service.ProjectService     `inject:""`
	InterfaceService   *service.InterfaceService   `inject:""`
	EnvironmentService *service.EnvironmentService `inject:""`
}

// Load 解析定义文件
func (c *SpecCtrl) Load(ctx iris.Context) {
	req := domain.SubmitSpecReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	doc3, info, err := c.SpecService.Load(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	project, err := c.ProjectService.CreateOrGetBySpec(req.File, req.Url)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	err = c.InterfaceService.Generate(doc3, project.ID)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	err = c.EnvironmentService.Generate(doc3, project.ID)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ret := iris.Map{"doc": doc3, "info": info}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}
