package model

import "github.com/deeptest-com/deeptest/cmd/server/v1/domain"

type ProjectPerm struct {
	BaseModel
	serverDomain.ProjectPermBase
}

func (ProjectPerm) TableName() string {
	return "biz_project_perm"
}
