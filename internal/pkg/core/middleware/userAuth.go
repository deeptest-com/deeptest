package middleware

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/integration/enum"
	"github.com/aaronchen2k/deeptest/integration/service/user"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/multi"
	"strconv"
	"strings"
	"time"
)

var whitelist []string

func init() {
	whitelist = []string{
		"/api/v1/account",
		"/swagger",
		"/upload",
		"/mocks",
		"/api/v1/message/receiveMcsApprovalData",
		"/api/v1/init/initdb",
		"/api/v1/openApi",
		"/api/v1/healthz",
		"/api/v1/saas",
	}
}

func isIgnore(path string) bool {
	if config.CONFIG.System.SysEnv == "" {
		return true
	}

	for _, item := range whitelist {
		if strings.HasPrefix(path, item) {
			return true
		}
	}
	return false
}

func UserAuth() iris.Handler {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader}
	verifier.ErrorHandler = func(ctx *context.Context, err error) {
		if isIgnore(ctx.Path()) {
			ctx.Next()
			return
		}

		appName, token, origin, tenantId := getAppName(ctx)
		user := user.NewUser(appName)

		if appName != "" {
			userInfo, err := user.GetUserInfoByToken(token, origin)
			if err == nil && userInfo.Username != "" {
				token, err := creatSession(tenantId, userInfo)
				if err == nil && token != "" {
					ctx.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
					ctx.Header("Authorization", token)
					ctx.Next()
					return
				}
			}
		}

		atoken := []byte(verifier.RequestToken(ctx))
		logUtils.Errorf("authorization failed, path:%s,token:%s,xToken:%", ctx.Path(), string(atoken), token)

		ctx.JSON(_domain.Response{
			Code: _domain.AuthErr.Code,
		})

	}
	return verifier.Verify()
}

func creatSession(tenantId consts.TenantId, userInfo integrationDomain.UserInfo) (token string, err error) {

	req := v1.UserReq{UserBase: v1.UserBase{
		Username:  userInfo.Username,
		Email:     userInfo.Mail,
		ImAccount: userInfo.WxName,
		Name:      userInfo.RealName,
		Password:  commonUtils.RandStr(8),
	}}
	userRepo := repo.UserRepo{DB: dao.GetDB()}
	userRepo.ProfileRepo = &repo.ProfileRepo{DB: dao.GetDB()}
	userRepo.RoleRepo = &repo.RoleRepo{DB: dao.GetDB()}
	userRepo.Create(tenantId, req)

	user, err := userRepo.GetByUsernameOrEmail(tenantId, userInfo.Username, userInfo.Mail)
	if err != nil {
		return
	}

	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(user.ID), 10),
		Username:      user.Username,
		AuthorityId:   "",
		AuthorityType: multi.AdminAuthority,
		LoginType:     multi.LoginTypeWx,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWx.Milliseconds(),
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return
	}
	return
}

func getAppName(ctx *context.Context) (appName enum.AppName, token, origin string, tenantId consts.TenantId) {

	origin = ctx.GetHeader("Origin")
	//origin = "http://192.168.5.60:804"

	token = ctx.GetHeader("X-Token")
	if token != "" {
		appName = enum.Leyan
		return
	}

	token = ctx.GetHeader("Token")
	//token = "61m3uc60xbeo"
	if token != "" {
		appName = enum.Lecang
		return
	}

	return
}
