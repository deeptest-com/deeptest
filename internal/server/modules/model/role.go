package model

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type SysRole struct {
	BaseModel
	v1.RoleBase
}

func (SysRole) TableName() string {
	return "sys_role"
}
