package model

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type Project struct {
	BaseModel
	v1.ProjectBase

	Spec          string `json:"spec"`
	EnvironmentId uint   `json:"environmentId"`
}

func (Project) TableName() string {
	return "biz_project"
}
