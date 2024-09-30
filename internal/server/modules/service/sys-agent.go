package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type SysAgentService struct {
	SysAgentRepo *repo.SysAgentRepo `inject:""`
}

func (s *SysAgentService) List(tenantId consts.TenantId, keywords string) (pos []model.SysAgent, err error) {
	pos, err = s.SysAgentRepo.List(tenantId, keywords)

	return
}

func (s *SysAgentService) Get(tenantId consts.TenantId, id uint) (po model.SysAgent, err error) {
	po, err = s.SysAgentRepo.Get(tenantId, id)

	return
}

func (s *SysAgentService) Save(tenantId consts.TenantId, req *model.SysAgent) (err error) {
	err = s.SysAgentRepo.Save(tenantId, req)
	return
}

func (s *SysAgentService) UpdateName(tenantId consts.TenantId, req v1.AgentReq) (err error) {
	err = s.SysAgentRepo.UpdateName(tenantId, req)
	return
}

func (s *SysAgentService) Delete(tenantId consts.TenantId, id uint) (err error) {
	return s.SysAgentRepo.Delete(tenantId, id)
}

func (s *SysAgentService) Disable(tenantId consts.TenantId, id uint) (err error) {
	return s.SysAgentRepo.Disable(tenantId, id)
}
