package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	thirdparty "github.com/aaronchen2k/deeptest/integration/thirdparty/service"
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
	RemoteService       *thirdparty.RemoteService `inject:""`
}

func (s *ProjectRolePermService) AllRoleList(tenantId consts.TenantId) (data []model.ProjectRole, err error) {
	if config.CONFIG.System.SysEnv == "ly" {
		data, err = s.GetRoleFromOther(tenantId)
	} else {
		data, err = s.ProjectRoleRepo.AllRoleList(tenantId)
	}

	return
}

func (s *ProjectRolePermService) GetRoleFromOther(tenantId consts.TenantId) (data []model.ProjectRole, err error) {
	spaceRoles, err := s.RemoteService.GetSpaceRoles(tenantId)
	if err != nil {
		return
	}

	//spaceRoles = s.DealSpaceRoles(spaceRoles)

	allRoleArr, spaceRoleValueMap, err := s.GetAllRoleValueMap(spaceRoles)
	if err != nil {
		return
	}

	err = s.ComplementRoleFromOther(tenantId, allRoleArr, spaceRoleValueMap)
	if err != nil {
		return
	}

	data = s.GetRoleListFromOther(tenantId, spaceRoles)

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

func (s *ProjectRolePermService) GetRoleListMap(tenantId consts.TenantId) (res map[consts.RoleType]model.ProjectRole, err error) {
	res = make(map[consts.RoleType]model.ProjectRole)
	roleList, err := s.ProjectRoleRepo.AllRoleList(tenantId)
	if err != nil {
		return
	}

	for _, v := range roleList {
		res[v.Name] = v
	}

	return
}
func (s *ProjectRolePermService) GetRoleListFromOther(tenantId consts.TenantId, spaceRoles []integrationDomain.SpaceRole) (data []model.ProjectRole) {
	roleMap, _ := s.GetRoleListMap(tenantId)
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

func (s *ProjectRolePermService) GetRolesNotExisted(tenantId consts.TenantId, allRoleArr []string) (notExistedRoles []string, err error) {
	existedRoles, err := s.ProjectRoleRepo.GetRoleNamesByNames(tenantId, allRoleArr)
	if err != nil {
		return
	}

	notExistedRoles = _commUtils.Difference(allRoleArr, existedRoles)

	return
}

func (s *ProjectRolePermService) ComplementRoleFromOther(tenantId consts.TenantId, allRoleArr []string, roleValueMap map[string]integrationDomain.SpaceRole) (err error) {
	notExistedRoles, err := s.GetRolesNotExisted(tenantId, allRoleArr)
	if err != nil || len(notExistedRoles) == 0 {
		return
	}

	err = s.BatchCreateSpaceRole(tenantId, roleValueMap, notExistedRoles)

	return
}

func (s *ProjectRolePermService) BatchCreateSpaceRole(tenantId consts.TenantId, roleValueMap map[string]integrationDomain.SpaceRole, notExistedRoles []string) (err error) {
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

	err = s.ProjectRoleRepo.BatchCreate(tenantId, roleNeedCreate)

	return
}

func (s *ProjectRolePermService) GetProjectUserRole(tenantId consts.TenantId, userId, projectId uint) (data model.ProjectRole, err error) {
	return s.ProjectRoleRepo.ProjectUserRoleList(tenantId, userId, projectId)
}

func (s *ProjectRolePermService) PaginateRolePerms(tenantId consts.TenantId, req v1.ProjectRolePermPaginateReq) (ret _domain.PageData, err error) {
	return s.ProjectRolePermRepo.PaginateRolePerms(tenantId, req)
}

func (s *ProjectRolePermService) PaginateUserPerms(tenantId consts.TenantId, req v1.ProjectUserPermsPaginate, userId uint) (ret _domain.PageData, err error) {
	return s.ProjectRolePermRepo.UserPermList(tenantId, req, userId)
}
