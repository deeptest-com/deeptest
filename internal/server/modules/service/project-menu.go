package service

import (
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

func (s *ProjectMenuService) GetUserButtonList(projectId, userId uint, xToken string) (ret []string, err error) {
	//var roleId uint
	isAdminUser, err := s.UserRepo.IsAdminUser(userId)
	if err != nil {
		return
	}

	if isAdminUser {
		// TODO 需要用角色名去乐研获取权限列表
		//adminProjectRole, err := s.ProjectRoleRepo.GetAdminRecord()
		//if err != nil {
		//	return ret, err
		//}
		//roleId = adminProjectRole.ID
	} else {
		//projectMemberRole, err := s.ProjectRepo.FindRolesByProjectAndUser(projectId, userId)
		//if err != nil {
		//	return ret, err
		//}
		//roleId = projectMemberRole.ProjectRoleId

		project, err := s.ProjectRepo.Get(projectId)
		if err != nil {
			return
		}
		ret, err = s.RemoteService.GetUserButtonPermissions(xToken, project.ShortName)
	}

	//ret, err = s.ProjectMenuRepo.GetRoleMenuList(roleId)
	return
}
