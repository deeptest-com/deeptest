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
	return new(SummaryDetailsService)
}

func (s *SummaryDetailsService) Card(projectId int64) (res v1.ResSummaryCard, err error) {
	var scenarioTotal, interfaceTotal, execTotal, oldScenarioTotal, oldInterfaceTotal int64
	var passRate, coverage, oldCoverage float64
	var summaryCardTotal, oldSummaryCardTotal model.SummaryCardTotal
	var summaryDetails, oldSummaryDetails model.SummaryDetails

	date := time.Now().AddDate(0, 0, -30)
	startTime, endTime := GetDate(date)
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
		res.CoverageHb, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", DecimalHB(coverage, oldCoverage)), 64)
	}

	if oldInterfaceTotal != 0 {
		res.InterfaceHb, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", DecimalHB(float64(interfaceTotal), float64(oldInterfaceTotal))), 64)
	}

	if oldScenarioTotal != 0 {
		res.ScenarioHb, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", DecimalHB(float64(scenarioTotal), float64(oldScenarioTotal))), 64)
	}

	return
}

func (s *SummaryDetailsService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	//从project表收集项目总数
	res.ProjectTotal, err = s.Count()
	res.UserProjectTotal, err = s.CountByUserId(userId)
	//查找用户参与的项目id
	projectIds, err := s.FindProjectIdsByUserId(userId)
	//查询details所有的项目信息
	allDetails, err := s.Find()
	//查询所有项目信息
	projectDetails, err := s.FindProjectDetails()
	//组装返回的json结构体
	res.ProjectList, res.UserProjectList = s.HandleSummaryDetails(projectIds, allDetails, projectDetails)
	return
}

func (s *SummaryDetailsService) HandleSummaryDetails(projectIds []int64, allDetails []model.SummaryDetails, projectDetails []model.Project) (resAllDetails []v1.ResSummaryDetails, resUserDetails []v1.ResSummaryDetails) {

	//遍历项目信息，匹配details表结果，进行字段复制，组装返回resAllDetails体
	for _, projectDetail := range projectDetails {

		for _, detail := range allDetails {
			var resDetail v1.ResSummaryDetails

			if int64(projectDetail.ID) == detail.ProjectId {
				//如果detail中有当前projectid对应的信息，则把detail数据赋值给结果resDetail
				resDetail = s.CopyProjectInfo(projectDetail, detail)
				resDetail.BugTotal, _ = s.CountBugsByProjectId(detail.ProjectId)
				break
			} else {
				//如果detail中没有当前projectid对应的信息，则把复制个空的detail数据给结果resDetail
				var nilDetail model.SummaryDetails
				resDetail = s.CopyProjectInfo(projectDetail, nilDetail)
			}

			userList, _ := s.FindUserIdAndNameByProjectId(detail.ProjectId)
			resDetail.UserList = userList
			resAllDetails = append(resAllDetails, resDetail)
			//当前项目如果是用户参与的项目，则添加到resUserDetails中
			for _, id := range projectIds {
				if detail.ProjectId == id {
					resUserDetails = append(resUserDetails, resDetail)
					break
				}
			}
		}
	}
	return
}

func (s *SummaryDetailsService) CopyProjectInfo(projectDetail model.Project, detail model.SummaryDetails) (resDetail v1.ResSummaryDetails) {

	copier.CopyWithOption(&resDetail, projectDetail, copier.Option{DeepCopy: true})
	copier.CopyWithOption(&resDetail, detail, copier.Option{DeepCopy: true})
	resDetail.Id = detail.ID
	resDetail.ProjectDescr = projectDetail.Desc
	resDetail.ProjectName = projectDetail.Name
	resDetail.ProjectShortName = projectDetail.ShortName
	resDetail.ProjectId = int64(projectDetail.ID)
	resDetail.CreatedAt = projectDetail.CreatedAt.Format("2022-01-11 11:22:12")
	resDetail.AdminName, _ = s.FindAdminNameByAdminId(int64(projectDetail.AdminId))
	return
}

func (s *SummaryDetailsService) CopyDetailsWithoutBaseModel(detail model.SummaryDetails) (ret model.SummaryDetails) {
	var newDetail model.SummaryDetails
	newDetail.ScenarioTotal = detail.ScenarioTotal
	newDetail.ProjectId = detail.ProjectId
	newDetail.Coverage = detail.Coverage
	newDetail.ExecTotal = detail.ExecTotal
	newDetail.PassRate = detail.PassRate
	newDetail.InterfaceTotal = detail.InterfaceTotal
	ret = newDetail
	return
}

