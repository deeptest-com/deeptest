package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"strings"
)

type ProjectRolePermCtrl struct {
	ProjectRolePermService *service.ProjectRolePermService `inject:""`
	BaseCtrl
}

// GetProjectUserRole
// @Tags	项目权限
// @summary	获取项目中用户的角色
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=model.ProjectRole}}
// @Router	/api/v1/projects/perms/userRole	[get]
func (c *ProjectRolePermCtrl) GetProjectUserRole(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)
	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	data, err := c.ProjectRolePermService.GetProjectUserRole(tenantId, userId, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// AllRoleList
// @Tags	项目权限
// @summary	所有项目角色列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ProjectRole}}
// @Router	/api/v1/projects/perms/rolesList	[get]
func (c *ProjectRolePermCtrl) AllRoleList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	data, err := c.ProjectRolePermService.AllRoleList(tenantId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ret := iris.Map{"result": data}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: ret, Msg: _domain.NoErr.Msg})
}

// RolePermList
// @Tags	项目权限
// @summary	项目角色的权限列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string									true	"Authentication header"
// @Param 	currProjectId				query	int										true	"当前项目ID"
// @Param 	ProjectRolePermPaginateReq	query	serverDomain.ProjectRolePermPaginateReq	true	"获取项目角色的权限列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ProjectPerm}}
// @Router	/api/v1/projects/perms/rolePermList	[get]
func (c *ProjectRolePermCtrl) RolePermList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.ProjectRolePermPaginateReq
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.ProjectRolePermService.PaginateRolePerms(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// UserPermList
// @Tags	项目权限
// @summary	项目中用户的权限列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization				header	string									true	"Authentication header"
// @Param 	currProjectId				query	int										true	"当前项目ID"
// @Param 	ProjectUserPermsPaginate	query	serverDomain.ProjectUserPermsPaginate	true	"获取项目中用户的权限列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]model.ProjectPerm}}
// @Router	/api/v1/projects/perms/userPermList	[get]
func (c *ProjectRolePermCtrl) UserPermList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)

	var req serverDomain.ProjectUserPermsPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	req.ConvertParams()

	data, err := c.ProjectRolePermService.PaginateUserPerms(tenantId, req, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}
