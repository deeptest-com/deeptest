package service

import (
	"github.com/aaronchen2k/deeptest/integration/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type ProjectMenuService struct {
	ProjectRepo      *repo.ProjectRepo         `inject:""`
	ProjectMenuRepo  *repo.ProjectMenuRepo     `inject:""`
	ProjectRoleRepo  *repo.ProjectRoleRepo     `inject:""`
	UserRepo         *repo.UserRepo            `inject:""`
	RemoteService    *service.RemoteService    `inject:""`
	PrivilegeService *service.PrivilegeService `inject:""`
	RoleService      *RoleService              `inject:""`
}

func (s *ProjectMenuService) GetUserMenuList(tenantId consts.TenantId, projectId, userId uint) (ret []model.ProjectMenu, err error) {
	var roleId uint
	isAdminUser, err := s.UserRepo.IsAdminUser(tenantId, userId)
	if err != nil {
		return
	}

	if isAdminUser {
		adminProjectRole, err := s.ProjectRoleRepo.GetAdminRecord(tenantId)
		if err != nil {
			return ret, err
		}
		roleId = adminProjectRole.ID
	} else {
		projectMemberRole, err := s.ProjectRepo.FindRolesByProjectAndUser(tenantId, projectId, userId)
		if err != nil {
			return ret, err
		}
		roleId = projectMemberRole.ProjectRoleId
	}

	ret, err = s.ProjectMenuRepo.GetRoleMenuList(tenantId, roleId)
	return
}

func (s *ProjectMenuService) GetAll(tenantId consts.TenantId, userId, projectRoleId uint, needSysAuth bool) (ret []string, err error) {
	if !needSysAuth {
		ret, err = s.RoleService.GetAuthByEnv(tenantId, userId)
		if err != nil {
			return
		}
	}

	projectRoleMenus, err := s.ProjectMenuRepo.GetRoleMenuCodeList(tenantId, projectRoleId)
	if err != nil {
		return
	}

	ret = append(ret, projectRoleMenus...)
	ret = commonUtils.ArrayUnique(ret)

	return
}

func (s *ProjectMenuService) GetUserMenuListNew(tenantId consts.TenantId, projectId, userId uint, userName string, needSysAuth bool) (ret []string, err error) {
	isAdminUser, err := s.UserRepo.IsAdminUser(tenantId, userId)
	if err != nil {
		return
	}

	var projectRole model.ProjectRole
	if isAdminUser {
		projectRole, err = s.ProjectRoleRepo.GetAdminRecord(tenantId)
	} else {
		projectRole, err = s.ProjectRoleRepo.GetRoleByProjectAndUser(tenantId, projectId, userId)
	}
	if err != nil {
		return
	}

	if config.CONFIG.System.SysEnv == "ly" && !isAdminUser {
		ret, err = s.PrivilegeService.GetAll(tenantId, userName, string(projectRole.Name), needSysAuth)
	} else {
		ret, err = s.GetAll(tenantId, userId, projectRole.ID, needSysAuth)
	}

	//if config.CONFIG.System.SysEnv == "ly" {
	//	project, err := s.ProjectRepo.Get(projectId)
	//	if err != nil {
	//		return ret, err
	//	}
	//	ret, err = s.RemoteService.GetUserButtonPermissions(userName, project.ShortName)
	//} else {
	//	var roleId uint
	//	isAdminUser, err := s.UserRepo.IsAdminUser(userId)
	//	if err != nil {
	//		return ret, err
	//	}
	//
	//	if isAdminUser {
	//		adminProjectRole, err := s.ProjectRoleRepo.GetAdminRecord()
	//		if err != nil {
	//			return ret, err
	//		}
	//		roleId = adminProjectRole.ID
	//	} else {
	//		projectMemberRole, err := s.ProjectRepo.FindRolesByProjectAndUser(projectId, userId)
	//		if err != nil {
	//			return ret, err
	//		}
	//		roleId = projectMemberRole.ProjectRoleId
	//	}
	//
	//	ret, err = s.ProjectMenuRepo.GetRoleMenuCodeList(roleId)
	//}

	return
}
