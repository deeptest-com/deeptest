package middleware

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/multi"
	"strconv"
	"strings"
	"time"
)

/*
func VerifyAuth() iris.Handler {
	return func(ctx *context.Context) {
		authorization := ctx.GetHeader("Authorization")
		xToken := ctx.GetHeader("X-Token")
		if xToken != "" && authorization == "" {
			userInfo, err := new(service.RemoteService).GetUserInfoByToken(xToken)
			if err == nil && userInfo.Username != "" {
				token, err := creatSession(userInfo)
				if err == nil && token != "" {
					ctx.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
					ctx.Header("Authorization", token)
				}
			}
		}

		ctx.Next()
	}
}
*/

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
	}
}

func isIgnore(path string) bool {
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

		xToken := ctx.GetHeader("X-Token")
		if xToken != "" {
			userInfo, err := new(service.RemoteService).GetUserInfoByToken(xToken)
			if err == nil && userInfo.Username != "" {
				token, err := creatSession(userInfo)
				if err == nil && token != "" {
					ctx.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
					ctx.Header("Authorization", token)
					ctx.Next()
					return
				}
			}
		}

		ctx.JSON(_domain.Response{
			Code: _domain.AuthErr.Code,
		})

	}
	return verifier.Verify()
}

func creatSession(userInfo v1.UserInfo) (token string, err error) {

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
	userRepo.Create(req)

	user, err := userRepo.GetByUsernameOrEmail(userInfo.Username, userInfo.Mail)
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
