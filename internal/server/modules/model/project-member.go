package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ProjectMember struct {
	BaseModel

	ProjectId     uint `json:"projectId"`
	ProjectRoleId uint `json:"projectRoleId"`
	UserId        uint `json:"userId"`

	ProjectRoleName consts.RoleType `gorm:"-" json:"-"`
}

func (ProjectMember) TableName() string {
	return "biz_project_member"
}

type ProjectMemberRole struct {
	Project
	RoleId   uint            `gorm:"role_id" json:"roleId"`
	RoleName consts.RoleType `gorm:"role_name" json:"roleName"`
}
