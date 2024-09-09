package service

import (
	thirdparty "github.com/aaronchen2k/deeptest/integration/thirdparty/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type PrivilegeService struct {
	RemoteService *thirdparty.RemoteService `inject:""`
}

func (s *PrivilegeService) GetAll(tenantId consts.TenantId, username, roleCode string, needSysAuth bool) (ret []string, err error) {
	if roleCode != "" && !needSysAuth {
		ret, err = s.RemoteService.GetRoleMenus(tenantId, roleCode)
		if err != nil {
			return
		}
	}

	var points []string
	points, err = s.RemoteService.GetUserButtonPermissions(tenantId, username)
	if err != nil {
		return
	}

	ret = append(ret, points...)
	ret = commonUtils.ArrayUnique(ret)

	return
}
