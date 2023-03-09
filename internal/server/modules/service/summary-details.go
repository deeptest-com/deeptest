package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SummaryDetailsService struct {
	SummaryDetailsRepo *repo.SummaryDetailsRepo `inject:""`
}

func NewSummaryDetailsService() *SummaryDetailsService {
	return &SummaryDetailsService{}
}

func (s *SummaryDetailsService) Find() (summaryDetails []model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.Find()
}

func (s *SummaryDetailsService) FindByProjectId(projectId uint) (summaryDetails []model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.FindByProjectId(projectId)
}

func (s *SummaryDetailsService) UpdateByProjectId(projectId uint, key string, value string) (err error) {
	//return s.SummaryDetailsRepo.UpdateByProjectId(projectId, key, value)
	return nil
}

func (s *SummaryDetailsService) Create(summaryDetails model.SummaryDetails) (err error) {
	return s.SummaryDetailsRepo.Create(summaryDetails)
}
