package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

type ServeCtrl struct {
	ServeService *service.ServeService `inject:""`
}

// 项目服务列表
func (c *ServeCtrl) ListByProject(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	projectId, err := ctx.URLParamInt("currProjectId")
	if projectId == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: "projectId"})
		return
	}

	serves, currServe, err := c.ServeService.ListByProject(projectId, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"serves": serves, "currServe": currServe}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret})

	return
}

// Index 服务列表
func (c *ServeCtrl) Index(ctx iris.Context) {
	var req v1.ServeReqPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.Paginate(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// Save 保存服务
func (c *ServeCtrl) Save(ctx iris.Context) {
	var req v1.ServeReq
	if err := ctx.ReadJSON(&req); err == nil {
		//req.CreateUser = "admin"
		req.CreateUser = multi.GetUsername(ctx)
		res, _ := c.ServeService.Save(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// Detail 服务详情
func (c *ServeCtrl) Detail(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res := c.ServeService.GetById(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Copy 克隆服务
func (c *ServeCtrl) Copy(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res := c.ServeService.Copy(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Delete 删除服务
func (c *ServeCtrl) Delete(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// Expire 禁用服务
func (c *ServeCtrl) Expire(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DisableById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

// SaveVersion 保存版本
func (c *ServeCtrl) SaveVersion(ctx iris.Context) {
	var req v1.ServeVersionReq
	if err := ctx.ReadJSON(&req); err == nil {
		if res, err := c.ServeService.SaveVersion(req); err != nil {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// ListVersion 获取版本列表
func (c *ServeCtrl) ListVersion(ctx iris.Context) {
	var req v1.ServeVersionPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.PaginateVersion(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) DeleteVersion(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteVersionById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) ExpireVersion(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DisableVersionById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) SaveSchema(ctx iris.Context) {
	var req v1.ServeSchemaReq
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.SaveSchema(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

// ListSchema 获取版本列表
func (c *ServeCtrl) ListSchema(ctx iris.Context) {
	var req v1.ServeSchemaPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.PaginateSchema(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) DeleteSchema(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.ServeService.DeleteSchemaById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) ListServer(ctx iris.Context) {
	serveId := ctx.URLParamUint64("serveId")
	res, err := c.ServeService.ListServer(uint(serveId))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) SaveServer(ctx iris.Context) {
	var req v1.ServeServer
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.SaveServer(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) ExampleToSchema(ctx iris.Context) {
	var req v1.JsonContent
	if err := ctx.ReadJSON(&req); err == nil {
		res := c.ServeService.Example2Schema(req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) SchemaToExample(ctx iris.Context) {
	var req v1.JsonContent
	if err := ctx.ReadJSON(&req); err == nil {
		res := c.ServeService.Schema2Example(req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) SchemaToYaml(ctx iris.Context) {
	var req v1.JsonContent
	if err := ctx.ReadJSON(&req); err == nil {
		res := c.ServeService.Schema2Yaml(req.Data)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) CopySchema(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res, _ := c.ServeService.CopySchema(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res.ID})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) BindEndpoint(ctx iris.Context) {
	var req v1.ServeVersionBindEndpointReq
	if err := ctx.ReadJSON(&req); err == nil {
		c.ServeService.BindEndpoint(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *ServeCtrl) ChangeServe(ctx iris.Context) {
	userId := multi.GetUserId(ctx)

	req := v1.ProjectReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	currServe, err := c.ServeService.ChangeServe(req.Id, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: currServe, Msg: _domain.NoErr.Msg})
}
