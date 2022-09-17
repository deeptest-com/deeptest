package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ReportService struct {
	ReportRepo *repo.ReportRepo `inject:""`
	LogRepo    *repo.LogRepo    `inject:""`
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) Paginate(req serverDomain.ReportReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ReportRepo.Paginate(req, projectId)
	return
}

func (s *ReportService) GetById(id uint) (report model.Report, err error) {
	report, err = s.ReportRepo.Get(id)
	return
}

func (s *ReportService) DeleteById(id uint) error {
	return s.ReportRepo.DeleteById(id)
}
