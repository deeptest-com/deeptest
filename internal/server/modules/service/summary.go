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

func (s *SummaryService) Bugs(projectId int64) (res v1.ResSummaryBugs, err error) {
	res, err = s.SummaryBugsService.Bugs(projectId)
	return
}

func (s *SummaryService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	//改为项目数据实时，但统计数据非实时
	res, err = s.SummaryDetailsService.Details(userId)
	return
}

func (s *SummaryService) ProjectUserRanking(cycle int64, projectId int64) (res v1.ResRankingList, err error) {
	res, err = s.SummaryProjectUserRankingService.ProjectUserRanking(cycle, projectId)
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
	s.SummaryProjectUserRankingService.SaveRanking()
	return
}

func (s *SummaryService) CollectionBugs() (err error) {
	//配置地址
	//请求对应系统,获取bug信息
	//bug转化,配置字段映射关系
	//调用存储
	//s.SummaryBugsService.CreateBug(bugs)

	return
}

func (s *SummaryService) CollectionDetails() (err error) {
	err = s.SummaryDetailsService.SaveDetails()
	return
}

func DecimalHB(newValue float64, oldValue float64) float64 {
	var value float64
	if newValue != 0 {
		value = newValue / oldValue
		value = value - 1
	} else {
		value = 1 - oldValue
	}
	return value * 100
}

func GetTodayStartAndEndTime() (startTime string, endTime string) {
	today := time.Now().AddDate(0, 0, 0)
	year, month, day := today.Date()
	startTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 00:00:00"
	endTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"
	return
}

func GetEarlierDateStartAndEndTime(earlier int64) (startTime string, endTime string) {
	earlierDate := time.Now().AddDate(0, 0, int(earlier))
	year, month, day := earlierDate.Date()
	startTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 00:00:00"
	endTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"
	return
}

func GetEarlierDateUntilTodayStartAndEndTime(earlier int64) (startTime string, endTime string) {
	earlierDate := time.Now().AddDate(0, 0, int(earlier))
	today := time.Now().AddDate(0, 0, 0)

	earlierYear, earlierMonth, earlierDay := earlierDate.Date()
	year, month, day := today.Date()

	startTime = strconv.Itoa(earlierYear) + "-" + strconv.Itoa(int(earlierMonth)) + "-" + strconv.Itoa(earlierDay) + " 00:00:00"
	endTime = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"
	return
}
