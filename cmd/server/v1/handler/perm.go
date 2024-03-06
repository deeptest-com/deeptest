package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type PermCtrl struct {
	BaseCtrl
	PermService *service.PermService `inject:""`
}

// GetAllPerms 分页列表
// @Tags	权限模块
// @summary	权限列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string							true	"Authentication header"
// @Param 	currProjectId	query	int								true	"当前项目ID"
// @Param 	PermReqPaginate	query	serverDomain.PermReqPaginate	true	"权限列表的请求参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]serverDomain.PermResp}}
// @Router	/api/v1/perms	[get]
func (c *PermCtrl) GetAllPerms(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.PermReqPaginate
	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	data, err := c.PermService.Paginate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// GetPerm 详情
// @Tags	权限模块
// @summary	权限详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"权限ID"
// @success	200	{object}	_domain.Response{data=serverDomain.PermResp}
// @Router	/api/v1/perms/{id}	[get]
func (c *PermCtrl) GetPerm(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	perm, err := c.PermService.FindById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: perm, Msg: _domain.NoErr.Msg})
}

// CreatePerm 添加
// @Tags	权限模块
// @summary	新建权限
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	PermReq			body	serverDomain.PermReq	true	"新建权限的请求参数"
// @success	200	{object}	_domain.Response{data=object{id=int}}
// @Router	/api/v1/perms	[post]
func (c *PermCtrl) CreatePerm(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.PermReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	id, err := c.PermService.Create(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

// UpdatePerm 更新
// @Tags	权限模块
// @summary	编辑权限
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	id				path	int						true	"权限ID"
// @Param 	PermReq			body	serverDomain.PermReq	true	"编辑权限的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/perms/{id}	[post]
func (c *PermCtrl) UpdatePerm(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var reqId _domain.ReqId
	if err := ctx.ReadParams(&reqId); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.PermReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.PermService.Update(tenantId, reqId.Id, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// DeletePerm 删除
// @Tags	权限模块
// @summary	删除权限
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"权限ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/perms/{id}	[delete]
func (c *PermCtrl) DeletePerm(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	err := c.PermService.DeleteById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
