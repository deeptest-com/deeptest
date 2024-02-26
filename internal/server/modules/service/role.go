package service

import (
	"encoding/json"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"io/ioutil"
)

type RoleService struct {
	RoleRepo *repo.RoleRepo `inject:""`
	UserRepo *repo.UserRepo `inject:""`
}

// Paginate
func (s *RoleService) Paginate(tenantId consts.TenantId, req v1.RoleReqPaginate) (ret _domain.PageData, err error) {
	return s.RoleRepo.Paginate(tenantId, req)
}

// FindByName
func (s *RoleService) FindByName(tenantId consts.TenantId, name string, ids ...uint) (v1.RoleResp, error) {
	return s.RoleRepo.FindByName(tenantId, name, ids...)
}

func (s *RoleService) Create(tenantId consts.TenantId, req v1.RoleReq) (uint, error) {
	return s.RoleRepo.Create(tenantId, req)
}

func (s *RoleService) Update(tenantId consts.TenantId, id uint, req v1.RoleReq) error {
	return s.RoleRepo.Update(tenantId, id, req)
}

func (s *RoleService) IsAdminRole(tenantId consts.TenantId, id uint) (bool, error) {
	return s.RoleRepo.IsAdminRole(tenantId, id)
}

func (s *RoleService) FindById(tenantId consts.TenantId, id uint) (v1.RoleResp, error) {
	return s.RoleRepo.FindById(tenantId, id)
}

func (s *RoleService) DeleteById(tenantId consts.TenantId, id uint) error {
	return s.RoleRepo.DeleteById(tenantId, id)
}

func (s *RoleService) FindInId(tenantId consts.TenantId, ids []string) ([]v1.RoleResp, error) {
	return s.RoleRepo.FindInId(tenantId, ids)
}

// AddPermForRole
func (s *RoleService) AddPermForRole(tenantId consts.TenantId, id uint, perms [][]string) error {
	return s.RoleRepo.AddPermForRole(tenantId, id, perms)
}

func (s *RoleService) GetRoleIds(tenantId consts.TenantId) ([]uint, error) {
	return s.RoleRepo.GetRoleIds(tenantId)
}

func (s *RoleService) AllRoleList(tenantId consts.TenantId) ([]v1.RoleResp, error) {
	return s.RoleRepo.GetAllRoles(tenantId)
}

func (s *RoleService) GetAuthByEnv(tenantId consts.TenantId, userId uint) (res []string, err error) {
	//if config.CONFIG.System.SysEnv != "ly" {
	user, err := s.UserRepo.FindDetailById(tenantId, userId)
	if err != nil {
		return []string{}, err
	}
	return s.GetRoleMenuConfig(user.SysRoles)
	//}

	//return []string{}, nil
}

func (s *RoleService) GetRoleMenuConfig(roles []string) (menus []string, err error) {

	data, err := ioutil.ReadFile("config/sample/sys-role-menu.json")
	if err != nil {
		logUtils.Errorf("load sys role menu config err ", zap.String("错误:", err.Error()))
		return
	}
	roleMenuConfigs := make([]v1.RoleMenuConfig, 0)
	err = json.Unmarshal(data, &roleMenuConfigs)
	if err != nil {
		logUtils.Errorf("unmarshall sys role menu config err ", zap.String("错误:", err.Error()))
		return
	}

	roleMenuConfigMap := make(map[string][]string)
	for _, v := range roleMenuConfigs {
		roleMenuConfigMap[v.RoleName] = v.Menus
	}

	for _, v := range roles {
		menus = append(menus, roleMenuConfigMap[v]...)
	}

	menus = _commUtils.ArrayRemoveDuplication(menus)
	return
}
