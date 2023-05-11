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
	//查找所有项目id
	projectIds, err := s.FindProjectIds()
	//查找所有项目对应的summaryDetail数据，并转为map
	allDetails := s.GetAllDetailOfMap(projectIds)
	//查找用户参与的项目id,并转为map
	userProjectIds, err := s.FindProjectIdsByUserId(userId)

	//查询所有项目信息
	allProjectsInfo, err := s.FindAllProjectInfo()
	//组装返回的json结构体
	res.ProjectList, res.UserProjectList = s.HandleSummaryDetails(projectIds, userProjectIds, allDetails, allProjectsInfo)
	return
}

func (s *SummaryDetailsService) HandleSummaryDetails(projectIds []int64, userProjectIds []int64, allDetails map[int64]model.SummaryDetails, allProjectsInfo []model.SummaryProjectInfo) (resAllDetails []v1.ResSummaryDetails, resUserDetails []v1.ResSummaryDetails) {
	var id, userId uint
	var projectsUserListOfMap map[int64][]v1.ResUserIdAndName
	projectsBugCount, _ := s.CountBugsGroupByProjectId()
	projectsUsers, _ := s.FindAllUserIdAndNameOfProject()
	//如果获取的
	if projectsUsers != nil {
		projectsUserListOfMap = s.LetUsersGroupByProjectId(projectIds, projectsUsers)
	} else {
		projectsUserListOfMap = nil
	}

	//遍历项目信息，匹配details表结果，进行字段复制，组装返回resAllDetails体
	for _, projectInfo := range allProjectsInfo {

		var resDetail v1.ResSummaryDetails
		hit := false
		for _, detail := range allDetails {
			if int64(projectInfo.ID) == detail.ProjectId {
				//如果detail中有当前projectid对应的信息，则把detail数据赋值给结果resDetail
				resDetail = s.CopyProjectInfo(projectInfo, allDetails[int64(projectInfo.ID)])
				hit = true
				break
			}
		}

		if !hit {
			//如果detail中没有当前projectid对应的信息，则把复制个空的detail数据给结果resDetail
			var nilDetail model.SummaryDetails
			resDetail = s.CopyProjectInfo(projectInfo, nilDetail)
		}
		id = id + 1
		resDetail.Id = id

		if projectsBugCount != nil {
			for _, bugCount := range projectsBugCount {
				if bugCount.ProjectId == int64(projectInfo.ID) {
					resDetail.BugTotal = bugCount.Count
				}
			}
		}

		if projectsUserListOfMap != nil {
			resDetail.UserList = projectsUserListOfMap[int64(projectInfo.ID)]
		}

		resAllDetails = append(resAllDetails, resDetail)
		//当前项目如果是用户参与的项目，则添加到resUserDetails中
		for _, id := range userProjectIds {
			if int64(projectInfo.ID) == id {
				userId = userId + 1
				resDetail.Id = userId
				resUserDetails = append(resUserDetails, resDetail)
				break
			}
		}
	}
	return
}

