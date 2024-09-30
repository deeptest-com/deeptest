package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type ProjectRecentlyVisitedService struct {
	ProjectRecentlyVisitedRepo *repo.ProjectRecentlyVisitedRepo `inject:""`
}

func (s *ProjectRecentlyVisitedService) Create(tenantId consts.TenantId, userId, projectId uint) (uint, error) {
	projectRecentlyVisitedBase := v1.ProjectRecentlyVisitedBase{UserId: userId, ProjectId: projectId}
	req := v1.ProjectRecentlyVisitedReq{ProjectRecentlyVisitedBase: projectRecentlyVisitedBase}
	return s.ProjectRecentlyVisitedRepo.Create(tenantId, req)
}
