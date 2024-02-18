package handler

import (
	"github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"strings"
)

type AccountCtrl struct {
	BaseCtrl
	AccountService *service.AccountService
	UserRepo       *repo.UserRepo `inject:""`
}

// Login
// @Tags	认证模块
// @summary	登录
// @accept	application/json
// @Produce	application/json
// @Param 	LoginReq	body	serverDomain.LoginReq	true	"登录的请求体"
// @success	200	{object}	_domain.Response{data=serverDomain.LoginResp}
// @Router	/api/v1/account/login	[post]
func (c *AccountCtrl) Login(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)

		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("ValidRequest()", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}
	resp, err := c.AccountService.Login(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.LoginErr.Code, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: resp, Msg: _domain.NoErr.Msg})
}

// Register
// @Tags	认证模块
// @summary	注册
// @accept	application/json
// @Produce	application/json
// @Param 	RegisterReq	body	serverDomain.RegisterReq	true	"注册的请求体"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/account/register	[post]
func (c *AccountCtrl) Register(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.RegisterReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)

		if len(errs) > 0 {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.AccountService.Register(tenantId, req)

	ctx.JSON(_domain.Response{Code: err.Code, Msg: err.Msg})
}

// ForgotPassword
// @Tags	认证模块
// @summary	忘记密码
// @accept	application/json
// @Produce	application/json
// @Param 	usernameOrPassword	query	string	true	"用户名或者密码"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/account/forgotPassword	[post]
// @x-creator "wangzhen"
func (c *AccountCtrl) ForgotPassword(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	usernameOrPassword := ctx.URLParam("usernameOrPassword")
	if usernameOrPassword == "" {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "no usernameOrPassword param"})
		return
	}

	err := c.AccountService.ForgotPassword(tenantId, usernameOrPassword)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

// ResetPassword
// @Tags	认证模块
// @summary	重置密码
// @accept	application/json
// @Produce	application/json
// @Param 	ResetPasswordReq	body	serverDomain.ResetPasswordReq	true	"重置密码的参数"
// @success	200	{object}	_domain.Response
// @Router	/api/v1/account/resetPassword	[post]
func (c *AccountCtrl) ResetPassword(ctx iris.Context) {
	tenantId := c.getTenantId(ctx)
	var req serverDomain.ResetPasswordReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			ctx.JSON(_domain.Response{Code: _domain.ParamErr.Code, Msg: _domain.ParamErr.Msg})
			return
		}
	}

	err = c.AccountService.ResetPassword(tenantId, req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
