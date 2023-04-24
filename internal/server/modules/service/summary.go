package service

import (
	"context"
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	//"github.com/aaronchen2k/deeptest/internal/server/core/CacheOption"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type SummaryService struct {
	SummaryDetailsService            *SummaryDetailsService            `inject:""`
	SummaryBugsService               *SummaryBugsService               `inject:""`
	SummaryProjectUserRankingService *SummaryProjectUserRankingService `inject:""`
}

func NewSummaryService() *SummaryService {
	return &SummaryService{}
}

func (s *SummaryService) CacheOption() redis.UniversalClient {
	universalOptions := &redis.UniversalOptions{
		Addrs:       strings.Split(config.CONFIG.Redis.Addr, ","),
		Password:    config.CONFIG.Redis.Password,
		PoolSize:    config.CONFIG.Redis.PoolSize,
		IdleTimeout: 300 * time.Second,
	}
	CACHE := redis.NewUniversalClient(universalOptions)

	return CACHE
}

func (s *SummaryService) ToString(arg interface{}) (str string) {
	switch arg.(type) {
	case int64:
		str = strconv.FormatInt(arg.(int64), 10)
	case string:
		str = arg.(string)
	}
	return
}

func (s *SummaryService) GetCache(key string, arg interface{}) (result []byte, err error) {
	str := s.ToString(arg)
	result, err = s.CacheOption().Get(context.Background(), key+str).Bytes()
	return
}

func (s *SummaryService) SetCache(key string, arg interface{}, object interface{}) {
	value, _ := json.Marshal(object)
	str := s.ToString(arg)
	s.CacheOption().Set(context.Background(), key+str, value, time.Duration(rand.Intn(30)+30)*time.Minute)
}

func (s *SummaryService) DelCache(key string, arg interface{}) {
	str := s.ToString(arg)
	s.CacheOption().Del(context.Background(), key+str)
}

func (s *SummaryService) Bugs(projectId int64) (res v1.ResSummaryBugs, err error) {
	bugsCache, err := s.GetCache("summaryBugs", projectId)

	if err != nil || len(bugsCache) == 0 {
		res, err = s.SummaryBugsService.Bugs(projectId)
		s.SetCache("summaryBugs", projectId, res)
	} else if err == nil && len(bugsCache) != 0 {
		err = json.Unmarshal(bugsCache, &res)
		if err != nil {
			s.DelCache("summaryBugs", projectId)
		}
	}

	return
}

func (s *SummaryService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	detailsCache, err := s.GetCache("summaryDetails", userId)

	if err != nil || len(detailsCache) == 0 {
		res, err = s.SummaryDetailsService.Details(userId)
		s.SetCache("summaryDetails", userId, res)
	} else if err == nil && len(detailsCache) != 0 {
		err = json.Unmarshal(detailsCache, &res)
		if err != nil {
			s.DelCache("summaryDetails", userId)
		}
	}
	return
}

func (s *SummaryService) ProjectUserRanking(projectId int64, cycle int64) (res v1.ResRankingList, err error) {
	rankingCache, err := s.GetCache("summary"+strconv.FormatInt(cycle, 10)+"Ranking"+strconv.FormatInt(projectId, 10), projectId)
	if err != nil || len(rankingCache) == 0 {
		res, err = s.SummaryProjectUserRankingService.ProjectUserRanking(projectId, cycle)
		s.SetCache("summary"+strconv.FormatInt(cycle, 10)+"Ranking"+strconv.FormatInt(projectId, 10), projectId, res)
	} else if err == nil && len(rankingCache) != 0 {
		err = json.Unmarshal(rankingCache, &res)
		if err != nil {
			s.DelCache("summary"+strconv.FormatInt(cycle, 10)+"Ranking"+strconv.FormatInt(projectId, 10), projectId)
		}
	}

	return
}

func (s *SummaryService) Card(projectId int64) (res v1.ResSummaryCard, err error) {
	cardCache, err := s.GetCache("summaryCard", projectId)
	if err != nil || len(cardCache) == 0 {
		res, err = s.SummaryDetailsService.Card(projectId)
		s.SetCache("summaryCard", projectId, res)
	} else if err == nil && len(cardCache) != 0 {
		err = json.Unmarshal(cardCache, &res)
		if err != nil {
			s.DelCache("summaryCard", projectId)
		}
	}
	return
}

func (s *SummaryService) Collection(scope string) (err error) {

	switch scope {
	case "all":
		err = s.CollectionDetails()
		err = s.CollectionBugs()
		err = s.CollectionRanking()
		return
	case "details":
		err = s.CollectionDetails()
		return
	case "ranking":
		err = s.CollectionRanking()
		return
	case "bugs":
		err = s.CollectionBugs()
		return
	}

	return
}

