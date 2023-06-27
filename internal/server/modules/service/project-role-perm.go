package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectRolePermService struct {
	ProjectRepo         *repo.ProjectRepo         `inject:""`
	ProjectRolePermRepo *repo.ProjectRolePermRepo `inject:""`
	ProjectRoleRepo     *repo.ProjectRoleRepo     `inject:""`
	ProjectRoleMenuRepo *repo.ProjectRoleMenuRepo `inject:""`
	ProfileRepo         *repo.ProfileRepo         `inject:""`
}

func (s *ProjectRolePermService) AllRoleList() (data []model.ProjectRole, err error) {
	return s.ProjectRoleRepo.AllRoleList()
}

func (s *ProjectRolePermService) GetProjectUserRole(userId, projectId uint) (data model.ProjectRole, err error) {
	return s.ProjectRoleRepo.ProjectUserRoleList(userId, projectId)
}

func (s *ProjectRolePermService) PaginateRolePerms(req v1.ProjectRolePermPaginateReq) (ret _domain.PageData, err error) {
	return s.ProjectRolePermRepo.PaginateRolePerms(req)
}

func (s *ProjectRolePermService) PaginateUserPerms(req v1.ProjectUserPermsPaginate, userId uint) (ret _domain.PageData, err error) {
	return s.ProjectRolePermRepo.UserPermList(req, userId)
}
