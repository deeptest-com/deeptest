package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/kataras/iris/v12"
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
		c.requestParser(&req)
		res, _ := c.EndpointService.Save(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

func (c *EndpointCtrl) Delete(ctx iris.Context) {

}

//构造参数构造auth，BasicAuth,BearerToken,OAuth20,ApiKey
func (c *EndpointCtrl) requestParser(req *v1.EndpointReq) {
	for key, item := range req.Interfaces {
		req.Interfaces[key].RequestBody.SchemaItem.Content = _commUtils.JsonEncode(item.RequestBody.SchemaItem.Content)
	}
}
