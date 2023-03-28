package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type ProjectPerm struct {
	BaseModel
	v1.ProjectPermBase
}

func (ProjectPerm) TableName() string {
	return "biz_project_perm"
}
