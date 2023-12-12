package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ProjectMenuService struct {
	ProjectRepo     *repo.ProjectRepo     `inject:""`
	ProjectMenuRepo *repo.ProjectMenuRepo `inject:""`
	ProjectRoleRepo *repo.ProjectRoleRepo `inject:""`
	UserRepo        *repo.UserRepo        `inject:""`
	RemoteService   *RemoteService        `inject:""`
}

func (s *ProjectMenuService) GetUserMenuList(projectId, userId uint) (ret []model.ProjectMenu, err error) {
	var roleId uint
	isAdminUser, err := s.UserRepo.IsAdminUser(userId)
	if err != nil {
		return
	}

	if isAdminUser {
		adminProjectRole, err := s.ProjectRoleRepo.GetAdminRecord()
		if err != nil {
			return ret, err
		}
		roleId = adminProjectRole.ID
	} else {
		projectMemberRole, err := s.ProjectRepo.FindRolesByProjectAndUser(projectId, userId)
		if err != nil {
			return ret, err
		}
		roleId = projectMemberRole.ProjectRoleId
	}

	ret, err = s.ProjectMenuRepo.GetRoleMenuList(roleId)
	return
}

func (s *ProjectMenuService) GetUserMenuListNew(projectId, userId uint, userName string) (ret []string, err error) {
	if config.CONFIG.System.SysEnv == "ly" {
		project, err := s.ProjectRepo.Get(projectId)
		if err != nil {
			return ret, err
		}
		ret, err = s.RemoteService.GetUserButtonPermissions(userName, project.ShortName)
	} else {
		projectMemberRole, err := s.ProjectRepo.FindRolesByProjectAndUser(projectId, userId)
		if err != nil {
			return ret, err
		}
		roleId := projectMemberRole.ProjectRoleId
		ret, err = s.ProjectMenuRepo.GetRoleMenuCodeList(roleId)
	}

	return
}
