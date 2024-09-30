package model

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
)

type SysUser struct {
	BaseModel

	serverDomain.UserBase

	Password string         `gorm:"type:varchar(250)" json:"password" validate:"required"`
	Vcode    string         `json:"vcode"`
	Profile  SysUserProfile `gorm:"foreignKey:user_id"`

	RoleIds []uint `gorm:"-" json:"role_ids"`
}

type Avatar struct {
	Avatar string `json:"avatar"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