func (s *SummaryDetailsService) CopyProjectInfo(projectInfo model.SummaryProjectInfo, detail model.SummaryDetails) (resDetail v1.ResSummaryDetails) {

	copier.CopyWithOption(&resDetail, projectInfo, copier.Option{DeepCopy: true})
	copier.CopyWithOption(&resDetail, detail, copier.Option{DeepCopy: true})
	resDetail.ProjectDescr = projectInfo.Descr
	resDetail.ProjectName = projectInfo.Name
	resDetail.ProjectShortName = projectInfo.ShortName
	resDetail.ProjectId = int64(projectInfo.ID)
	resDetail.CreatedAt = projectInfo.CreatedAt.Format("2006-01-02 15:04:05")
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

func (s *SummaryDetailsService) GetAllDetailOfMap(projectIds []int64) (ret map[int64]model.SummaryDetails) {

	//从biz_scenario表根据projectid,查找场景总数
	ScenariosTotal, _ := s.CountAllScenarioTotalProjectId()

	//根据projectid,从biz_scenario_report表,获得所有报告总数,然后计算
	execsTotal, _ := s.CountAllExecTotalProjectId()

	//从biz_interface表根据projectid,查找接口总数
	interfacesTotal, _ := s.CountAllEndpointTotalProjectId()

	//从biz_scenario_report拿到assertion的相关数据,计算后存储
	passRates, _ := s.FindAllPassRateByProjectId()

	//从biz_endpoint需要获取当前项目的所有接口,然后从biz_processor_interface检查哪些在场景中出现过
	interfaceIds, _ := s.FindAllEndpointIdsGroupByProjectId(projectIds)
	coverages, _ := s.CoverageAllByProjectId(interfaceIds)

	ret = make(map[int64]model.SummaryDetails, len(projectIds))

	for _, projectId := range projectIds {
		details := s.CollectDetail(projectId, ScenariosTotal[projectId], interfacesTotal[projectId], execsTotal[projectId], passRates[projectId], coverages[projectId])
		//返回的数组，需要处理成map形式
		ret[projectId] = details
	}

	return
}

func (s *SummaryDetailsService) LetUsersGroupByProjectId(projectIds []int64, projectsUsers []model.UserIdAndName) (ret map[int64][]v1.ResUserIdAndName) {
	//将拿到的userList，根据projectid，装入map中，key是projectid，value是userList    []v1.ResUserIdAndName
	var m1 map[int64][]v1.ResUserIdAndName
	m1 = make(map[int64][]v1.ResUserIdAndName, len(projectIds))
	for _, projectId := range projectIds {
		var tmpUsers []v1.ResUserIdAndName
		for _, projectUsers := range projectsUsers {
			if projectUsers.ProjectId == projectId {
				var tmpUser v1.ResUserIdAndName
				copier.CopyWithOption(&tmpUser, projectUsers, copier.Option{DeepCopy: true})
				tmpUsers = append(tmpUsers, tmpUser)
			}
		}
		m1[projectId] = tmpUsers
	}
	return m1
}

func (s *SummaryDetailsService) Create(req model.SummaryDetails) (err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Create(req)
}

func (s *SummaryDetailsService) CreateByDate(req model.SummaryDetails) (err error) {
	now := time.Now()
	startTime, endTime := GetDate(now)
	ret, err := s.HasDataOfDate(startTime, endTime, req.ProjectId)
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

func (s *SummaryDetailsService) FindAllProjectInfo() (projectDetails []model.SummaryProjectInfo, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindAllProjectInfo()

}

func (s *SummaryDetailsService) FindAllAdminNameByAdminId(adminId int64) (adminName string, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindAdminNameByAdminId(adminId)
}

func (s *SummaryDetailsService) FindProjectIdsByUserId(userId int64) (count []int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindProjectIdsByUserId(userId)
}

func (s *SummaryDetailsService) FindEndpointIdsByProjectId(projectId int64) (ids []int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindEndpointIdsByProjectId(projectId)
}

func (s *SummaryDetailsService) FindAllEndpointIdsGroupByProjectId(projectIds []int64) (ids map[int64][]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	results, err := r.FindAllEndpointIdsGroupByProjectId()
	ids = make(map[int64][]int64, len(projectIds))
	for _, projectId := range projectIds {
		var tmpUsers []int64
		for _, result := range results {
			if result.ProjectId == projectId {
				tmpUsers = append(tmpUsers, result.Id)
			}
		}
		ids[projectId] = tmpUsers
	}
	return
}

func (s *SummaryDetailsService) CoverageByProjectId(projectId int64, interfaceIds []int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CoverageByProjectId(projectId, interfaceIds)
}

func (s *SummaryDetailsService) CoverageAllByProjectId(interfaceIds map[int64][]int64) (count map[int64]int64, err error) {
	count = make(map[int64]int64, len(interfaceIds))
	r := repo.NewSummaryDetailsRepo()

	for projectId, ids := range interfaceIds {
		result, _ := r.CoverageByProjectId(projectId, ids)
		count[projectId] = result
	}
	return
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

func (s *SummaryDetailsService) FindProjectIds() (ids []int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindProjectIds()
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

func (s *SummaryDetailsService) FindAllUserIdAndNameOfProject() (users []model.UserIdAndName, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindAllUserIdAndNameOfProject()
}

func (s *SummaryDetailsService) FindCreateUserNameByProjectId(projectId int64) (userName string, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindCreateUserNameByProjectId(projectId)
}

func (s *SummaryDetailsService) CountBugsGroupByProjectId() (bugsCount []model.ProjectsBugCount, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountBugsGroupByProjectId()
}

func (s *SummaryDetailsService) CountScenarioTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountScenarioTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountAllScenarioTotalProjectId() (scenarioTotal map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	scenariosTotal, _ := r.CountAllScenarioTotalProjectId()

	scenarioTotal = make(map[int64]int64, len(scenariosTotal))
	for key, value := range scenariosTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		scenarioTotal[int64(key)] = value.Id
	}

	return
}

func (s *SummaryDetailsService) CountEndpointTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountEndpointTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountAllEndpointTotalProjectId() (counts map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	endpointsTotal, _ := r.CountAllEndpointTotalProjectId()
	counts = make(map[int64]int64, len(endpointsTotal))
	for key, value := range endpointsTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[int64(key)] = value.Id
	}
	return
}

func (s *SummaryDetailsService) CountExecTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountExecTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountAllExecTotalProjectId() (counts map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	execsTotal, _ := r.CountAllExecTotalProjectId()
	counts = make(map[int64]int64, len(execsTotal))
	for key, value := range execsTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[int64(key)] = value.Id
	}
	return
}

func (s *SummaryDetailsService) FindPassRateByProjectId(projectId int64) (passRate float64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindPassRateByProjectId(projectId)
}

func (s *SummaryDetailsService) FindAllPassRateByProjectId() (passRate map[int64]float64, err error) {
	r := repo.NewSummaryDetailsRepo()
	passRates, _ := r.FindAllPassRateByProjectId()

	passRate = make(map[int64]float64, len(passRates))

	for projectId, rate := range passRates {
		passRate[int64(projectId)] = rate.Coverage
	}
	return
}

func (s *SummaryDetailsService) HasDataOfDate(startTime string, endTime string, projectId int64) (ret bool, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.HasDataOfDate(startTime, endTime, projectId)
}

//func (s *SummaryDetailsService) CheckCardUpdated(lastUpdateTime *time.Time) (result bool, err error) {
//	r := repo.NewSummaryDetailsRepo()
//	return r.CheckCardUpdated(lastUpdateTime)
//}

//检查是否有今日数据,没有则copy最后一条,然后进行数据是否更新检查
//func (s *SummaryDetailsService) CheckDetailsUpdated(lastUpdateTime *time.Time) (result bool, err error) {
//	r := repo.NewSummaryDetailsRepo()
//	now := time.Now()
//	startTime, endTime := GetDate(now)
//	ret, err := s.HasDataOfDate(startTime, endTime)
//	if !ret {
//		details, _ := s.Find()
//		for _, detail := range details {
//			newDetail := s.CopyDetailsWithoutBaseModel(detail)
//			s.Create(newDetail)
//		}
//	}
//	return r.CheckDetailsUpdated(lastUpdateTime)
//}

func (s *SummaryDetailsService) CollectDetail(projectId int64, ScenariosTotal int64, interfacesTotal int64, execsTotal int64, passRates float64, coverages int64) (ret model.SummaryDetails) {
	ret.ProjectId = projectId
	ret.ScenarioTotal = ScenariosTotal
	ret.InterfaceTotal = interfacesTotal
	ret.ExecTotal = execsTotal
	ret.PassRate, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", passRates), 64)
	coveragesCount := coverages

	var coverage float64
	if ret.InterfaceTotal != 0 {
		coverage = float64(coveragesCount / ret.InterfaceTotal)
	} else {
		coverage = 0
	}
	ret.Coverage, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", coverage), 64)
	return
}

func (s *SummaryDetailsService) CollectDetailByProjectId(projectId int64) (ret model.SummaryDetails) {
	//从biz_scenario表根据projectid,查找场景总数
	ret.ScenarioTotal, _ = s.CountScenarioTotalProjectId(projectId)

	//从biz_interface表根据projectid,查找接口总数
	ret.InterfaceTotal, _ = s.CountEndpointTotalProjectId(projectId)

	//根据projectid,从biz_scenario_report表,获得所有报告总数,然后计算
	ret.ExecTotal, _ = s.CountExecTotalProjectId(projectId)

	//从biz_scenario_report拿到assertion的相关数据,计算后存储
	passRate, _ := s.FindPassRateByProjectId(projectId)
	ret.PassRate, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", passRate), 64)

	//从biz_interface需要获取当前项目的所有接口,然后从biz_processor_interface检查哪些在场景中出现过
	interfaceIds, _ := s.FindEndpointIdsByProjectId(projectId)
	count, _ := s.CoverageByProjectId(projectId, interfaceIds)
	var coverage float64
	if ret.InterfaceTotal != 0 {
		coverage = float64(count / ret.InterfaceTotal)
	} else {
		coverage = 0
	}
	ret.Coverage, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", coverage), 64)
	return
}
