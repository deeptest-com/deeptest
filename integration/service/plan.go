package service

import (
	integrationDomain "github.com/deeptest-com/deeptest/integration/domain"
	"github.com/deeptest-com/deeptest/integration/thirdparty/common"
	thirdparty "github.com/deeptest-com/deeptest/integration/thirdparty/service"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type PlanService struct {
	PlanRepo      *repo.PlanRepo            `inject:""`
	RemoteService *thirdparty.RemoteService `inject:""`
}

func (s *PlanService) SyncPlan(tenantId consts.TenantId, id uint) (err error) {
	plan, err := s.PlanRepo.Get(tenantId, id)
	if err != nil || !plan.IsLy {
		return err
	}
	data := integrationDomain.SyncPlan{
		Number:   plan.SerialNumber,
		Name:     plan.Name,
		IsDelete: plan.Deleted,
	}

	data.Status, _ = common.PlanStatus[string(plan.Status)]

	err = s.RemoteService.SyncPlan(tenantId, data)
	return

}
