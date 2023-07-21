package handler

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointTagCtrl struct {
	EndpointTagService *service.EndpointTagService `inject:""`
}

func (c *EndpointTagCtrl) ListTags(ctx iris.Context) {
	projectId, _ := ctx.URLParamInt("currProjectId")

	tags, err := c.EndpointTagService.ListTagsByProject(uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: tags, Msg: _domain.NoErr.Msg})
}
