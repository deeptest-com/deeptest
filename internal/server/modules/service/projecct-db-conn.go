package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type DatabaseConnService struct {
	DatabaseConnRepo *repo.DatabaseConnRepo `inject:""`
}

func (s *DatabaseConnService) List(tenantId consts.TenantId, keywords string, projectId int, ignoreDisabled bool) (ret []model.DatabaseConn, err error) {
	ret, err = s.DatabaseConnRepo.List(tenantId, keywords, projectId, ignoreDisabled)
	return
}

func (s *DatabaseConnService) Get(tenantId consts.TenantId, id uint) (model.DatabaseConn, error) {
	return s.DatabaseConnRepo.Get(tenantId, id)
}

func (s *DatabaseConnService) Save(tenantId consts.TenantId, po *model.DatabaseConn) (err error) {
	return s.DatabaseConnRepo.Save(tenantId, po)
}

func (s *DatabaseConnService) Delete(tenantId consts.TenantId, id uint) (err error) {
	return s.DatabaseConnRepo.Delete(tenantId, id)
}

func (s *DatabaseConnService) Disable(tenantId consts.TenantId, id uint) (err error) {
	return s.DatabaseConnRepo.Disable(tenantId, id)
}

func (s *DatabaseConnService) UpdateName(tenantId consts.TenantId, req v1.DbConnReq) (err error) {
	err = s.DatabaseConnRepo.UpdateName(tenantId, req)

	return
}
