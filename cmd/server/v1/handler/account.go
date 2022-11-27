package handler

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/aaronchen2k/deeptest/pkg/lib/log"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type AccountCtrl struct {
	AccountService *service.AccountService `inject:""`
	UserRepo       *repo.UserRepo          `inject:""`
}

func (c *AccountCtrl) Login(ctx iris.Context) {
	var req v1.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)

		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("ValidRequest()", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}
	token, err := c.AccountService.Login(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"token": token}, Msg: _domain.NoErr.Msg})
}

func (c *AccountCtrl) Register(ctx iris.Context) {
	var req v1.RegisterReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)

		if len(errs) > 0 {
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}

	err := c.AccountService.Register(req)

	ctx.JSON(_domain.Response{Code: err.Code, Msg: err.Msg})
}

func (c *AccountCtrl) ForgotPassword(ctx iris.Context) {
	usernameOrPassword := ctx.URLParam("usernameOrPassword")
	if usernameOrPassword == "" {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: "no usernameOrPassword param"})
		return
	}

	err := c.AccountService.ForgotPassword(usernameOrPassword)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}

func (c *AccountCtrl) ResetPassword(ctx iris.Context) {
	var req v1.ResetPasswordReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Msg: err.Error()})
		return
	}

	err = c.AccountService.ResetPassword(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Msg: _domain.NoErr.Msg})
}
