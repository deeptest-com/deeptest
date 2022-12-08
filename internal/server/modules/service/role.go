package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type RoleService struct {
	RoleRepo *repo.RoleRepo `inject:""`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// Paginate
func (s *RoleService) Paginate(req v1.RoleReqPaginate) (ret _domain.PageData, err error) {
	return s.RoleRepo.Paginate(req)
}

// FindByName
func (s *RoleService) FindByName(name string, ids ...uint) (v1.RoleResp, error) {
	return s.RoleRepo.FindByName(name, ids...)
}

func (s *RoleService) Create(req v1.RoleReq) (uint, error) {
	return s.RoleRepo.Create(req)
}

func (s *RoleService) Update(id uint, req v1.RoleReq) error {
	return s.RoleRepo.Update(id, req)
}

func (s *RoleService) IsAdminRole(id uint) (bool, error) {
	return s.RoleRepo.IsAdminRole(id)
}

func (s *RoleService) FindById(id uint) (v1.RoleResp, error) {
	return s.RoleRepo.FindById(id)
}

func (s *RoleService) DeleteById(id uint) error {
	return s.RoleRepo.DeleteById(id)
}

func (s *RoleService) FindInId(ids []string) ([]v1.RoleResp, error) {
	return s.RoleRepo.FindInId(ids)
}

// AddPermForRole
func (s *RoleService) AddPermForRole(id uint, perms [][]string) error {
	return s.RoleRepo.AddPermForRole(id, perms)
}

func (s *RoleService) GetRoleIds() ([]uint, error) {
	return s.RoleRepo.GetRoleIds()
}
