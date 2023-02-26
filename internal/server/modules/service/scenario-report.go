package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ReportService struct {
	ReportRepo *repo2.PlanReportRepo `inject:""`
	LogRepo    *repo2.LogRepo        `inject:""`
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) Paginate(req v1.ReportReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ReportRepo.Paginate(req, projectId)
	return
}

func (s *ReportService) GetById(id uint) (report model.ScenarioReport, err error) {
	report, err = s.ReportRepo.Get(id)
	return
}

func (s *ReportService) DeleteById(id uint) error {
	return s.ReportRepo.DeleteById(id)
}
