package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"sort"
	"time"
)

type SummaryProjectUserRankingService struct {
	SummaryProjectUserRankingRepo *repo.SummaryProjectUserRankingRepo `inject:""`
}

func (s *SummaryProjectUserRankingService) ProjectUserRanking(tenantId consts.TenantId, cycle int64, projectId int64) (resRankingList v1.ResRankingList, err error) {

	//获取即时数据
	newRankings, _ := s.GetRanking(tenantId, projectId)

	//查询所有用户名字，map的key为userId
	userInfo, _ := s.FindAllUserName(tenantId)
	var lastWeekRanking map[int64]model.SummaryProjectUserRanking
	lastWeekRanking = make(map[int64]model.SummaryProjectUserRanking)

	for _, newRanking := range newRankings {
		var resRanking v1.ResUserRanking
		var testTotal int64
		var scenarioTotal int64

		//查询7天前直到现在，最靠前的数据
		earlierDateStartTime, todayEndTime := GetEarlierDateUntilTodayStartAndEndTime(-7)
		lastWeekRanking, err = s.FindMinDataByDateAndProjectIdOfMap(tenantId, earlierDateStartTime, todayEndTime, projectId)

		if cycle == 1 {
			//全部范围数据,就是最新的数据，newRanking
			scenarioTotal = newRanking.ScenarioTotal
			testTotal = newRanking.TestCaseTotal
		} else if cycle == 0 {
			//当月范围数据
			//先查询31天前的数据
			earlierDateStartTime, earlierDateEndTime := GetEarlierDateStartAndEndTime(-31)
			lastMonthLastDayRanking, _ := s.FindMaxDataByDateAndProjectIdOfMap(tenantId, earlierDateStartTime, earlierDateEndTime, projectId)
			//那newRanking的所有数据，减去30天前的，就是当月增量数据情况
			scenarioTotal = newRanking.ScenarioTotal - lastMonthLastDayRanking[newRanking.UserId].ScenarioTotal
			testTotal = newRanking.TestCaseTotal - lastMonthLastDayRanking[newRanking.UserId].TestCaseTotal
		}
		if lastWeekRanking[newRanking.UserId].Sort != 0 {
			resRanking.Hb = lastWeekRanking[newRanking.UserId].Sort - newRanking.Sort
		}
		resRanking.Sort = newRanking.Sort
		resRanking.ScenarioTotal = scenarioTotal
		resRanking.TestCaseTotal = testTotal
		resRanking.UserName = userInfo[newRanking.UserId]
		lastUpdateTime, _ := s.FindUserLastUpdateTestCasesByProjectId(tenantId, projectId)
		if lastUpdateTime[newRanking.UserId] != nil {
			resRanking.UpdatedAt = lastUpdateTime[newRanking.UserId].Format("2006-01-02 15:04:05")
		} else {
			resRanking.UpdatedAt = "------"
		}
		resRanking.UserId = newRanking.UserId
		resRankingList.UserRankingList = append(resRankingList.UserRankingList, resRanking)
	}

	//由于存在当月选项，当月数据需要重新进行排序，不累积
	if len(resRankingList.UserRankingList) != 0 {
		resRankingList, _ = s.SortRankingList(resRankingList)
	}

	return
}

func (s *SummaryProjectUserRankingService) HandlerSummaryProjectUserRankingRepo() *repo.SummaryProjectUserRankingRepo {
	return repo.NewSummaryProjectUserRankingRepo()
}

func (s *SummaryProjectUserRankingService) Create(tenantId consts.TenantId, req model.SummaryProjectUserRanking) (err error) {
	return s.HandlerSummaryProjectUserRankingRepo().Create(tenantId, req)
}

func (s *SummaryProjectUserRankingService) CreateByDate(tenantId consts.TenantId, req model.SummaryProjectUserRanking) (err error) {
	startTime, endTime := GetTodayStartAndEndTime()
	id, err := s.Existed(tenantId, startTime, endTime, req.ProjectId, req.UserId)
	if id == 0 {
		err = s.Create(tenantId, req)
	} else {
		err = s.UpdateColumnsByDate(tenantId, id, req)
	}
	return
}

func (s *SummaryProjectUserRankingService) UpdateColumnsByDate(tenantId consts.TenantId, id int64, req model.SummaryProjectUserRanking) (err error) {
	return s.HandlerSummaryProjectUserRankingRepo().UpdateColumnsByDate(tenantId, id, req)
}

func (s *SummaryProjectUserRankingService) FindProjectIds(tenantId consts.TenantId) (projectIds []int64, err error) {
	return s.HandlerSummaryProjectUserRankingRepo().FindProjectIds(tenantId)
}

func (s *SummaryProjectUserRankingService) Existed(tenantId consts.TenantId, startTime string, endTiem string, projectId int64, userId int64) (id int64, err error) {
	return s.HandlerSummaryProjectUserRankingRepo().Existed(tenantId, startTime, endTiem, projectId, userId)
}

