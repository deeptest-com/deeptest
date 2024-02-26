package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PermService struct {
	PermRepo *repo.PermRepo `inject:""`
}

// Paginate
func (s *PermService) Paginate(tenantId consts.TenantId, req v1.PermReqPaginate) (data _domain.PageData, err error) {
	return s.PermRepo.Paginate(tenantId, req)
}

// FindByNameAndAct
// db *gorm.DB
// name 名称
// act 方法
// ids 当 ids 的 len = 1 ，排除次 id 数据
func (s *PermService) FindByNameAndAct(tenantId consts.TenantId, name, act string, ids ...uint) (v1.PermResp, error) {
	return s.PermRepo.FindByNameAndAct(tenantId, name, act, ids...)
}

// Create
func (s *PermService) Create(tenantId consts.TenantId, req v1.PermReq) (uint, error) {
	return s.PermRepo.Create(tenantId, req)
}

// CreatenInBatches
func (s *PermService) CreatenInBatches(tenantId consts.TenantId, perms []model.SysPerm) error {
	return s.PermRepo.CreateInBatches(tenantId, perms)
}

// Update
func (s *PermService) Update(tenantId consts.TenantId, id uint, req v1.PermReq) error {
	return s.PermRepo.Update(tenantId, id, req)
}

// checkNameAndAct
func (r *PermService) checkNameAndAct(tenantId consts.TenantId, req v1.PermReq, ids ...uint) bool {
	return r.PermRepo.CheckNameAndAct(tenantId, req, ids...)
}

// FindById
func (s *PermService) FindById(tenantId consts.TenantId, id uint) (v1.PermResp, error) {
	return s.PermRepo.FindById(tenantId, id)
}

// DeleteById
func (s *PermService) DeleteById(tenantId consts.TenantId, id uint) error {
	return s.PermRepo.DeleteById(tenantId, id)
}

// GetPermsForRole
func (s *PermService) GetPermsForRole(tenantId consts.TenantId) ([][]string, error) {
	return s.PermRepo.GetPermsForRole(tenantId)
}
