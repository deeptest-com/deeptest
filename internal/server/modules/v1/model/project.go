package model

import serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"

type Project struct {
	BaseModel
	serverDomain.ProjectBase

	DefaultEnvironmentId uint `json:"defaultEnvironmentId"`
}

func (Project) TableName() string {
	return "biz_project"
}
