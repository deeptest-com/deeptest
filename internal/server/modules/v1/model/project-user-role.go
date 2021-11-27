package model

import ()

type ProjectUserRole struct {
	BaseModel

	ProjectId uint `json:"projectId"`
	UserId    uint `json:"userId"`
	RoleId    uint `json:"roleId"`
}

func (ProjectUserRole) TableName() string {
	return "r_project_user_role"
}
