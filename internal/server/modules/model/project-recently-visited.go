package model

import "github.com/deeptest-com/deeptest/cmd/server/v1/domain"

type ProjectRecentlyVisited struct {
	BaseModel
	serverDomain.ProjectRecentlyVisitedBase
}

func (ProjectRecentlyVisited) TableName() string {
	return "biz_project_recently_visited"
}