func (s *SummaryProjectUserRankingService) FindByProjectId(tenantId consts.TenantId, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {
	return s.HandlerSummaryProjectUserRankingRepo().FindByProjectId(tenantId, projectId)
}

func (s *SummaryProjectUserRankingService) FindMaxDataByDateAndProjectId(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {

	return s.HandlerSummaryProjectUserRankingRepo().FindMaxDataByDateAndProjectId(tenantId, startTime, endTime, projectId)
}

func (s *SummaryProjectUserRankingService) FindMinDataByDateAndProjectId(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (summaryProjectUserRanking []model.SummaryProjectUserRanking, err error) {

	return s.HandlerSummaryProjectUserRankingRepo().FindMinDataByDateAndProjectId(tenantId, startTime, endTime, projectId)
}

func (s *SummaryProjectUserRankingService) FindMaxDataByDateAndProjectIdOfMap(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (result map[int64]model.SummaryProjectUserRanking, err error) {
	summaryProjectUserRanking, _ := s.FindMaxDataByDateAndProjectId(tenantId, startTime, endTime, projectId)

	result = make(map[int64]model.SummaryProjectUserRanking, len(summaryProjectUserRanking))
	for _, ranking := range summaryProjectUserRanking {
		result[ranking.UserId] = ranking
	}
	return
}

func (s *SummaryProjectUserRankingService) FindMinDataByDateAndProjectIdOfMap(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (result map[int64]model.SummaryProjectUserRanking, err error) {
	summaryProjectUserRanking, _ := s.FindMinDataByDateAndProjectId(tenantId, startTime, endTime, projectId)

	result = make(map[int64]model.SummaryProjectUserRanking, len(summaryProjectUserRanking))
	for _, ranking := range summaryProjectUserRanking {
		result[ranking.UserId] = ranking
	}
	return
}

func (s *SummaryProjectUserRankingService) CheckUpdated(tenantId consts.TenantId, lastUpdateTime *time.Time) (result bool, err error) {
	return s.HandlerSummaryProjectUserRankingRepo().CheckUpdated(tenantId, lastUpdateTime)
}

func (s *SummaryProjectUserRankingService) FindScenarioTotalOfUserGroupByProject(tenantId consts.TenantId) (ScenariosTotal map[int64][]model.ProjectUserTotal, err error) {

	results, err := s.HandlerSummaryProjectUserRankingRepo().FindProjectUserScenarioTotal(tenantId)
	ScenariosTotal = make(map[int64][]model.ProjectUserTotal, len(results))

	for _, result := range results {
		ScenariosTotal[result.ProjectId] = append(ScenariosTotal[result.ProjectId], result)
	}
	return
}

func (s *SummaryProjectUserRankingService) FindTestCasesTotalOfUserGroupByProject(tenantId consts.TenantId) (testCasesTotal map[int64][]model.ProjectUserTotal, err error) {

	results, err := s.HandlerSummaryProjectUserRankingRepo().FindProjectUserTestCasesTotal(tenantId)
	testCasesTotal = make(map[int64][]model.ProjectUserTotal, len(results))

	for _, result := range results {
		testCasesTotal[result.ProjectId] = append(testCasesTotal[result.ProjectId], result)
	}
	return
}

func (s *SummaryProjectUserRankingService) FindCasesTotalByProjectId(tenantId consts.TenantId, projectId int64) (result map[int64]int64, err error) {

	counts, err := s.HandlerSummaryProjectUserRankingRepo().FindCasesTotalByProjectId(tenantId, projectId)
	result = make(map[int64]int64, len(counts))
	for _, tmp := range counts {
		result[tmp.CreateUserId] = tmp.Count
	}

	return
}

func (s *SummaryProjectUserRankingService) FindScenariosTotalByProjectId(tenantId consts.TenantId, projectId int64) (result map[int64]int64, err error) {

	counts, err := s.HandlerSummaryProjectUserRankingRepo().FindScenariosTotalByProjectId(tenantId, projectId)
	result = make(map[int64]int64, len(counts))
	for _, tmp := range counts {
		result[tmp.CreateUserId] = tmp.Count
	}

	return
}

func (s *SummaryProjectUserRankingService) FindUserLastUpdateTestCasesByProjectId(tenantId consts.TenantId, projectId int64) (result map[int64]*time.Time, err error) {

	updateTime, err := s.HandlerSummaryProjectUserRankingRepo().FindUserLastUpdateTestCasesByProjectId(tenantId, projectId)
	result = make(map[int64]*time.Time, len(updateTime))
	for _, tmp := range updateTime {
		result[tmp.CreatedBy] = tmp.UpdatedAt
	}
	return
}

func (s *SummaryProjectUserRankingService) FindAllUserName(tenantId consts.TenantId) (result map[int64]string, err error) {
	users, err := s.HandlerSummaryProjectUserRankingRepo().FindAllUserName(tenantId)
	result = make(map[int64]string, len(users))
	for _, user := range users {
		result[user.Id] = user.Name
	}
	return
}

func (s *SummaryProjectUserRankingService) FindUserByProjectId(tenantId consts.TenantId, projectId int64) (users []model.RankingUser, err error) {
	users, err = s.HandlerSummaryProjectUserRankingRepo().FindUserByProjectId(tenantId, projectId)
	return
}

func (s *SummaryProjectUserRankingService) FindUserIdsByProjectId(tenantId consts.TenantId, projectId int64) (userIds []int64, err error) {
	userIds, err = s.HandlerSummaryProjectUserRankingRepo().FindUserIdsByProjectId(tenantId, projectId)
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

type elseif struct {
}

func (s *SummaryProjectUserRankingService) SortRanking(data []model.SummaryProjectUserRanking) []model.SummaryProjectUserRanking {

	sort.Slice(data, func(i, j int) bool {
		// 1. 当data[i].ScenarioTotal 和 data[j].ScenarioTotal 不相等时候，最大的排前边
		if data[i].ScenarioTotal != data[j].ScenarioTotal {
			return data[i].ScenarioTotal > data[j].ScenarioTotal
		}

		// 2. 当data[i].ScenarioTotal 和 data[j].ScenarioTotal 相等，判断data[i].TestCaseTotal 和 data[j].TestCaseTotal，最大的放最前边
		if data[i].TestCaseTotal != data[j].TestCaseTotal {
			return data[i].TestCaseTotal > data[j].TestCaseTotal
		}

		// 3. 当data[i].TestCaseTotal 和 data[j].TestCaseTotal相等，判断data[i].UpdatedAt 和 data[j].UpdatedAt
		// 注意：UpdatedAt是*time.Time，要判断空情况
		if data[i].UpdatedAt != nil && data[j].UpdatedAt != nil {
			return data[i].UpdatedAt.Before(*data[j].UpdatedAt)
		} else if data[i].UpdatedAt == nil {
			return false // 当data[i]的UpdatedAt为nil时，将data[j]排在前面
		} else {
			return true // 当data[j]的UpdatedAt为nil时，将data[i]排在前面
		}
	})

	for i, value := range data {
		value.Sort = int64(i + 1)
		data[i] = value
	}

	return data
}

func (s *SummaryProjectUserRankingService) SortRankingList(req v1.ResRankingList) (ret v1.ResRankingList, err error) {
	data := req.UserRankingList

	sort.Slice(data, func(i, j int) bool {
		// 1. 当data[i].ScenarioTotal 和 data[j].ScenarioTotal 不相等时候，最大的排前边
		if data[i].ScenarioTotal != data[j].ScenarioTotal {
			return data[i].ScenarioTotal > data[j].ScenarioTotal
		}

		// 2. 当data[i].ScenarioTotal 和 data[j].ScenarioTotal 相等，判断data[i].TestCaseTotal 和 data[j].TestCaseTotal，最大的放最前边
		if data[i].TestCaseTotal != data[j].TestCaseTotal {
			return data[i].TestCaseTotal > data[j].TestCaseTotal
		}

		// 3. 当data[i].TestCaseTotal 和 data[j].TestCaseTotal相等，判断data[i].UpdatedAt 和 data[j].UpdatedAt
		// 注意：UpdatedAt是string类型，需要解析为时间对象进行比较
		timeFormat := "2006-01-02 15:04:05"
		timeI, errI := time.Parse(timeFormat, data[i].UpdatedAt)
		timeJ, errJ := time.Parse(timeFormat, data[j].UpdatedAt)

		if errI == nil && errJ == nil {
			return timeI.Before(timeJ)
		} else if errI != nil {
			return false // 当data[i]的UpdatedAt解析失败时，将data[j]排在前面
		} else {
			return true // 当data[j]的UpdatedAt解析失败时，将data[i]排在前面
		}
	})

	for i, value := range data {
		value.Sort = int64(i + 1)
		data[i] = value
	}

	ret.UserRankingList = data

	return
}

func (s *SummaryProjectUserRankingService) GetRanking(tenantId consts.TenantId, projectId int64) (rankings []model.SummaryProjectUserRanking, err error) {

	users, err := s.FindUserIdsByProjectId(tenantId, projectId)

	cases, err := s.FindCasesTotalByProjectId(tenantId, projectId)
	scenarios, err := s.FindScenariosTotalByProjectId(tenantId, projectId)
	lastUpdateTime, _ := s.FindUserLastUpdateTestCasesByProjectId(tenantId, projectId)

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
		rankings = s.SortRanking(rankings)
	}

	return
}

func (s *SummaryProjectUserRankingService) SaveRanking(tenantId consts.TenantId) (err error) {
	projectIds, err := s.FindProjectIds(tenantId)
	for _, projectId := range projectIds {
		rankings, _ := s.GetRanking(tenantId, projectId)
		for _, ranking := range rankings {
			err := s.CreateByDate(tenantId, ranking)
			if err != nil {
				return err
			}
		}
	}
	return
}
