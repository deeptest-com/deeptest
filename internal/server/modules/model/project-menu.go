package model

import "github.com/deeptest-com/deeptest/cmd/server/v1/domain"

type ProjectMenu struct {
	BaseModel
	serverDomain.ProjectMenuBase
}

func (ProjectMenu) TableName() string {
	return "biz_project_menu"
}
