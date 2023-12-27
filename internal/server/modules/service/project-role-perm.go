package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
)

type ProjectRolePermService struct {
	ProjectRepo         *repo.ProjectRepo         `inject:""`
	ProjectRolePermRepo *repo.ProjectRolePermRepo `inject:""`
	ProjectRoleRepo     *repo.ProjectRoleRepo     `inject:""`
	ProjectRoleMenuRepo *repo.ProjectRoleMenuRepo `inject:""`
	ProfileRepo         *repo.ProfileRepo         `inject:""`
	RemoteService       *RemoteService            `inject:""`
}

func (s *ProjectRolePermService) AllRoleList() (data []model.ProjectRole, err error) {
	if config.CONFIG.System.SysEnv == "ly" {
		err = s.ComplementRole()
		if err != nil {
			return
		}
	}

	return s.ProjectRoleRepo.AllRoleList()
}

func (s *ProjectRolePermService) ComplementRole() (err error) {
	spaceRoles, err := s.RemoteService.GetSpaceRoles()
	if err != nil {
		return
	}

	var allRoleArr []string
	roleValueMap := make(map[string]v1.SpaceRole)
	for _, v := range spaceRoles {
		allRoleArr = append(allRoleArr, v.RoleValue)
		roleValueMap[v.RoleValue] = v
	}

	notExistedRoles, err := s.GetRolesNotExisted(allRoleArr)
	if err != nil || len(notExistedRoles) == 0 {
		return
	}

	err = s.BatchCreateSpaceRole(roleValueMap, notExistedRoles)
	return
}

func (s *ProjectRolePermService) GetRolesNotExisted(allRoleArr []string) (notExistedRoles []string, err error) {
	existedRoles, err := s.ProjectRoleRepo.GetRoleNamesByNames(allRoleArr)
	if err != nil {
		return
	}

	notExistedRoles = _commUtils.Difference(allRoleArr, existedRoles)

	return
}

func (s *ProjectRolePermService) BatchCreateSpaceRole(roleValueMap map[string]v1.SpaceRole, notExistedRoles []string) (err error) {
	var roleNeedCreate []model.ProjectRole
	for _, v := range notExistedRoles {
		if roleValue, ok := roleValueMap[v]; ok {
			projectRole := model.ProjectRole{
				Name:        consts.RoleType(roleValue.RoleValue),
				DisplayName: roleValue.RoleName,
			}
			roleNeedCreate = append(roleNeedCreate, projectRole)
		}
	}

	err = s.ProjectRoleRepo.BatchCreate(roleNeedCreate)

	return
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
