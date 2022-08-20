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

func (s *ReportService) Paginate(req serverDomain.ReportReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ReportRepo.Paginate(req)
	return
}

func (s *ReportService) FindById(id uint) (model.Report, error) {
	return s.ReportRepo.Get(id)
}

func (s *ReportService) DeleteById(id uint) error {
	return s.ReportRepo.DeleteById(id)
}
