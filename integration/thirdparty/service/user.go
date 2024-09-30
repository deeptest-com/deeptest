package service

import (
	v1 "github.com/deeptest-com/deeptest/integration/domain"
	"github.com/deeptest-com/deeptest/integration/enum"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/cache"
	"github.com/snowlyg/helper/arr"
	"time"
)

type User struct {
}

func (s *User) GetUserInfoByToken(tenantId consts.TenantId, token, origin string) (user v1.UserInfo, err error) {
	user, err = new(RemoteService).GetUserInfoByToken(tenantId, token)
	return
}

func (s *User) IsSuperAdmin(tenantId consts.TenantId, username string) (ret bool, err error) {
	roleValueArr, err := s.GetUserRoleArr(tenantId, username)
	if err != nil {
		return
	}

	ret = arr.InArrayS(roleValueArr, enum.SuperAdmin)

	return
}

func (s *User) GetUserRoleArr(tenantId consts.TenantId, username string) (ret []string, err error) {
	roles, err := new(RemoteService).GetUserOpenRoles(tenantId, username)
	if err != nil {
		return
	}

	for _, v := range roles {
		ret = append(ret, v.RoleValue)
	}

	return
}

func (s *User) IsSuperAdminInCache(tenantId consts.TenantId, username string) (ret bool, err error) {
	redisKey := string(tenantId) + "-" + "isAdmin-" + username

	isAdmin, err := cache.GetCacheString(redisKey)
	if err == nil {
		if isAdmin == serverConsts.IsAdminRole {
			ret = true
		}
		return ret, err
	}

	return
}

func (s *User) SetIsSuperAdminCache(tenantId consts.TenantId, username string) (ret bool, err error) {
	ret, err = s.IsSuperAdminInCache(tenantId, username)
	if err == nil {
		return
	}

	ret, err = s.IsSuperAdmin(tenantId, username)
	if err != nil {
		return
	}

	redisKey := string(tenantId) + "-" + "isAdmin-" + username
	if ret {
		err = cache.SetCache(redisKey, serverConsts.IsAdminRole, time.Second*180)
	} else {
		err = cache.SetCache(redisKey, serverConsts.IsNotAdminRole, time.Second*180)
	}

	return
}