func (s *SummaryService) CollectionRanking() (err error) {
	projectIds, err := s.SummaryProjectUserRankingService.FindProjectIds()
	for i := 0; i < 2; i++ {
		for projectId := range projectIds {
			res, err := s.SummaryProjectUserRankingService.ProjectUserRanking(int64(projectId), int64(i))
			if err == nil {
				s.SetCache("summary"+strconv.Itoa(i)+"Ranking"+strconv.Itoa(projectId), "", res)
			}
		}
	}
	checkTime := time.Now().Local()
	s.SetCache("summaryDataUpdatedAt", "ranking", &checkTime)
	return
}

func (s *SummaryService) CollectionBugs() (err error) {
	//配置地址

	//请求对应系统,获取bug信息

	//bug转化,配置字段映射关系
	bugs := model.SummaryBugs{}

	//调用存储
	//s.SummaryBugsService.Create(bugs)
	ids, err := s.SummaryBugsService.FindProjectIds()

	for id := range ids {
		value, _ := s.SummaryBugsService.Bugs(int64(id))
		s.SetCache("summaryBugs", bugs.ProjectId, value)
	}
	checkTime := time.Now().Local()
	s.SetCache("summaryDataUpdatedAt", "bugs", &checkTime)
	return
}

func (s *SummaryService) CollectionDetails() (err error) {

	var details []model.SummaryDetails

	//从project表获取所有项目id、name、描述、简称、创建时间
	details, err = s.SummaryDetailsService.CollectionProjectInfo()

	for index, detail := range details {

		//根据projectid获取所有用户列表,将id最小的名字赋值进来,现成的方法返回getusersByprojectid

		detail.AdminUser, _ = s.SummaryDetailsService.FindCreateUserNameByProjectId(detail.ProjectId)

		//从biz_scenario表根据projectid,查找场景总数
		detail.ScenarioTotal, _ = s.SummaryDetailsService.CountScenarioTotalProjectId(detail.ProjectId)

		//从biz_interface表根据projectid,查找接口总数
		detail.InterfaceTotal, _ = s.SummaryDetailsService.CountInterfaceTotalProjectId(detail.ProjectId)

		//根据projectid,从biz_scenario_report表,获得所有报告总数,然后计算
		detail.ExecTotal, _ = s.SummaryDetailsService.CountExecTotalProjectId(detail.ProjectId)

		//从biz_scenario_report拿到assertion的相关数据,计算后存储
		passRate, _ := s.SummaryDetailsService.FindPassRateByProjectId(detail.ProjectId)
		detail.PassRate, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", passRate), 64)

		//从biz_interface需要获取当前项目的所有接口,然后从biz_processor_interface检查哪些在场景中出现过
		interfaceIds, _ := s.SummaryDetailsService.FindInterfaceIdsByProjectId(detail.ProjectId)
		count, _ := s.SummaryDetailsService.CoverageByProjectId(detail.ProjectId, interfaceIds)
		var coverage float64
		if detail.InterfaceTotal != 0 {
			coverage = float64(count / detail.InterfaceTotal)
		} else {
			coverage = 0
		}
		detail.Coverage, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", coverage), 64)

		s.SummaryDetailsService.CreateByDate(detail)

		details[index] = detail

		userList, _ := s.SummaryDetailsService.FindUserIdAndNameByProjectId(detail.ProjectId)
		for _, user := range userList {
			var resDetail v1.ResSummaryDetails
			copier.CopyWithOption(&resDetail, detail, copier.Option{DeepCopy: true})
			resDetail.Id = detail.ID
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", detail.ProjectCreateTime, time.Local)
			resDetail.CreatedAt = t.Format("2006-01-02 15:04:05")
			resDetail.Disabled = detail.Disabled
			resDetail.BugTotal, _ = s.SummaryDetailsService.CountBugsByProjectId(detail.ProjectId)
			resDetail.UserList = userList
			value, _ := json.Marshal(resDetail)
			s.SetCache("summaryDetails", user.UserId, value)
		}
	}
	checkTime := time.Now().Local()
	s.SetCache("summaryDataUpdatedAt", "details", &checkTime)
	return
}

// SummaryDataCheck corn task
func (s *SummaryService) SummaryDataCheck() (err error) {
	modelList := []string{"details", "ranking", "bugs"}

	for _, m := range modelList {
		values, err := s.GetCache("summaryDataUpdatedAt", m)
		if err != nil || len(values) == 0 {
			s.Collection("all")
		} else {
			var t *time.Time
			switch m {
			case "details":
				err = json.Unmarshal(values, &t)
				ret, err := s.SummaryDetailsService.CheckDetailsUpdated(t)
				if err == nil && ret != false {
					s.Collection("details")
				}
			case "ranking":
				err = json.Unmarshal(values, &t)
				ret, err := s.SummaryProjectUserRankingService.CheckUpdated(t)
				if err == nil && ret != false {
					s.Collection("ranking")
				}
			case "bugs":
				err = json.Unmarshal(values, &t)
				ret, err := s.SummaryBugsService.CheckUpdated(t)
				if err == nil && ret != false {
					s.Collection("bugs")
				}
			}

		}

	}
	return
}
