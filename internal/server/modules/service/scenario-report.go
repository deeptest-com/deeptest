package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioReportService struct {
	ScenarioReportRepo *repo2.ScenarioReportRepo `inject:""`
	LogRepo            *repo2.LogRepo            `inject:""`
}

func NewReportService() *ScenarioReportService {
	return &ScenarioReportService{}
}

func (s *ScenarioReportService) Paginate(req v1.ReportReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioReportRepo.Paginate(req, projectId)
	return
}

func (s *ScenarioReportService) GetById(id uint) (report model.ScenarioReport, err error) {
	report, err = s.ScenarioReportRepo.Get(id)
	return
}

func (s *ScenarioReportService) DeleteById(id uint) error {
	return s.ScenarioReportRepo.DeleteById(id)
}
