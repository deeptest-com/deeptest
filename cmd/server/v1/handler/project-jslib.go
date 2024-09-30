package handler

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/service"
	"github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type JslibCtrl struct {
	JslibService *service.JslibService `inject:""`
	BaseCtrl
}

func (c *JslibCtrl) List(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	keywords := ctx.URLParam("keywords")

	res, err := c.JslibService.List(tenantId, keywords, projectId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
}

func (c *JslibCtrl) Get(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	datapool, err := c.JslibService.Get(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: datapool, Msg: _domain.NoErr.Msg})
}

func (c *JslibCtrl) Save(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	req := model.SysJslib{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)

	userName := multi.GetUsername(ctx)
	if req.ID > 0 {
		req.UpdateUser = userName
	} else {
		req.CreateUser = userName
	}

	err = c.JslibService.Save(tenantId, &req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrNameExist.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *JslibCtrl) UpdateName(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")

	req := v1.JslibReq{}
	err = ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	req.ProjectId = uint(projectId)
	req.UpdateUser = multi.GetUsername(ctx)

	err = c.JslibService.UpdateName(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ErrNameExist.Code, Data: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code})
}

func (c *JslibCtrl) Delete(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.JslibService.Delete(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *JslibCtrl) Disable(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.JslibService.Disable(tenantId, uint(id))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
