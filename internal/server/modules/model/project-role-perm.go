package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type ProjectRolePerm struct {
	BaseModel
	v1.ProjectRolePermBase
}

func (ProjectRolePerm) TableName() string {
	return "biz_project_role_perm"
}
