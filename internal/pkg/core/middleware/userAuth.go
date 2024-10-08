package middleware

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/deeptest-com/deeptest/integration/domain"
	"github.com/deeptest-com/deeptest/integration/enum"
	"github.com/deeptest-com/deeptest/integration/service/user"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/deeptest-com/deeptest/saas/common"
	"github.com/deeptest-com/deeptest/saas/tenant"
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
		logUtils.Infof("authorization, appName:%s,token:%s,origin:%s,tenantId:%s", appName, token, origin, tenantId)
		if appName != "" {
			userInfo, err := user.GetUserInfoByToken(tenantId, token, origin)
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
	//修改saas默认示例用户为租户管理员
	updateManagerAccount(tenantId)

	req := v1.UserReq{UserBase: v1.UserBase{
		Username:  userInfo.Username,
		Email:     userInfo.Mail,
		ImAccount: userInfo.WxName,
		Name:      userInfo.RealName,
		Password:  commonUtils.RandStr(8),
	}}
	userRepo := repo.UserRepo{}
	userRepo.ProfileRepo = &repo.ProfileRepo{}
	userRepo.RoleRepo = &repo.RoleRepo{}
	userRepo.Create(tenantId, req)

	user, err := userRepo.GetByUsernameOrEmail(tenantId, userInfo.Username, userInfo.Mail)
	if err != nil {
		logUtils.Errorf("%s", err.Error())
		return
	}

	authorityType, _ := strconv.Atoi(string(tenantId))
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(user.ID), 10),
		Username:      user.Username,
		AuthorityId:   string(tenantId),
		AuthorityType: authorityType,
		LoginType:     multi.LoginTypeWx,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWx.Milliseconds(),
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		logUtils.Errorf("%s", err.Error())
		return
	}

	return
}

func getAppName(ctx *context.Context) (appName enum.AppName, token, origin string, tenantId consts.TenantId) {

	tenantId = common.GetTenantId(ctx)
	origin = ctx.GetHeader("Origin")
	//origin = "http://192.168.5.60:804"

	token = ctx.GetHeader("X-Token")
	if token != "" {
		appName = enum.Thirdparty
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

func updateManagerAccount(tenantId consts.TenantId) {
	if tenantId != "" {
		var ids []uint
		dao.GetDB(tenantId).Table("sys_user").Where("id = 2 and name = 'sys' ").Pluck("id", &ids)
		if len(ids) == 0 {
			return
		}

		info := tenant.NewTenant().GetInfo(tenantId, "")
		if info.ManagerId != 0 {
			if info.ManagerMail == "" {
				info.ManagerMail = fmt.Sprintf("%d", info.ManagerId)
			}
			sqls := []string{
				fmt.Sprintf("update sys_user set  username = '%d',name = '%s',email = '%s'  where id = 2 and  name = 'sys'", info.ManagerId, info.ManagerName, info.ManagerMail),
				fmt.Sprintf("update biz_endpoint set  create_user = '%d',update_user = '%d'  where project_id = 1 and create_user = 'sys'", info.ManagerId, info.ManagerId),
				fmt.Sprintf("update biz_scenario set  create_user_name = '%d'  where project_id = 1 and create_user_name = 'sys'", info.ManagerId),
			}
			for _, sql := range sqls {
				dao.GetDB(tenantId).Exec(sql)
			}

		}
	}
}
