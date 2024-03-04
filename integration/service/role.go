package service

import (
	"github.com/aaronchen2k/deeptest/integration/enum"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/snowlyg/helper/arr"
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
