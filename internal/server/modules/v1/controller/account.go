package controller

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/internal/server/core/web/validate"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/service"
	"strings"

	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

type AccountCtrl struct {
	AuthService *service.AccountService `inject:""`
	UserRepo    *repo.UserRepo          `inject:""`
}

func NewAccountCtrl() *AccountCtrl {
	return &AccountCtrl{}
}

func (c *AccountCtrl) Login(ctx iris.Context) {
	var req serverDomain.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)

		if len(errs) > 0 {
			_logUtils.Errorf("参数验证失败", zap.String("ValidRequest()", strings.Join(errs, ";")))
			ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}
	token, err := c.AuthService.GetAccessToken(req)
	if err != nil {
		ctx.JSON(_domain.Response{Code: _domain.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(_domain.Response{Code: _domain.NoErr.Code, Data: iris.Map{"token": token}, Msg: _domain.NoErr.Msg})
}
