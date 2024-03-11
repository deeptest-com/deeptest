package middleware

import (
	"errors"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	casbinServer "github.com/aaronchen2k/deeptest/internal/server/core/casbin"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/saas/common"
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// Casbin Casbin 权鉴中间件
func Casbin() iris.Handler {
	return func(ctx *context.Context) {
		userId := multi.GetUserId(ctx)
		tenantId := common.GetTenantId(ctx)
		check, err := Check(tenantId, ctx.Request(), strconv.FormatUint(uint64(userId), 10))
		if err != nil || !check {
			ctx.JSON(_domain.Response{Code: _domain.AuthActionErr.Code, Data: nil, Msg: err.Error()})
			ctx.StopExecution()
			return
		}

		ctx.Next()
	}
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.
func Check(tenantId consts.TenantId, r *http.Request, userId string) (bool, error) {
	method := r.Method
	path := r.URL.Path

	ok, err := casbinServer.Instance(tenantId).Enforce(userId, path, method)
	if err != nil {
		logUtils.Errorf(fmt.Sprintf("验证权限报错：%s-%s-%s", userId, path, method), zap.String("casbinServer.Instance().Enforce()", err.Error()))
		return false, err
	}

	logUtils.Debugf(fmt.Sprintf("权限：%s-%s-%s", userId, path, method))
	fmt.Println(fmt.Sprintf("权限：%s-%s-%s", userId, path, method))

	if !ok {
		return ok, errors.New("你未拥有当前系统操作权限，请联系管理员")
	}

	return ok, nil
}
