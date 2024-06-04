package service

import (
	integrationDomain "github.com/aaronchen2k/deeptest/integration/domain"
	leyan "github.com/aaronchen2k/deeptest/integration/leyan/service"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ReportService struct {
	PlanReportRepo *repo.PlanReportRepo `inject:""`
	RemoteService  *leyan.RemoteService `inject:""`
	PlanRepo       *repo.PlanRepo       `inject:""`
	UserRepo       *repo.UserRepo       `inject:""`
}

func (s *ReportService) SaveReport(tenantId consts.TenantId, id uint) (err error) {
	report, err := s.PlanReportRepo.Get(tenantId, id)
	if err != nil {
		return err
	}

	// 保存到leyan
	data := integrationDomain.CreateReport{}
	data.Name = report.Name
	data.ApiReportNumber = report.SerialNumber

	plan, err := s.PlanRepo.Get(tenantId, report.PlanId)
	if err != nil || !plan.IsLy {
		return err
	}
	data.ApiPlanNumber = plan.SerialNumber

	user, err := s.UserRepo.GetByUserId(tenantId, report.CreateUserId)
	if err != nil {
		return err
	}
	data.CreatedBy = user.Username

	s.RemoteService.SaveReport(tenantId, data)

	return

}
