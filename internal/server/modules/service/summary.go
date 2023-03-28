package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/core/cache"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"math/rand"
	"strconv"
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

func (s *SummaryService) Bugs(projectId int64) (res v1.ResSummaryBugs, err error) {
	bugsCache, err := cache.GetCacheBytes("summaryBugs-" + strconv.FormatInt(projectId, 36))
	kong, err := cache.GetCacheString("kong")
	fmt.Println(bugsCache)
	fmt.Println(kong)

	if err != nil || len(bugsCache) == 0 {
		res, err = s.SummaryBugsService.Bugs(projectId)
		cache.SetCache("summaryBugs-"+strconv.FormatInt(projectId, 36), res, time.Duration(rand.Int63n(50)+50))
	} else {
		json.Unmarshal(bugsCache, &res)
	}

	return
}

func (s *SummaryService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	detailsCache, err := cache.GetCacheBytes("summaryDetails-" + strconv.FormatInt(userId, 36))
	kong, err := cache.GetCacheString("kong")
	fmt.Println(detailsCache)
	fmt.Println(kong)

	if err != nil || len(detailsCache) == 0 {
		res, err = s.SummaryDetailsService.Details(userId)
		cache.SetCache("summaryDetails-"+strconv.FormatInt(userId, 36), res, time.Duration(rand.Int63n(50)+50))
	} else {
		json.Unmarshal(detailsCache, &res)
	}

	return
}

func (s *SummaryService) ProjectUserRanking(projectId int64, cycle int64) (res v1.ResRankingList, err error) {
	rankingCache, err := cache.GetCacheBytes("summaryRanking-" + strconv.FormatInt(projectId, 36))
	kong, err := cache.GetCacheString("kong")
	fmt.Println(rankingCache)
	fmt.Println(kong)

	if err != nil || len(rankingCache) == 0 {
		res, err = s.SummaryProjectUserRankingService.ProjectUserRanking(projectId, cycle)
		cache.SetCache("summaryRanking-"+strconv.FormatInt(projectId, 36), res, time.Duration(rand.Int63n(50)+50))
	} else {
		json.Unmarshal(rankingCache, &res)
	}

	return
}

func (s *SummaryService) Card(projectId int64) (res v1.ResSummaryCard, err error) {
	cardCache, err := cache.GetCacheBytes("summaryCard-" + strconv.FormatInt(projectId, 36))
	kong, err := cache.GetCacheString("kong")
	fmt.Println(cardCache)
	fmt.Println(kong)

	if err != nil || len(cardCache) == 0 {
		res, err = s.SummaryDetailsService.Card(projectId)
		cache.SetCache("summaryCard-"+strconv.FormatInt(projectId, 36), res, time.Duration(rand.Int63n(50)+50))
	} else {
		json.Unmarshal(cardCache, &res)
		cache.DeleteCache("summaryCard-" + strconv.FormatInt(projectId, 36))
		cache.SetCache("summaryCard-"+strconv.FormatInt(projectId, 36), res, time.Duration(rand.Int63n(50)+50))
	}

	return
}

func (s *SummaryService) Collection(scope string) (err error) {

	switch scope {
	case "all":
		err = s.CollectionDetails()
		return
	case "details":
		err = s.CollectionDetails()
		return
	case "card":
		return
	case "ranking":
		err = s.CollectionRanking()
		return
	case "bugs":
		err = s.CollectionBugs()
		return
	}

	//set cache
	//sql data
	//create struct

	return
}

func (s *SummaryService) CollectionRanking() (err error) {

	//res, err := s.SummaryDetailsService.Card(0)
	//cache.SetCache("summaryDetails-"+strconv.FormatInt(user.UserId, 36), resDetail, time.Duration(rand.Int63n(50)+50))
	return
}

func (s *SummaryService) CollectionBugs() (err error) {
	//配置地址

	//请求对应系统,获取bug信息

	//bug转化,配置字段映射关系
	bugs := model.SummaryBugs{}

	//调用存储

	s.SummaryBugsService.Create(bugs)
	cache.SetCache("summaryDetails-"+strconv.FormatInt(bugs.ProjectId, 36), bugs, time.Duration(rand.Int63n(50)+50))
	return
}

func (s *SummaryService) CollectionDetails() (err error) {
	//SummaryDetailsService := NewSummaryDetailsService()

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
		passRate, _ := s.SummaryDetailsService.FindPassRate(detail.ProjectId)
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

			cache.DeleteCache("summaryDetails-" + strconv.FormatInt(user.UserId, 36))
			cache.SetCache("summaryDetails-"+strconv.FormatInt(user.UserId, 36), resDetail, time.Duration(rand.Int63n(50)+50))
		}
	}
	return
}

// SummaryDataCheck corn task
func (s *SummaryService) SummaryDataCheck() (err error) {

	//SummaryBugsService := NewSummaryBugsService()
	//SummaryDetailsService := NewSummaryDetailsService()
	//SummaryProjectUserRankingService := NewSummaryProjectUserRankingService()

	var checks []v1.SummaryDataCheck

	values, err := cache.GetCacheBytes("summaryDataUpdatedAt")

	if err != nil || len(values) == 0 {
		s.Collection("all")
	} else {
		json.Unmarshal(values, &checks)
		for _, value := range checks {
			switch value.CacheKey {
			//由于所有数据都基于details表抓取,所以details表需要特殊处理
			case "details":
				ret, err := s.SummaryDetailsService.CheckDetailsUpdated(value.CacheValue)
				if err == nil && ret != false {
					s.Collection("details")
				}
			case "card":
				ret, err := s.SummaryDetailsService.CheckCardUpdated(value.CacheValue)
				if err == nil && ret != false {
					s.Collection("card")
				}
			case "ranking":
				ret, err := s.SummaryProjectUserRankingService.CheckUpdated(value.CacheValue)
				if err == nil && ret != false {
					s.Collection("ranking")
				}
			case "bugs":
				ret, err := s.SummaryBugsService.CheckUpdated(value.CacheValue)
				if err == nil && ret != false {
					s.Collection("bugs")
				}
			}

		}
	}

	return
}
