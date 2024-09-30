package model

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
)

type SysRole struct {
	BaseModel
	serverDomain.RoleBase
}

func (SysRole) TableName() string {
	return "sys_role"
}
