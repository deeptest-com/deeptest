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
	return new(SummaryProjectUserRankingService)
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
		if err == nil && oldSummaryProjectUserRankings != nil && len(oldSummaryProjectUserRankings) != 0 {
			for _, oldRanking := range oldSummaryProjectUserRankings {
				if oldRanking.UserId == newRanking.UserId {
					resUserRanking.UserName = newRanking.UserName
					resUserRanking.UserId = newRanking.UserId
					resUserRanking.ScenarioTotal = newRanking.ScenarioTotal
					resUserRanking.TestCaseTotal = newRanking.TestCaseTotal
					resUserRanking.UpdatedAt = newRanking.UpdatedAt
					resUserRanking.Sort = newRanking.Sort
					resUserRanking.Hb = oldRanking.Sort - newRanking.Sort
				}
			}
		} else {
			resUserRanking.UserName = newRanking.UserName
			resUserRanking.UserId = newRanking.UserId
			resUserRanking.ScenarioTotal = newRanking.ScenarioTotal
			resUserRanking.TestCaseTotal = newRanking.TestCaseTotal
			resUserRanking.UpdatedAt = newRanking.UpdatedAt
			resUserRanking.Sort = newRanking.Sort
			resUserRanking.Hb = 0
		}
		resRankingList.UserRankingList = append(resRankingList.UserRankingList, resUserRanking)
	}
	return
}
func (s *SummaryProjectUserRankingService) Create(req model.SummaryProjectUserRanking) (err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.Create(req)
}

func (s *SummaryProjectUserRankingService) CreateByDate(req model.SummaryProjectUserRanking) (err error) {
	now := time.Now()
	year, month, day := now.Date()
	startTime := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 00:00:00"
	endTime := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"
	ret, err := s.HasDataOfDate(startTime, endTime)
	if ret {
		err = s.Create(req)
	} else {
		err = s.UpdateColumnsByDate(req, startTime, endTime)
	}
	return
}

func (s *SummaryProjectUserRankingService) UpdateColumnsByDate(req model.SummaryProjectUserRanking, startTime string, endTime string) (err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.UpdateColumnsByDate(req, startTime, endTime)
}

func (s *SummaryProjectUserRankingService) FindProjectIds() (projectIds []int64, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.FindProjectIds()
}

func (s *SummaryProjectUserRankingService) HasDataOfDate(startTime string, endTiem string) (ret bool, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.HasDataOfDate(startTime, endTiem)
}

func (s *SummaryProjectUserRankingService) FindByProjectId(projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.FindByProjectId(projectId)
}

func (s *SummaryProjectUserRankingService) FindByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.FindByDateAndProjectId(startTime, endTime, projectId)
}

func (s *SummaryProjectUserRankingService) CheckUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	r := *repo.NewSummaryProjectUserRankingRepo()
	return r.CheckUpdated(lastUpdateTime)
}
