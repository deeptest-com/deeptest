package model

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type ProjectRecentlyVisited struct {
	BaseModel
	v1.ProjectRecentlyVisitedBase
}

func (ProjectRecentlyVisited) TableName() string {
	return "biz_project_recently_visited"
}
