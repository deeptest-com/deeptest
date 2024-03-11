package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

type UserCtrl struct {
	BaseCtrl
	UserService *service.UserService `inject:""`
	UserRepo    *repo.UserRepo       `inject:""`
}

// ListAll
// @Tags	用户管理
// @summary	用户列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string				true	"Authentication header"
// @Param 	currProjectId	query	int					true	"当前项目ID"
// @Param 	UserReqPaginate		query	serverDomain.UserReqPaginate	true	"查询参数"
// @success	200	{object}	_domain.Response{data=_domain.PageData{result=[]serverDomain.UserResp}}
// @Router	/api/v1/users	[get]
func (c *UserCtrl) ListAll(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.UserReqPaginate

	if err := ctx.ReadQuery(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	data, err := c.UserRepo.Paginate(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// GetUser 详情
// @Tags	用户管理
// @summary	用户详情
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"用户ID"
// @success	200	{object}	_domain.Response{data=serverDomain.UserResp}
// @Router	/api/v1/users/{id}	[get]
func (c *UserCtrl) GetUser(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		_logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	user, err := c.UserRepo.FindDetailById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: _domain.SystemErr.Msg})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// Invite 邀请用户
// @Tags	用户管理
// @summary	邀请用户
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	InviteUserReq	body	serverDomain.InviteUserReq	true	"邀请用户的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/invite	[post]
func (c *UserCtrl) Invite(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.InviteUserReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	projectId, err := ctx.URLParamInt("currProjectId")
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}
	req.ProjectId = projectId
	_, bizErr := c.UserService.Invite(tenantId, req)
	if bizErr != nil {
		ctx.JSON(_domain.Response{Code: bizErr.Code})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// UpdateEmail 修改邮箱
// @Tags	用户管理
// @summary	修改邮箱
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	UpdateUserReq	body	serverDomain.UpdateUserReq	true	"修改邮箱的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/updateEmail	[post]
func (c *UserCtrl) UpdateEmail(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)
	req := serverDomain.UpdateUserReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	po, _ := c.UserRepo.FindByEmail(tenantId, req.Email, userId)
	if po.Id > 0 {
		bizErr := _domain.ErrEmailExist
		ctx.JSON(_domain.Response{Code: bizErr.Code})
		return
	}

	err = c.UserRepo.UpdateEmail(tenantId, req.Email, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	user, err := c.UserRepo.FindDetailById(tenantId, userId)
	user.Password = ""
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// UpdateName 修改名称
// @Tags	用户管理
// @summary	修改名称
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	UpdateUserReq	body	serverDomain.UpdateUserReq	true	"修改名称的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.UserResp}
// @Router	/api/v1/users/updateName	[post]
func (c *UserCtrl) UpdateName(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)
	req := serverDomain.UpdateUserReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	po, _ := c.UserRepo.FindByUserName(tenantId, req.Username, userId)
	if po.Id > 0 {
		bizErr := _domain.ErrUsernameExist
		ctx.JSON(_domain.Response{Code: bizErr.Code})
		return
	}

	err = c.UserRepo.UpdateName(tenantId, req.Username, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	user, err := c.UserRepo.FindDetailById(tenantId, userId)
	user.Password = ""
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// UpdatePassword 修改密码
// @Tags	用户管理
// @summary	修改密码
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string						true	"Authentication header"
// @Param 	currProjectId	query	int							true	"当前项目ID"
// @Param 	UpdateUserReq	body	serverDomain.UpdateUserReq	true	"修改密码的请求参数"
// @success	200	{object}	_domain.Response{data=serverDomain.UserResp}
// @Router	/api/v1/users/updatePassword	[post]
func (c *UserCtrl) UpdatePassword(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)

	req := serverDomain.UpdateUserReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	err = c.UserRepo.ChangePassword(tenantId, req, userId)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	user, err := c.UserRepo.FindDetailById(tenantId, userId)
	user.Password = ""
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// Profile 个人信息
// @Tags	用户管理
// @summary	个人信息
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=serverDomain.UserResp}
// @Router	/api/v1/users/profile	[get]
func (c *UserCtrl) Profile(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	id := multi.GetUserId(ctx)
	if id == 0 {
		ctx.JSON(_domain.Response{Code: _domain.ErrNoUser.Code, Msg: _domain.SystemErr.Msg})
		return
	}

	user, err := c.UserRepo.FindDetailById(tenantId, id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	user.Password = ""

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: user, Msg: _domain.NoErr.Msg})
}

// Message 消息
func (c *UserCtrl) Message(ctx iris.Context) {
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// CreateUser 添加
// @Tags	用户管理
// @summary	新建用户
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	UserReq			body	serverDomain.UserReq	true	"新建用户的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users	[post]
func (c *UserCtrl) CreateUser(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	req := serverDomain.UserReq{}
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	//区分手动手动添加的账号和域账号登录，true 为手动创建，非true 为域账号
	req.Type = true
	id, err := c.UserRepo.Create(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"id": id}, Msg: _domain.NoErr.Msg})
}

// UpdateUser 更新
// @Tags	用户管理
// @summary	编辑用户
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string					true	"Authentication header"
// @Param 	currProjectId	query	int						true	"当前项目ID"
// @Param 	id				path	int						true	"用户ID"
// @Param 	UserReq			body	serverDomain.UserReq	true	"编辑用户的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/{id}	[post]
func (c *UserCtrl) UpdateUser(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	userId := multi.GetUserId(ctx)
	var reqId _domain.ReqId
	if err := ctx.ReadParams(&reqId); err != nil {
		_logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}

	var req serverDomain.UserReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.UserRepo.Update(tenantId, userId, reqId.Id, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// DeleteUser 删除
// @Tags	用户管理
// @summary	删除用户
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @Param 	id				path	int		true	"用户ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/{id}	[delete]
func (c *UserCtrl) DeleteUser(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req _domain.ReqId
	if err := ctx.ReadParams(&req); err != nil {
		_logUtils.Errorf("参数解析失败", zap.String("错误:", err.Error()))
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
		return
	}
	err := c.UserRepo.DeleteById(tenantId, req.Id)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Logout 退出
// @Tags	用户管理
// @summary	退出登录
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/logout	[get]
func (c *UserCtrl) Logout(ctx iris.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "授权凭证为空"})
		return
	}
	err := c.UserRepo.DelToken(string(token))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// Clear 清空 token
// @Tags	用户管理
// @summary	清空 token
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/clear	[get]
func (c *UserCtrl) Clear(ctx iris.Context) {
	token := multi.GetVerifiedToken(ctx)
	if token == nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "授权凭证为空"})
		return
	}
	if err := c.UserRepo.CleanToken(multi.AdminAuthority, string(token)); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ChangeAvatar 修改头像
// @Tags	用户管理
// @summary	修改头像
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string			true	"Authentication header"
// @Param 	currProjectId	query	int				true	"当前项目ID"
// @Param 	Avatar			body	model.Avatar	true	"头像"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/change_avatar	[post]
func (c *UserCtrl) ChangeAvatar(ctx iris.Context) {
	avatar := &model.Avatar{}
	if err := ctx.ReadJSON(avatar); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	err := c.UserRepo.UpdateAvatar(multi.GetUserId(ctx), avatar.Avatar)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// GetUsersNotExistedInProject
// @Tags	用户管理
// @summary	获取项目中没有的用户列表
// @accept 	application/json
// @Produce application/json
// @Param	Authorization	header	string	true	"Authentication header"
// @Param 	currProjectId	query	int		true	"当前项目ID"
// @success	200	{object}	_domain.Response{data=object{result=[]serverDomain.UserResp}}
// @Router	/api/v1/users/usersNotExistedInProject	[get]
func (c *UserCtrl) GetUsersNotExistedInProject(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	projectId, _ := ctx.URLParamInt("currProjectId")

	users, err := c.UserService.GetUsersNotExistedInProject(tenantId, uint(projectId))
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	data := iris.Map{"result": users}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: data, Msg: _domain.NoErr.Msg})
}

// ChangeUserSysRole
// @Tags	用户管理
// @summary	更新用户系统角色
// @accept 	application/json
// @Produce application/json
// @Param	Authorization		header	string							true	"Authentication header"
// @Param 	currProjectId		query	int								true	"当前项目ID"
// @Param 	UpdateUserRoleReq	body	serverDomain.UpdateUserRoleReq	true	"更新用户系统角色的请求参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/users/changeUserSysRole	[post]
func (c *UserCtrl) ChangeUserSysRole(ctx iris.Context) {
	//SAAS

	tenantId := c.getTenantId(ctx)
	req := serverDomain.UpdateUserRoleReq{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: err.Error()})
		return
	}

	if err = c.UserService.UpdateSysRoleForUser(tenantId, req.UserId, req.RoleIds); err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
