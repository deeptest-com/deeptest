package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type ScenarioReportService struct {
	ScenarioReportRepo *repo2.ScenarioReportRepo `inject:""`
	LogRepo            *repo2.LogRepo            `inject:""`
	PlanReportRepo     *repo2.PlanReportRepo     `inject:""`
}

func NewReportService() *ScenarioReportService {
	return &ScenarioReportService{}
}

func (s *ScenarioReportService) Paginate(req v1.ReportReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioReportRepo.Paginate(req)
	return
}

func (s *ScenarioReportService) GetById(id uint) (report model.ScenarioReport, err error) {
	report, err = s.ScenarioReportRepo.Get(id)
	return
}

func (s *ScenarioReportService) DeleteById(id uint) error {
	return s.ScenarioReportRepo.DeleteById(id)
}

func (s *ScenarioReportService) CreatePlanReport(id uint) (err error) {
	var report model.ScenarioReport
	var planReport model.PlanReport
	report, err = s.ScenarioReportRepo.Get(id)
	copier.CopyWithOption(&planReport, report, copier.Option{DeepCopy: true})
	planReport.ID = 0
	err = s.PlanReportRepo.Create(&planReport)
	if err != nil {
		return
	}
	err = s.ScenarioReportRepo.UpdatePlanReportId(id, planReport.ID)

	return
}
