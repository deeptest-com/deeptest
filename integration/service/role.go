package service

type RoleService struct {
	RemoteService *RemoteService `inject:""`
}

func (s *RoleService) GetRoleValueNameMap() (res map[string]string, err error) {
	spaceRoles, err := s.RemoteService.GetSpaceRoles()
	if err != nil {
		return
	}

	res = make(map[string]string)
	for _, v := range spaceRoles {
		res[v.RoleValue] = v.RoleName
	}

	return
}

func (s *RoleService) GetRoleNameByValue(value string) (res string, err error) {
	roleValueNameMap, err := s.GetRoleValueNameMap()
	if err != nil {
		return
	}

	if name, ok := roleValueNameMap[value]; ok {
		res = name
	}

	return
}