package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type ProjectMenu struct {
	BaseModel
	v1.ProjectMenuBase
}

func (ProjectMenu) TableName() string {
	return "biz_project_menu"
}
