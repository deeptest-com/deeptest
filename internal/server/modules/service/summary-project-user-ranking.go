package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"time"
)

type SummaryProjectUserRankingService struct {
	SummaryProjectUserRankingRepo *repo.SummaryProjectUserRankingRepo `inject:""`
}

func NewSummaryProjectUserRankingService() *SummaryProjectUserRankingService {
	return new(SummaryProjectUserRankingService)
}

func (s *SummaryProjectUserRankingService) ProjectUserRanking(projectId int64, cycle int64) (resRankingList v1.ResRankingList, err error) {

	//获取即时数据
	newRankings, _ := s.GetRanking(projectId)

	//查询所有用户名字，map的key为userId
	userInfo, _ := s.FindAllUserName()
	var lastWeekRanking map[int64]model.SummaryProjectUserRanking
	lastWeekRanking = make(map[int64]model.SummaryProjectUserRanking)

	for _, newRanking := range newRankings {
		var resRanking v1.ResUserRanking
		var testTotal int64
		var scenarioTotal int64

		if cycle == 1 {
			//全部范围数据
			//查询上周的数据
			lastWeek := time.Now().AddDate(0, 0, -7)
			lastWeekStartTime, lastWeekEndTime := GetDate(lastWeek)
			lastWeekRanking, err = s.FindByDateAndProjectIdOfMap(lastWeekStartTime, lastWeekEndTime, projectId)
			scenarioTotal = newRanking.ScenarioTotal
			testTotal = newRanking.TestCaseTotal
		} else if cycle == 0 {
			//当月范围数据
			lastMonthLastDay := time.Now().AddDate(0, 0, -30)
			lastMonthLastDayStartTime, lastMonthLastDayEndTime := GetDate(lastMonthLastDay)
			lastMonthLastDayRanking, _ := s.FindByDateAndProjectIdOfMap(lastMonthLastDayStartTime, lastMonthLastDayEndTime, projectId)
			scenarioTotal = newRanking.ScenarioTotal - lastMonthLastDayRanking[newRanking.UserId].ScenarioTotal
			testTotal = newRanking.TestCaseTotal - lastMonthLastDayRanking[newRanking.UserId].TestCaseTotal
		}
		if lastWeekRanking[newRanking.UserId].Sort != 0 {
			resRanking.Hb = lastWeekRanking[newRanking.UserId].Sort - newRanking.Sort
		}
		resRanking.ScenarioTotal = scenarioTotal
		resRanking.TestCaseTotal = testTotal
		resRanking.UserName = userInfo[newRanking.UserId]
		resRanking.UpdatedAt = newRanking.UpdatedAt.Format("2006-01-02 15:04:05")
		resRanking.UserId = newRanking.UserId
		resRankingList.UserRankingList = append(resRankingList.UserRankingList, resRanking)
	}

	//由于存在当月选项，当月数据需要重新进行排序，不累积
	if len(resRankingList.UserRankingList) != 0 {
		resRankingList, _ = s.SortRankingList(resRankingList)
	}

	return
}

func (s *SummaryProjectUserRankingService) Create(req model.SummaryProjectUserRanking) (err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.Create(req)
}

func (s *SummaryProjectUserRankingService) CreateByDate(req model.SummaryProjectUserRanking) (err error) {
	now := time.Now()
	startTime, endTime := GetDate(now)
	id, err := s.Existed(startTime, endTime, req.ProjectId, req.UserId)
	if id == 0 {
		err = s.Create(req)
	} else {
		err = s.UpdateColumnsByDate(id, req)
	}
	return
}

func (s *SummaryProjectUserRankingService) UpdateColumnsByDate(id int64, req model.SummaryProjectUserRanking) (err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.UpdateColumnsByDate(id, req)
}

func (s *SummaryProjectUserRankingService) FindProjectIds() (projectIds []int64, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.FindProjectIds()
}

func (s *SummaryProjectUserRankingService) Existed(startTime string, endTiem string, projectId int64, userId int64) (id int64, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.Existed(startTime, endTiem, projectId, userId)
}

func (s *SummaryProjectUserRankingService) FindByProjectId(projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.FindByProjectId(projectId)
}

func (s *SummaryProjectUserRankingService) FindByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	return r.FindByDateAndProjectId(startTime, endTime, projectId)
}

func (s *SummaryProjectUserRankingService) FindByDateAndProjectIdOfMap(startTime string, endTime string, projectId int64) (result map[int64]model.SummaryProjectUserRanking, err error) {
	summaryProjectUserRanking, _ := s.FindByDateAndProjectId(startTime, endTime, projectId)

	result = make(map[int64]model.SummaryProjectUserRanking, len(summaryProjectUserRanking))
	for _, ranking := range summaryProjectUserRanking {
		result[ranking.UserId] = ranking
	}
	return
}

func (s *SummaryProjectUserRankingService) CheckUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	r := *repo.NewSummaryProjectUserRankingRepo()
	return r.CheckUpdated(lastUpdateTime)
}

func (s *SummaryProjectUserRankingService) FindScenarioTotalOfUserGroupByProject() (ScenariosTotal map[int64][]model.ProjectUserTotal, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	results, err := r.FindProjectUserScenarioTotal()
	ScenariosTotal = make(map[int64][]model.ProjectUserTotal, len(results))

	for _, result := range results {
		ScenariosTotal[result.ProjectId] = append(ScenariosTotal[result.ProjectId], result)
	}
	return
}

