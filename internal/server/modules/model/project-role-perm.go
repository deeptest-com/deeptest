package model

import "github.com/deeptest-com/deeptest/cmd/server/v1/domain"

type ProjectRolePerm struct {
	BaseModel
	serverDomain.ProjectRolePermBase
}

func (ProjectRolePerm) TableName() string {
	return "biz_project_role_perm"
}
