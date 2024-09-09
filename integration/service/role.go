package service

import (
	thirdparty "github.com/aaronchen2k/deeptest/integration/thirdparty/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type RoleService struct {
	RemoteService *thirdparty.RemoteService `inject:""`
	UserService   *thirdparty.User          `inject:""`
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
