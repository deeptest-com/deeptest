package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
)

type SummaryDetailsService struct {
	SummaryDetailsRepo *repo.SummaryDetailsRepo `inject:""`
}

func NewSummaryDetailsService() *SummaryDetailsService {
	return &SummaryDetailsService{}
}

func (s *SummaryDetailsService) Card(projectId int64) (res v1.ResSummaryCard, err error) {
	var scenarioTotal int64
	var interfaceTotal int64
	var execTotal int64
	var passRate float64
	var coverage float64
	var oldCoverage float64
	var oldScenarioTotal int64
	var oldInterfaceTotal int64

	var summaryCardTotal model.SummaryCardTotal
	var oldSummaryCardTotal model.SummaryCardTotal
	var summaryDetails model.SummaryDetails
	var oldSummaryDetails model.SummaryDetails
	date := time.Now().AddDate(0, 0, -30)
	year, month, day := date.Date()
	startTime := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 00:00:00"
	endTime := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + " 23:59:59"
	if projectId == 0 {
		summaryCardTotal, err = s.SummaryCard()
		res.ProjectTotal, err = s.Count()
		oldSummaryCardTotal, err = s.SummaryCardByDate(startTime, endTime)

		coverage = summaryCardTotal.Coverage
		interfaceTotal = summaryCardTotal.InterfaceTotal
		execTotal = summaryCardTotal.ExecTotal
		passRate = summaryCardTotal.PassRate
		scenarioTotal = summaryCardTotal.ScenarioTotal
		oldCoverage = oldSummaryCardTotal.Coverage
		oldScenarioTotal = oldSummaryCardTotal.ScenarioTotal
		oldInterfaceTotal = oldSummaryCardTotal.InterfaceTotal

	} else {
		summaryDetails, err = s.FindByProjectId(projectId)
		res.ProjectTotal = 1
		oldSummaryDetails, err = s.FindByProjectIdAndDate(startTime, endTime, projectId)

		coverage = summaryDetails.Coverage
		interfaceTotal = summaryDetails.InterfaceTotal
		execTotal = summaryDetails.ExecTotal
		passRate = summaryDetails.PassRate
		scenarioTotal = summaryDetails.ScenarioTotal
		oldCoverage = oldSummaryDetails.Coverage
		oldScenarioTotal = oldSummaryDetails.ScenarioTotal
		oldInterfaceTotal = oldSummaryDetails.InterfaceTotal
	}

	res.Coverage = coverage
	res.InterfaceTotal = interfaceTotal
	res.ExecTotal = execTotal
	res.PassRate = passRate
	res.ScenarioTotal = scenarioTotal

	if oldCoverage != 0 {
		res.CoverageHb = DecimalHB(coverage, oldCoverage)
		res.CoverageHb, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", res.CoverageHb), 64)
	} else {
		res.CoverageHb = 0.0
	}

	if oldInterfaceTotal != 0 {
		res.InterfaceHb = DecimalHB(float64(interfaceTotal), float64(oldInterfaceTotal))
		res.InterfaceHb, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", res.InterfaceHb), 64)
	} else {
		res.CoverageHb = 0.0
	}

	if oldScenarioTotal != 0 {
		res.ScenarioHb = DecimalHB(float64(scenarioTotal), float64(oldScenarioTotal))
		res.ScenarioHb, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", res.ScenarioHb), 64)
	} else {
		res.CoverageHb = 0.0
	}

	return
}

func (s *SummaryDetailsService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	var resDetail v1.ResSummaryDetails
	var resDetails []v1.ResSummaryDetails
	var allDetails []model.SummaryDetails
	if userId == 0 {
		res.ProjectTotal, err = s.Count()
		allDetails, err = s.Find()

	} else {
		var projectIds []int64
		res.ProjectTotal, err = s.CountByUserId(userId)
		projectIds, err = s.FindProjectIdsByUserId(userId)
		allDetails, err = s.FindByProjectIds(projectIds)
	}

	for _, detail := range allDetails {
		copier.CopyWithOption(&resDetail, detail, copier.Option{DeepCopy: true})
		resDetail.Id = detail.ID
		resDetail.CreatedAt = time.Unix(detail.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")
		resDetail.Disabled = detail.Disabled
		resDetails = append(resDetails, resDetail)
	}
	res.ProjectList = resDetails
	return
}

func DecimalHB(newValue float64, oldValue float64) float64 {
	value := newValue / oldValue
	value = value - 1
	return value * 100
}

func (s *SummaryDetailsService) CreateByDate(req model.SummaryDetails) (err error) {

	//
	//var user v1.ResUserList
	//var userA v1.ResUserList
	//user.UserId = 1
	//user.UserName = "yanggggggg"
	//userA.UserId = 2
	//userA.UserName = "xiggggg"
	//resDetails.Coverage = 10.1
	//resDetails.InterfaceTotal = 5
	//resDetails.ScenarioTotal = 10
	//resDetails.PassRate = 11.5
	//resDetails.ExecTotal = 15
	//resDetails.AdminUser = "auto"
	//resDetails.ProjectCreateTime = "2023-03-17 09:15:15"
	//resDetails.ProjectName = "projectAuto"
	//resDetails.ProjectChineseName = "自动化创建"
	//resDetails.ProjectDes = "miaoshu"
	//resDetails.ProjectId = 10

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

func (s *SummaryDetailsService) SummaryCard() (summaryCardTotal model.SummaryCardTotal, err error) {
	return s.SummaryDetailsRepo.SummaryCard()
}

func (s *SummaryDetailsService) SummaryCardByDate(startTime string, endTime string) (summaryCardTotal model.SummaryCardTotal, err error) {
	return s.SummaryDetailsRepo.SummaryCardByDate(startTime, endTime)
}

func (s *SummaryDetailsService) FindByProjectIdAndDate(startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.FindByProjectIdAndDate(startTime, endTime, projectId)
}
