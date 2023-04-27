package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ProjectMenuService struct {
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ProjectMenuRepo *repo.ProjectMenuRepo `inject:""`
	ProjectRoleRepo *repo.ProjectRoleRepo `inject:""`
}

func NewProjectMenuService() *ProjectMenuService {
	return &ProjectMenuService{}
}

func (s *ProjectMenuService) GetUserMenuList(userId uint) (ret []model.ProjectMenu, err error) {
	projectMemberRole, err := s.ProjectRepo.GetCurrProjectMemberRoleByUser(userId)
	if err != nil {
		return
	}

	projectRole, err := s.ProjectRoleRepo.FindById(projectMemberRole.ProjectRoleId)
	if err != nil {
		return
	}

	return s.ProjectMenuRepo.GetRoleMenuList(projectMemberRole.ProjectRoleId, projectRole.Name)
}
