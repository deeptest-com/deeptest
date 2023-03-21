package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"strconv"
	"time"
)

type SummaryProjectUserRankingService struct {
	SummaryProjectUserRankingRepo *repo.SummaryProjectUserRankingRepo `inject:""`
}

func NewSummaryProjectUserRankingService() *SummaryProjectUserRankingService {
	return &SummaryProjectUserRankingService{}
}

func (s *SummaryProjectUserRankingService) ProjectUserRanking(projectId int64, cycle int64) (resRankingList v1.ResRankingList, err error) {
	switch cycle {
	case 0:
		cycle = -7
	case 1:
		cycle = -30
	}

	date := time.Now().AddDate(0, 0, int(cycle))
	year, month, day := date.Date()
	startTime := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 00:00:00"
	endTime := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"

	summaryProjectUserRankings, err := s.FindByProjectId(projectId)
	oldSummaryProjectUserRankings, err := s.FindByDateAndProjectId(startTime, endTime, projectId)

	for _, newRanking := range summaryProjectUserRankings {
		var resUserRanking v1.ResUserRanking
		for _, oldRanking := range oldSummaryProjectUserRankings {
			if oldRanking.CreatedAt != nil {
				if oldRanking.UserId == newRanking.UserId {
					resUserRanking.UserName = newRanking.UserName
					resUserRanking.UserId = newRanking.UserId
					resUserRanking.ScenarioTotal = newRanking.ScenarioTotal
					resUserRanking.TestcasesTotal = newRanking.TestcasesTotal
					resUserRanking.UpdateTime = newRanking.UpdatedAt
					resUserRanking.Sort = newRanking.Sort
					resUserRanking.Hb = oldRanking.Sort - newRanking.Sort
				}
			} else {
				resUserRanking.UserName = newRanking.UserName
				resUserRanking.UserId = newRanking.UserId
				resUserRanking.ScenarioTotal = newRanking.ScenarioTotal
				resUserRanking.TestcasesTotal = newRanking.TestcasesTotal
				resUserRanking.UpdateTime = newRanking.UpdatedAt
				resUserRanking.Sort = newRanking.Sort
				resUserRanking.Hb = 0
			}
		}
		resRankingList.UserRankingList = append(resRankingList.UserRankingList, resUserRanking)
	}
	return
}

func (s *SummaryProjectUserRankingService) Create(summaryProjectUserRanking model.SummaryProjectUserRanking) (err error) {
	return s.SummaryProjectUserRankingRepo.Create(summaryProjectUserRanking)
}

func (s *SummaryProjectUserRankingService) FindByProjectId(projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	return s.SummaryProjectUserRankingRepo.FindByProjectId(projectId)
}

func (s *SummaryProjectUserRankingService) FindByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	return s.SummaryProjectUserRankingRepo.FindByDateAndProjectId(startTime, endTime, projectId)
}
