package handler

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	encoder "github.com/zwgblue/yaml-encoder"
)

type EndpointCtrl struct {
	EndpointService *service.EndpointService `inject:""`
}

func (c *EndpointCtrl) Index(ctx iris.Context) {
	var req v1.EndpointReqPaginate
	if err := ctx.ReadJSON(&req); err == nil {
		res, _ := c.EndpointService.Paginate(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

func (c *EndpointCtrl) Save(ctx iris.Context) {
	var req v1.EndpointReq
	if err := ctx.ReadJSON(&req); err == nil {
		//req.CreateUser = multi.GetUsername(ctx)
		req.CreateUser = "admin"
		endpoint := c.requestParser(&req)
		res, _ := c.EndpointService.Save(endpoint)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

func (c *EndpointCtrl) Detail(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	if id != 0 {
		res := c.EndpointService.GetById(uint(id))
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) Delete(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.DeleteById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

//构造参数构造auth，BasicAuth,BearerToken,OAuth20,ApiKey
func (c *EndpointCtrl) requestParser(req *v1.EndpointReq) (endpoint model.Endpoint) {
	for _, item := range req.Interfaces {
		fmt.Println(_commUtils.JsonEncode(item.ResponseBodies))
		//req.Interfaces[key].RequestBody.SchemaItem.Content = _commUtils.JsonEncode(item.RequestBody.SchemaItem.Content)
	}
	copier.CopyWithOption(&endpoint, req, copier.Option{DeepCopy: true})
	return
}

func (c *EndpointCtrl) Expire(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.DisableById(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) Copy(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	res, err := c.EndpointService.Copy(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) Yaml(ctx iris.Context) {
	var req v1.EndpointReq
	if err := ctx.ReadJSON(&req); err == nil {
		endpoint := c.requestParser(&req)
		res := c.EndpointService.Yaml(endpoint)
		content, _ := encoder.NewEncoder(res).Encode()
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: string(content)})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}
