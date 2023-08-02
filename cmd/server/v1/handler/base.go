package handler

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type BaseCtrl struct {
	CommonService *service.CommonService `inject:""`
}

// BatchUpdateField
// @Tags	设计器
// @summary	批量更新字段内容
// @accept	application/json
// @Produce	application/json
// @Param 	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	BatchUpdateReq 	body 	serverDomain.BatchUpdateReq	true 	"批量更新字段内容的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/endpoint/batchUpdateField	[post]
func (c *BaseCtrl) BatchUpdateField(ctx iris.Context) {
	var req serverDomain.BatchUpdateReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	if err := c.CommonService.BatchUpdateField(req); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	//if err := c.EndpointService.BatchUpdateByField(req); err != nil {
	//	ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	//	return
	//}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
