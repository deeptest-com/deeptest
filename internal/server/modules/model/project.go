package model

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
)

type Project struct {
	BaseModel
	v1.ProjectBase

	Spec           string `json:"spec"`
	Spec2          string `json:"spec2"`
	EnvironmentId  uint   `json:"environmentId"`
	Logo           string `json:"logo"`
	ShortName      string `json:"shortName"`
	IncludeExample bool   `json:"includeExample"`
	AdminId        uint   `json:"AdminId"`
}

func (Project) TableName() string {
	return "biz_project"
}
