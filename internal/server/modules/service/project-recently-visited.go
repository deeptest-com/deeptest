package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ProjectRecentlyVisitedService struct {
	ProjectRecentlyVisitedRepo *repo.ProjectRecentlyVisitedRepo `inject:""`
}

func (s *ProjectRecentlyVisitedService) Create(userId, projectId uint) (uint, error) {
	projectRecentlyVisitedBase := v1.ProjectRecentlyVisitedBase{UserId: userId, ProjectId: projectId}
	req := v1.ProjectRecentlyVisitedReq{ProjectRecentlyVisitedBase: projectRecentlyVisitedBase}
	return s.ProjectRecentlyVisitedRepo.Create(req)
}
