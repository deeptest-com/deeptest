package service

import (
	"github.com/aaronchen2k/deeptest/integration/enum"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/snowlyg/helper/arr"
	"time"
)

type RoleService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *RoleService) GetRoleValueNameMap(tenantId consts.TenantId) (res map[string]string, err error) {
	spaceRoles, err := s.RemoteService.GetSpaceRoles(tenantId)
	if err != nil {
		return
	}

	res = make(map[string]string)
	for _, v := range spaceRoles {
		res[v.RoleValue] = v.RoleName
	}

	return
}

func (s *RoleService) GetRoleNameByValue(tenantId consts.TenantId, value string) (res string, err error) {
	roleValueNameMap, err := s.GetRoleValueNameMap(tenantId)
	if err != nil {
		return
	}

	if name, ok := roleValueNameMap[value]; ok {
		res = name
	}

	return
}

func (s *RoleService) GetUserRoleArr(tenantId consts.TenantId, username string) (ret []string, err error) {
	roles, err := s.RemoteService.GetUserOpenRoles(tenantId, username)
	if err != nil {
		return
	}

	for _, v := range roles {
		ret = append(ret, v.RoleValue)
	}

	return
}

func (s *RoleService) IsSuperAdmin(tenantId consts.TenantId, username string) (ret bool, err error) {
	roleValueArr, err := s.GetUserRoleArr(tenantId, username)
	if err != nil {
		return
	}

	ret = arr.InArrayS(roleValueArr, enum.SuperAdmin)

	return
}

func (s *RoleService) IsSuperAdminInCache(tenantId consts.TenantId, username string) (ret bool, err error) {
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

func (s *RoleService) SetIsSuperAdminCache(tenantId consts.TenantId, username string) (ret bool, err error) {
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
		err = cache.SetCache(redisKey, serverConsts.IsAdminRole, time.Hour*4)
	} else {
		err = cache.SetCache(redisKey, serverConsts.IsNotAdminRole, time.Hour*4)
	}

	return
}