func (s *SummaryProjectUserRankingService) FindTestCasesTotalOfUserGroupByProject() (testCasesTotal map[int64][]model.ProjectUserTotal, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	results, err := r.FindProjectUserTestCasesTotal()
	testCasesTotal = make(map[int64][]model.ProjectUserTotal, len(results))

	for _, result := range results {
		testCasesTotal[result.ProjectId] = append(testCasesTotal[result.ProjectId], result)
	}
	return
}

func (s *SummaryProjectUserRankingService) FindCasesTotalByProjectId(projectId int64) (result map[int64]int64, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	counts, err := r.FindCasesTotalByProjectId(projectId)
	result = make(map[int64]int64, len(counts))
	for _, tmp := range counts {
		result[tmp.CreateUserId] = tmp.Count
	}

	return
}

func (s *SummaryProjectUserRankingService) FindScenariosTotalByProjectId(projectId int64) (result map[int64]int64, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	counts, err := r.FindScenariosTotalByProjectId(projectId)
	result = make(map[int64]int64, len(counts))
	for _, tmp := range counts {
		result[tmp.CreateUserId] = tmp.Count
	}

	return
}

func (s *SummaryProjectUserRankingService) FindUserLastUpdateTestCasesByProjectId(projectId int64) (result map[int64]*time.Time, err error) {
	r := repo.NewSummaryProjectUserRankingRepo()
	updateTime, err := r.FindUserLastUpdateTestCasesByProjectId(projectId)
	result = make(map[int64]*time.Time, len(updateTime))
	for _, tmp := range updateTime {
		result[tmp.CreatedBy] = tmp.UpdatedAt
	}
	return
}

func (s *SummaryProjectUserRankingService) FindAllUserName() (result map[int64]string, err error) {
	r := *repo.NewSummaryProjectUserRankingRepo()
	users, err := r.FindAllUserName()
	result = make(map[int64]string, len(users))
	for _, user := range users {
		result[user.Id] = user.Name
	}
	return
}

func (s *SummaryProjectUserRankingService) FindUserByProjectId(projectId int64) (users []model.RankingUser, err error) {
	r := *repo.NewSummaryProjectUserRankingRepo()
	users, err = r.FindUserByProjectId(projectId)
	return
}

func (s *SummaryProjectUserRankingService) FindUserIdsByProjectId(projectId int64) (userIds []int64, err error) {
	r := *repo.NewSummaryProjectUserRankingRepo()
	userIds, err = r.FindUserIdsByProjectId(projectId)
	return
}

func (s *SummaryProjectUserRankingService) ForMap(userTotal []model.UserTotal) (ret []map[int64]int64, err error) {

	user := make(map[int64]int64, len(userTotal))
	for _, u := range userTotal {
		user[u.CreateUserId] = u.Count
		ret = append(ret, user)
	}
	return
}

func (s *SummaryProjectUserRankingService) SortRanking(data []model.SummaryProjectUserRanking) (ret []model.SummaryProjectUserRanking, err error) {
	length := len(data)
	for i := 0; i < length; i++ {
		mix := data[i].TestCaseTotal

		for x := i + 1; x < length; x++ {
			if data[x].TestCaseTotal < mix {

				tmp := data[i]
				data[i] = data[x]
				data[x] = tmp
				data[i].Sort = int64(x)
			}
		}
	}
	data[length-1].Sort = int64(length - 1)
	ret = data
	return
}

func (s *SummaryProjectUserRankingService) SortRankingList(data v1.ResRankingList) (ret v1.ResRankingList, err error) {
	list := data.UserRankingList
	length := len(list)
	for i := 0; i < length; i++ {
		mix := list[i].TestCaseTotal
		for x := i + 1; x < length; x++ {
			if list[x].TestCaseTotal < mix {
				tmp := list[i]
				list[i] = list[x]
				list[x] = tmp
				list[i].Sort = int64(x)
			}
		}
	}
	list[length-1].Sort = int64(length - 1)
	ret = data
	return
}

func (s *SummaryProjectUserRankingService) GetRanking(projectId int64) (rankings []model.SummaryProjectUserRanking, err error) {

	users, err := s.FindUserIdsByProjectId(projectId)

	cases, err := s.FindCasesTotalByProjectId(projectId)
	scenarios, err := s.FindScenariosTotalByProjectId(projectId)
	lastUpdateTime, _ := s.FindUserLastUpdateTestCasesByProjectId(projectId)

	for _, user := range users {
		var ranking model.SummaryProjectUserRanking
		ranking.UserId = user
		ranking.ProjectId = projectId
		ranking.ScenarioTotal = scenarios[user]
		ranking.TestCaseTotal = cases[user]
		ranking.UpdatedAt = lastUpdateTime[user]
		rankings = append(rankings, ranking)
	}
	if len(rankings) != 0 {
		rankings, err = s.SortRanking(rankings)
	}

	return
}

func (s *SummaryProjectUserRankingService) SaveRanking() (err error) {
	projectIds, err := s.FindProjectIds()
	for _, projectId := range projectIds {
		rankings, _ := s.GetRanking(projectId)
		for _, ranking := range rankings {
			err := s.CreateByDate(ranking)
			if err != nil {
				return err
			}
		}
	}
	return
}
