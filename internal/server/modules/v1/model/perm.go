package model

import serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"

type SysPerm struct {
	BaseModel
	serverDomain.PermBase
}

func (SysPerm) TableName() string {
	return "sys_perm"
}
