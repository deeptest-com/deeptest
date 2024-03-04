package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationService "github.com/aaronchen2k/deeptest/integration/service"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
)

type OpenCtrl struct {
	ProjectService            *service.ProjectService            `inject:""`
	IntegrationProjectService *integrationService.ProjectService `inject:""`
	BaseCtrl
}

func (c *OpenCtrl) AllProjectList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	username := ctx.URLParam("username")

	data, err := c.ProjectService.AllProjectList(tenantId, username)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *OpenCtrl) GetProjectsBySpace(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	spaceCode := ctx.URLParam("spaceCode")
	username := ctx.URLParam("username")

	logUtils.Infof("GetProjectsBySpace tenantId:%+v, spaceCode:%+v, username:%+v", tenantId, spaceCode, username)
	data, err := c.IntegrationProjectService.GetListWithRoleBySpace(tenantId, spaceCode, username)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *OpenCtrl) SaveSpaceRelatedProjects(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := v1.SaveSpaceRelatedProjectsReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.IntegrationProjectService.SaveSpaceRelatedProjects(tenantId, req.SpaceCode, req.ProjectShortNames)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *OpenCtrl) GetProjectRole(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	username := ctx.URLParam("username")
	projectCode := ctx.URLParam("projectCode")
	if role, err := c.ProjectService.GetProjectRole(tenantId, username, projectCode); err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: role})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.ErrUserNotInProject.Code, Msg: err.Error()})
	}
}
