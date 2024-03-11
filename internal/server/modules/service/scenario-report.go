package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type ScenarioReportService struct {
	ScenarioReportRepo *repo2.ScenarioReportRepo `inject:""`
	LogRepo            *repo2.LogRepo            `inject:""`
	PlanReportRepo     *repo2.PlanReportRepo     `inject:""`
	UserRepo           *repo2.UserRepo           `inject:""`
	ScenarioService    *ScenarioService          `inject:""`
}

func (s *ScenarioReportService) Paginate(tenantId consts.TenantId, req v1.ReportReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioReportRepo.Paginate(tenantId, req)
	return
}

func (s *ScenarioReportService) GetById(tenantId consts.TenantId, id uint) (report model.ScenarioReport, err error) {
	report, err = s.ScenarioReportRepo.Get(tenantId, id)

	scenario, err := s.ScenarioService.GetById(tenantId, report.ScenarioId)
	if err != nil {
		return
	}
	report.CreateUserName = scenario.CreatorName

	createUser, _ := s.UserRepo.GetByUserId(tenantId, report.CreateUserId)
	report.ExecUserName = createUser.Name
	report.Priority = scenario.Priority
	return
}

func (s *ScenarioReportService) DeleteById(tenantId consts.TenantId, id uint) error {
	return s.ScenarioReportRepo.DeleteById(tenantId, id)
}

func (s *ScenarioReportService) CreatePlanReport(tenantId consts.TenantId, id uint) (err error) {
	var report model.ScenarioReport
	var planReport model.PlanReport
	report, err = s.ScenarioReportRepo.Get(tenantId, id)
	copier.CopyWithOption(&planReport, report, copier.Option{DeepCopy: true})
	planReport.ID = 0
	err = s.PlanReportRepo.Create(tenantId, &planReport)
	if err != nil {
		return
	}
	err = s.ScenarioReportRepo.UpdatePlanReportId(tenantId, id, planReport.ID)

	return
}
