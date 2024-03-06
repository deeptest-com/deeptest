package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointTagCtrl struct {
	BaseCtrl
	EndpointTagService *service.EndpointTagService `inject:""`
}

// ListTags
// @Tags	设计器/标签
// @summary	用例详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=[]model.EndpointTagRel}
// @Router	/api/v1/endpoint/tags	[get]
func (c *EndpointTagCtrl) ListTags(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	tags, err := c.EndpointTagService.ListTagsByProject(tenantId, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: tags, Msg: _domain.NoErr.Msg})
}
