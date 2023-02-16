package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type ServeCtrl struct {
	ServeService *service.ServeService `inject:""`
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

// Clone 克隆服务
func (c *ServeCtrl) Copy(ctx iris.Context) {

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
		res, _ := c.ServeService.SaveVersion(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

// ListVersion 获取版本列表
func (c *ServeCtrl) ListVersion(ctx iris.Context) {
	id := ctx.URLParamUint64("serveId")
	res, err := c.ServeService.ListVersion(uint(id))
	if err == nil {
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

}

/*
func (c *ServeCtrl) SaveServer(ctx iris.Context) {
	var req v1.ServeSeverReq
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.ServeService.SaveVersion(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}
*/

func (c *ServeCtrl) ListServer(ctx iris.Context) {

}
