package service

import (
	integrationDomain "github.com/deeptest-com/deeptest/integration/domain"
	thirdparty "github.com/deeptest-com/deeptest/integration/thirdparty/service"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type ReportService struct {
	PlanReportRepo *repo.PlanReportRepo      `inject:""`
	RemoteService  *thirdparty.RemoteService `inject:""`
	PlanRepo       *repo.PlanRepo            `inject:""`
	UserRepo       *repo.UserRepo            `inject:""`
}

func (s *ReportService) SaveReport(tenantId consts.TenantId, id uint) (err error) {
	report, err := s.PlanReportRepo.Get(tenantId, id)
	if err != nil {
		return err
	}

	// 保存到thirdparty
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
	data.EnvName = report.ExecEnv
	data.ExecUserName = report.ExecUserName
	data.TestRate = report.TestRate
	err = s.RemoteService.SaveReport(tenantId, data)

	return

}

func (s *ReportService) DeleteReport(tenantId consts.TenantId, id uint) (err error) {
	report, err := s.PlanReportRepo.Get(tenantId, id)
	if err != nil {
		return err
	}
	err = s.RemoteService.DeleteReport(tenantId, report.SerialNumber)
	return
}
