package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
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
	res, err = s.SummaryBugsService.Bugs(projectId)
	return
}

func (s *SummaryService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	//改为项目数据实时，但统计数据非实时
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
	err = s.CollectionDetails()
	err = s.CollectionBugs()
	err = s.CollectionRanking()
	return
}

func (s *SummaryService) CollectionRanking() (err error) {
	//projectIds, err := s.SummaryProjectUserRankingService.FindProjectIds()
	//
	////从各地方获取ranking数据然后存储
	//sort bigint
	//project_id text
	//user_id bigint
	//user_name text
	//scenario_total text
	//testcases_total text
	//
	//s.SummaryProjectUserRankingService.CreateByDate();

	return
}

func (s *SummaryService) CollectionBugs() (err error) {
	//配置地址
	//请求对应系统,获取bug信息
	//bug转化,配置字段映射关系
	//调用存储
	//s.SummaryBugsService.Create(bugs)

	return
}

func (s *SummaryService) CollectionDetails() (err error) {
	//从project表获取所有项目id
	ids, err := s.SummaryDetailsService.FindProjectIds()
	for _, id := range ids {
		detail := s.SummaryDetailsService.CollectDetailByProjectId(id)
		s.SummaryDetailsService.CreateByDate(detail)
	}
	return
}

func DecimalHB(newValue float64, oldValue float64) float64 {
	value := newValue / oldValue
	value = value - 1
	return value * 100
}

func GetDate(date time.Time) (startTime string, endTime string) {
	year, month, day := date.Date()
	startTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 00:00:00"
	endTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"
	return
}
