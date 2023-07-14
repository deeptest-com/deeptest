package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	encoder "github.com/zwgblue/yaml-encoder"
)

type EndpointCtrl struct {
	EndpointService *service.EndpointService `inject:""`
	ServeService    *service.ServeService    `inject:""`
}

func (c *EndpointCtrl) Index(ctx iris.Context) {
	var req serverDomain.EndpointReqPaginate
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	res, _ := c.EndpointService.Paginate(req)
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})

	return
}

func (c *EndpointCtrl) Save(ctx iris.Context) {
	var req serverDomain.EndpointReq
	err := ctx.ReadJSON(&req)

	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}

	req.CreateUser = multi.GetUsername(ctx)
	endpoint := c.requestParser(req)

	/*
		if endpoint.CategoryId == 0 {
			endpoint.CategoryId = 0
		}
	*/

	if res, err := c.EndpointService.Save(endpoint); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	}

	return
}

func (c *EndpointCtrl) Detail(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	version := ctx.URLParamDefault("version", c.EndpointService.GetLatestVersion(uint(id)))
	if id != 0 {
		res := c.EndpointService.GetById(uint(id), version)
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

func (c *EndpointCtrl) BatchDelete(ctx iris.Context) {
	var req []uint
	if err := ctx.ReadJSON(&req); err == nil {
		c.EndpointService.BatchDelete(req)
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

//构造参数构造auth，BasicAuth,BearerToken,OAuth20,ApiKey
func (c *EndpointCtrl) requestParser(req serverDomain.EndpointReq) (endpoint model.Endpoint) {

	for key, item := range req.Interfaces {
		req.Interfaces[key].Body = item.RequestBody.Examples
		req.Interfaces[key].BodyType = item.RequestBody.MediaType
		req.Interfaces[key].Name = req.Title
		/*
					if req.Interfaces[key].RequestBody.Examples == "" {
						var examples []map[string]string
			//			example := c.ServeService.Schema2Example(req.ServeId, item.RequestBody.SchemaItem.Content)
			//			examples = append(examples, map[string]string{"name": "default", "content": commonUtils.JsonEncode(example)})
			//			req.Interfaces[key].RequestBody.Examples = commonUtils.JsonEncode(examples)
					}
		*/

	}

	if req.CategoryId == 0 {
		req.CategoryId = -1
	}

	if req.Status == 0 {
		req.Status = 1
	}

	copier.CopyWithOption(&endpoint, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
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

func (c *EndpointCtrl) Publish(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.Publish(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) Develop(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	err := c.EndpointService.Develop(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) Copy(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	version := ctx.URLParamDefault("version", c.EndpointService.GetLatestVersion(uint(id)))
	res, err := c.EndpointService.Copy(uint(id), version)
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) Yaml(ctx iris.Context) {
	var req serverDomain.EndpointReq
	if err := ctx.ReadJSON(&req); err == nil {
		endpoint := c.requestParser(req)
		res := c.EndpointService.Yaml(endpoint)
		var ret interface{}
		commonUtils.JsonDecode(commonUtils.JsonEncode(res), &ret)
		content, _ := encoder.NewEncoder(ret).Encode()
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: string(content)})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
	return
}

func (c *EndpointCtrl) UpdateStatus(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	status := ctx.URLParamUint64("status")
	err := c.EndpointService.UpdateStatus(uint(id), int64(status))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
	}
}

func (c *EndpointCtrl) AddVersion(ctx iris.Context) {
	var req serverDomain.EndpointVersionReq
	if err := ctx.ReadJSON(&req); err == nil {
		var version model.EndpointVersion
		copier.CopyWithOption(&version, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
		err = c.EndpointService.AddVersion(&version)
		if err == nil {
			ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: req.Version})
		} else {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		}
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EndpointCtrl) ListVersions(ctx iris.Context) {
	id := ctx.URLParamUint64("id")
	res, err := c.EndpointService.GetVersionsByEndpointId(uint(id))
	if err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	}
}

func (c *EndpointCtrl) BatchUpdateField(ctx iris.Context) {
	var req serverDomain.BatchUpdateReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if err := c.EndpointService.BatchUpdateByField(req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
	return
}

/*
func (c *EndpointCtrl) Index() {
	c.EndpointService.GetVersionsByEndpointId(1)
}
*/
