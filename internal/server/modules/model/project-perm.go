package model

import "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type ProjectPerm struct {
	BaseModel
	serverDomain.ProjectPermBase
}

func (ProjectPerm) TableName() string {
	return "biz_project_perm"
}
