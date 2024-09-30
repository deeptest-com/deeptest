package model

import (
	"github.com/deeptest-com/deeptest/cmd/server/v1/domain"
)

type Project struct {
	BaseModel
	serverDomain.ProjectBase

	Spec          string `json:"spec"`
	Spec2         string `json:"spec2"`
	EnvironmentId uint   `json:"environmentId"`
}

func (Project) TableName() string {
	return "biz_project"
}
