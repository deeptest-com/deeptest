package service

import (
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	leyan "github.com/aaronchen2k/deeptest/integration/leyan/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type PlanService struct {
	PlanRepo      *repo.PlanRepo       `inject:""`
	RemoteService *leyan.RemoteService `inject:""`
}

func (s *PlanService) SyncPlan(tenantId consts.TenantId, id uint) (err error) {
	plan, err := s.PlanRepo.Get(tenantId, id)
	if err != nil || !plan.IsLy {
		return err
	}
	data := integrationDomain.SyncPlan{
		Number:   plan.SerialNumber,
		Name:     plan.Name,
		Status:   1,
		IsDelete: plan.Deleted,
	}
	err = s.RemoteService.SyncPlan(tenantId, data)
	return

}