func (s *SummaryDetailsService) Create(req model.SummaryDetails) (err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Create(req)
}

func (s *SummaryDetailsService) CreateByDate(req model.SummaryDetails) (err error) {
	now := time.Now()
	startTime, endTime := GetDate(now)
	ret, err := s.HasDataOfDate(startTime, endTime)
	if ret {
		err = s.Create(req)
	} else {
		err = s.UpdateColumnsByDate(req, startTime, endTime)
	}
	return
}

func (s *SummaryDetailsService) UpdateColumnsByDate(req model.SummaryDetails, startTime string, endTime string) (err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.UpdateColumnsByDate(req, startTime, endTime)
}

func (s *SummaryDetailsService) Count() (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Count()
}

func (s *SummaryDetailsService) CountByUserId(userId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountByUserId(userId)
}

func (s *SummaryDetailsService) FindProjectDetails() (projectDetails []model.Project, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindProjectDetails()

}

func (s *SummaryDetailsService) FindAdminNameByAdminId(adminId int64) (adminName string, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindAdminNameByAdminId(adminId)
}

func (s *SummaryDetailsService) FindProjectIdsByUserId(userId int64) (count []int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindProjectIdsByUserId(userId)
}

func (s *SummaryDetailsService) FindInterfaceIdsByProjectId(projectId int64) (ids []int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindInterfaceIdsByProjectId(projectId)
}

func (s *SummaryDetailsService) CoverageByProjectId(projectId int64, interfaceIds []int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CoverageByProjectId(projectId, interfaceIds)
}

func (s *SummaryDetailsService) FindByProjectId(projectId int64) (summaryDetail model.SummaryDetails, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindByProjectId(projectId)
}

func (s *SummaryDetailsService) Find() (details []model.SummaryDetails, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Find()
}

func (s *SummaryDetailsService) FindByProjectIds(projectIds []int64) (details []model.SummaryDetails, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindByProjectIds(projectIds)
}

func (s *SummaryDetailsService) SummaryCard() (summaryCardTotal model.SummaryCardTotal, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.SummaryCard()
}

func (s *SummaryDetailsService) SummaryCardByDate(startTime string, endTime string) (summaryCardTotal model.SummaryCardTotal, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.SummaryCardByDate(startTime, endTime)
}

func (s *SummaryDetailsService) FindByProjectIdAndDate(startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindByProjectIdAndDate(startTime, endTime, projectId)
}

func (s *SummaryDetailsService) FindUserIdAndNameByProjectId(projectId int64) (userIdAndName []v1.ResUserIdAndName, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindUserIdAndNameByProjectId(projectId)
}

func (s *SummaryDetailsService) FindCreateUserNameByProjectId(projectId int64) (userName string, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindCreateUserNameByProjectId(projectId)
}

func (s *SummaryDetailsService) CountBugsByProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountBugsByProjectId(projectId)
}

func (s *SummaryDetailsService) CountScenarioTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountScenarioTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountInterfaceTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountInterfaceTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountExecTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountExecTotalProjectId(projectId)
}

func (s *SummaryDetailsService) FindPassRateByProjectId(projectId int64) (passRate float64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindPassRateByProjectId(projectId)
}

func (s *SummaryDetailsService) HasDataOfDate(startTime string, endTiem string) (ret bool, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.HasDataOfDate(startTime, endTiem)
}

func (s *SummaryDetailsService) CheckCardUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CheckCardUpdated(lastUpdateTime)
}

//检查是否有今日数据,没有则copy最后一条,然后进行数据是否更新检查
func (s *SummaryDetailsService) CheckDetailsUpdated(lastUpdateTime *time.Time) (result bool, err error) {
	r := repo.NewSummaryDetailsRepo()
	now := time.Now()
	startTime, endTime := GetDate(now)
	ret, err := s.HasDataOfDate(startTime, endTime)
	if !ret {
		details, _ := s.Find()
		for _, detail := range details {
			newDetail := s.CopyDetailsWithoutBaseModel(detail)
			s.Create(newDetail)
		}
	}
	return r.CheckDetailsUpdated(lastUpdateTime)
}

func (s *SummaryDetailsService) CollectionProjectInfo() (details []model.SummaryDetails, err error) {
	return s.SummaryDetailsRepo.CollectionProjectInfo()
}
