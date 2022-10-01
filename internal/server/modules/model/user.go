package model

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type SysUser struct {
	BaseModel

	v1.UserBase

	Password string `gorm:"type:varchar(250)" json:"password" validate:"required"`

	Profile SysUserProfile `gorm:"foreignKey:user_id"`

	RoleIds []uint `gorm:"-" json:"role_ids"`
}

type Avatar struct {
	Avatar string `json:"avatar"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
