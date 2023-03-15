package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SummaryDetailsService struct {
	SummaryDetailsRepo *repo.SummaryDetailsRepo `inject:""`
}

func NewSummaryDetailsService() *SummaryDetailsService {
	return &SummaryDetailsService{}
}

func (s *SummaryDetailsService) Card(projectId int64) (res v1.ResSummaryBugs, err error) {
	return
}

func (s *SummaryDetailsService) Details(projectId int64) (res v1.ResSummaryBugs, err error) {
	return
}
func (s *SummaryDetailsService) CreateByDate(req model.SummaryDetails) (err error) {
	return s.SummaryDetailsRepo.CreateByDate(req)
}

func (s *SummaryDetailsService) UpdateColumnsByDate(req model.SummaryDetails, startTime string, endTime string) (err error) {
	return s.SummaryDetailsRepo.UpdateColumnsByDate(req, startTime, endTime)
}

func (s *SummaryDetailsService) Count() (count int64, err error) {
	return s.SummaryDetailsRepo.Count()
}

func (s *SummaryDetailsService) CountByUserId(userId int64) (count int64, err error) {
	return s.SummaryDetailsRepo.CountByUserId(userId)
}

func (s *SummaryDetailsService) FindProjectIdsByUserId(userId int64) (count []int64, err error) {
	return s.SummaryDetailsRepo.FindProjectIdsByUserId(userId)
}

func (s *SummaryDetailsService) FindByProjectId(projectId int64) (summaryDetail model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.FindByProjectId(projectId)
}

func (s *SummaryDetailsService) Find() (details []model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.Find()
}

func (s *SummaryDetailsService) FindByProjectIds(projectIds []int64) (details []model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.FindByProjectIds(projectIds)
}

func (s *SummaryDetailsService) SummaryCard() (summaryCardTotal []model.SummaryCardTotal, err error) {
	return s.SummaryDetailsRepo.SummaryCard()
}

func (s *SummaryDetailsService) SummaryCardByDate(startTime string, endTime string) (summaryDetails []model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.SummaryCardByDate(startTime, endTime)
}

func (s *SummaryDetailsService) FindByProjectIdAndDate(startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.FindByProjectIdAndDate(startTime, endTime, projectId)
}
