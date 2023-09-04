package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type PlanReportService struct {
	ReportRepo *repo2.PlanReportRepo `inject:""`
	LogRepo    *repo2.LogRepo        `inject:""`
}

func (s *PlanReportService) Paginate(req v1.PlanReportReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ReportRepo.Paginate(req, projectId)
	return
}

func (s *PlanReportService) GetById(id uint) (report model.PlanReportDetail, err error) {
	report, err = s.ReportRepo.Get(id)
	return
}

func (s *PlanReportService) DeleteById(id uint) error {
	return s.ReportRepo.DeleteById(id)
}
