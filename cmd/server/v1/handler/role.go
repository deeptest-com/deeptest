package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/snowlyg/multi"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type RoleCtrl struct {
	BaseCtrl
	RoleService *service.RoleService `inject:""`
}

// GetAllRoles 分页列表
// @Tags	角色模块
// @summary	角色列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	RoleReqPaginate	query	serverDomain.RoleReqPaginate	true	"获取角色列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]serverDomain.RoleResp}}
// @Router	/api/v1/roles	[get]
func (c *RoleCtrl) GetAllRoles(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.RoleReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	data, err := c.RoleService.Paginate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// GetRole 详情
// @Tags	角色模块
// @summary	角色列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"角色ID"
// @success	200	{object}	_domain.Response{data=serverDomain.RoleResp}
// @Router	/api/v1/roles/{id}	[get]
func (c *RoleCtrl) GetRole(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	role, err := c.RoleService.FindById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: role, Msg: _domain.NoErr.Msg})
}

// CreateRole 添加
// @Tags	角色模块
// @summary	新建角色
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	RoleReq			body	serverDomain.RoleReq		true	"新建角色的请求参数"
// @success	200	{object}	_domain.Response{data=object{id=int}}
// @Router	/api/v1/roles	[post]
func (c *RoleCtrl) CreateRole(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.RoleReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	id, err := c.RoleService.Create(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

// UpdateRole 更新
// @Tags	角色模块
// @summary	编辑角色
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	id				path	int						true	"角色ID"
// @Param 	RoleReq			body	serverDomain.RoleReq	true	"编辑角色的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/roles/{id}	[post]
func (c *RoleCtrl) UpdateRole(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id, _ := ctx.Params().GetInt("id")

	var req serverDomain.RoleReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.RoleService.Update(tenantId, uint(id), req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// DeleteRole 删除
// @Tags	角色模块
// @summary	删除角色
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"角色ID"
// @success	200	{object}	_domain.Response{data=serverDomain.RoleResp}
// @Router	/api/v1/roles/{id}	[delete]
func (c *RoleCtrl) DeleteRole(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	err := c.RoleService.DeleteById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// AllRoleList
// @Tags	角色模块
// @summary	无分页的角色列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=object{result=[]serverDomain.RoleResp}}
// @Router	/api/v1/roles/all	[get]
func (c *RoleCtrl) AllRoleList(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	roles, err := c.RoleService.AllRoleList(tenantId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	data := iris.Map{"result": roles}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

func (c *RoleCtrl) GetAuthByEnv(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)
	res, err := c.RoleService.GetAuthByEnv(tenantId, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: res, Msg: _domain.NoErr.Msg})
}
