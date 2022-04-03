package model

import serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"

type SysRole struct {
	BaseModel
	serverDomain.RoleBase
}

func (SysRole) TableName() string {
	return "sys_role"
}
