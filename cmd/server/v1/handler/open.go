package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type OpenCtrl struct {
	ProjectService *service.ProjectService `inject:""`
	BaseCtrl
}

func (c *OpenCtrl) AllProjectList(ctx iris.Context) {
	username := ctx.URLParam("username")

	data, err := c.ProjectService.AllProjectList(username)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *OpenCtrl) GetProjectsBySpace(ctx iris.Context) {
	spaceCode := ctx.URLParam("spaceCode")

	data, err := c.ProjectService.GetListWithRoleBySpace(spaceCode)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *OpenCtrl) SaveSpaceRelatedProjects(ctx iris.Context) {
	req := v1.SaveSpaceRelatedProjectsReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.ProjectService.SaveSpaceRelatedProjects(req.SpaceCode, req.ProjectShortNames)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *OpenCtrl) CheckProjectAndUser(ctx iris.Context) {
	projectCode := ctx.URLParam("project_code")
	username := ctx.URLParam("username")

	project, userInProject, err := c.ProjectService.CheckProjectAndUserByName(projectCode, username)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
	} else if project.ID == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ErrProjectNotExist.Code, Msg: _domain.ErrProjectNotExist.Msg, Data: project})
	} else if !userInProject {
		ctx.JSON(_domain.Response{Code: _domain.ErrUserNotInProject.Code, Msg: _domain.ErrUserNotInProject.Msg, Data: project})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: project})
	}

	return
}

func (c *OpenCtrl) GetProjectRole(ctx iris.Context) {
	username := ctx.URLParam("username")
	projectCode := ctx.URLParam("project_code")
	if role, err := c.ProjectService.GetProjectRole(username, projectCode); err == nil {
		ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg, Data: role})
	} else {
		ctx.JSON(_domain.Response{Code: _domain.ErrUserNotInProject.Code, Msg: err.Error()})
	}
}
