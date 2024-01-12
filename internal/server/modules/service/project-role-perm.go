package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	"github.com/aaronchen2k/deeptest/integration/service"
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
	RemoteService       *service.RemoteService    `inject:""`
}

func (s *ProjectRolePermService) AllRoleList() (data []model.ProjectRole, err error) {
	if config.CONFIG.System.SysEnv == "ly" {
		data, err = s.GetRoleFromOther()
	} else {
		data, err = s.ProjectRoleRepo.AllRoleList()
	}

	return
}

func (s *ProjectRolePermService) GetRoleFromOther() (data []model.ProjectRole, err error) {
	spaceRoles, err := s.RemoteService.GetSpaceRoles()
	if err != nil {
		return
	}

	spaceRoles = s.DealSpaceRoles(spaceRoles)

	allRoleArr, spaceRoleValueMap, err := s.GetAllRoleValueMap(spaceRoles)
	if err != nil {
		return
	}

	err = s.ComplementRoleFromOther(allRoleArr, spaceRoleValueMap)
	if err != nil {
		return
	}

	data = s.GetRoleListFromOther(spaceRoles)

	return
}

func (s *ProjectRolePermService) DealSpaceRoles(spaceRoles []integrationDomain.SpaceRole) (res []integrationDomain.SpaceRole) {
	//TODO 临时角色，后期要改
	tmpSpaceRoleArr := []string{"api-admin", "general", "space-test-engineer", "space-server-engineer", "space-web-engineer", "space-product-manager"}
	for _, v := range spaceRoles {
		if _commUtils.InArray(v.RoleValue, tmpSpaceRoleArr) {
			res = append(res, v)
		}
		//if v.RoleValue != "space-visitor" {
		//	res = append(res, v)
		//}
	}

	return
}

func (s *ProjectRolePermService) GetRoleListMap() (res map[consts.RoleType]model.ProjectRole, err error) {
	res = make(map[consts.RoleType]model.ProjectRole)
	roleList, err := s.ProjectRoleRepo.AllRoleList()
	if err != nil {
		return
	}

	for _, v := range roleList {
		res[v.Name] = v
	}

	return
}
func (s *ProjectRolePermService) GetRoleListFromOther(spaceRoles []integrationDomain.SpaceRole) (data []model.ProjectRole) {
	roleMap, _ := s.GetRoleListMap()
	for _, spaceRole := range spaceRoles {
		projectRoleTmp := model.ProjectRole{
			Name:        consts.RoleType(spaceRole.RoleValue),
			DisplayName: spaceRole.RoleName,
			Description: spaceRole.Remark,
		}
		if role, ok := roleMap[consts.RoleType(spaceRole.RoleValue)]; ok {
			projectRoleTmp.ID = role.ID
		}
		data = append(data, projectRoleTmp)
	}

	return
}

func (s *ProjectRolePermService) GetAllRoleValueMap(spaceRoles []integrationDomain.SpaceRole) (allRoleArr []string, roleValueMap map[string]integrationDomain.SpaceRole, err error) {
	roleValueMap = make(map[string]integrationDomain.SpaceRole)
	for _, v := range spaceRoles {
		allRoleArr = append(allRoleArr, v.RoleValue)
		roleValueMap[v.RoleValue] = v
	}

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

func (s *ProjectRolePermService) ComplementRoleFromOther(allRoleArr []string, roleValueMap map[string]integrationDomain.SpaceRole) (err error) {
	notExistedRoles, err := s.GetRolesNotExisted(allRoleArr)
	if err != nil || len(notExistedRoles) == 0 {
		return
	}

	err = s.BatchCreateSpaceRole(roleValueMap, notExistedRoles)

	return
}

func (s *ProjectRolePermService) BatchCreateSpaceRole(roleValueMap map[string]integrationDomain.SpaceRole, notExistedRoles []string) (err error) {
	var roleNeedCreate []model.ProjectRole
	for _, v := range notExistedRoles {
		if roleValue, ok := roleValueMap[v]; ok {
			projectRole := model.ProjectRole{
				Name:        consts.RoleType(roleValue.RoleValue),
				DisplayName: roleValue.RoleName,
				Description: roleValue.Remark,
				Source:      consts.RoleSourceLy,
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
