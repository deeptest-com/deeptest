package model

import "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type SysPerm struct {
	BaseModel
	serverDomain.PermBase
}

func (SysPerm) TableName() string {
	return "sys_perm"
}
