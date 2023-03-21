package service

import v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"

type SummaryService struct {
	SummaryDetailsService            *SummaryDetailsService            `inject:""`
	SummaryBugsService               *SummaryBugsService               `inject:""`
	SummaryProjectUserRankingService *SummaryProjectUserRankingService `inject:""`
}

func NewSummaryService() *SummaryService {
	return &SummaryService{}
}

func (s *SummaryService) Bugs(projectId int64) (res v1.ResSummaryBugs, err error) {
	res, err = s.SummaryBugsService.Bugs(projectId)
	return
}

func (s *SummaryService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	res, err = s.SummaryDetailsService.Details(userId)
	return
}

func (s *SummaryService) ProjectUserRanking(projectId int64, cycle int64) (res v1.ResRankingList, err error) {
	res, err = s.SummaryProjectUserRankingService.ProjectUserRanking(projectId, cycle)
	return
}

func (s *SummaryService) Card(projectId int64) (res v1.ResSummaryCard, err error) {
	res, err = s.SummaryDetailsService.Card(projectId)
	return
}

func (s *SummaryService) Collection() (err error) {

	return

}
