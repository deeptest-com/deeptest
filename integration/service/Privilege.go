package service

import commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"

type PrivilegeService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *PrivilegeService) GetAll(username, roleCode string) (ret []string, err error) {
	ret, err = s.RemoteService.GetRoleMenus(roleCode)
	if err != nil {
		return
	}
	var points []string
	points, err = s.RemoteService.GetUserButtonPermissions(username)
	if err != nil {
		return
	}

	ret = append(ret, points...)
	ret = commonUtils.ArrayUnique(ret)

	return
}
