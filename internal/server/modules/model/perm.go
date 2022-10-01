package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type SysPerm struct {
	BaseModel
	v1.PermBase
}

func (SysPerm) TableName() string {
	return "sys_perm"
}
