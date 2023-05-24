package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"strconv"
)

type SummaryDetailsService struct {
	SummaryDetailsRepo *repo.SummaryDetailsRepo `inject:""`
}

func NewSummaryDetailsService() *SummaryDetailsService {
	return new(SummaryDetailsService)
}

func (s *SummaryDetailsService) Card(projectId int64) (res v1.ResSummaryCard, err error) {
	var summaryCardTotal, oldSummaryCardTotal model.SummaryCardTotal

	startTime, endTime := GetEarlierDateUntilTodayStartAndEndTime(-30)

	if projectId == 0 {
		summaryCardTotal, err = s.SummaryCard()
		res.ProjectTotal, err = s.Count()
		oldSummaryCardTotal, err = s.SummaryCardByDate(startTime, endTime)
		res.UserTotal, err = s.CountUserTotal()
	} else {
		summaryCardTotal, err = s.SummaryCardByProjectId(projectId)
		res.ProjectTotal = 1
		oldSummaryCardTotal, err = s.SummaryCardByDateAndProjectId(startTime, endTime, projectId)
		res.UserTotal, err = s.CountProjectUserTotal(projectId)
	}

	copier.CopyWithOption(&res, summaryCardTotal, copier.Option{DeepCopy: true})

	if oldSummaryCardTotal.Coverage != 0 {
		res.CoverageHb, err = strconv.ParseFloat(fmt.Sprintf("%.1f", DecimalHB(res.Coverage, oldSummaryCardTotal.Coverage)), 64)
	}

	if oldSummaryCardTotal.InterfaceTotal != 0 {
		res.InterfaceHb, err = strconv.ParseFloat(fmt.Sprintf("%.1f", DecimalHB(float64(res.InterfaceTotal), float64(oldSummaryCardTotal.InterfaceTotal))), 64)
	}

	if oldSummaryCardTotal.ScenarioTotal != 0 {
		res.ScenarioHb, err = strconv.ParseFloat(fmt.Sprintf("%.1f", DecimalHB(float64(res.ScenarioTotal), float64(oldSummaryCardTotal.ScenarioTotal))), 64)
	}

	return
}

func (s *SummaryDetailsService) Details(userId int64) (res v1.ResSummaryDetail, err error) {
	//从project表收集项目总数
	res.ProjectTotal, err = s.Count()
	res.UserProjectTotal, err = s.CountByUserId(userId)
	//查找所有项目对应的summaryDetail数据，并转为map
	allDetails, err := s.GetAllDetailGroupByProjectId()
	//查找用户参与的项目id,并转为map
	userProjectIds, err := s.FindProjectIdsByUserId(userId)
	//查询所有项目信息
	allProjectsInfo, err := s.FindAllProjectInfo()
	//组装返回的json结构体
	res.ProjectList, res.UserProjectList, err = s.HandleSummaryDetails(userProjectIds, allDetails, allProjectsInfo)
	return
}

func (s *SummaryDetailsService) HandleSummaryDetails(userProjectIds []int64, allDetails map[int64]model.SummaryDetails, allProjectsInfo []model.SummaryProjectInfo) (resAllDetails []v1.ResSummaryDetails, resUserDetails []v1.ResSummaryDetails, err error) {
	projectsBugCount, err := s.CountBugsGroupByProjectId()
	projectsUsers, err := s.FindAllUserIdAndNameOfProject()
	projectsUserListGroupByProject := s.LetUsersGroupByProjectId(allProjectsInfo, projectsUsers)

	//遍历项目信息，匹配details表结果，进行字段复制，组装返回resAllDetails体
	for allProjectsInfoIndex, projectInfo := range allProjectsInfo {
		var resDetail v1.ResSummaryDetails
		var Detail model.SummaryDetails
		Detail = allDetails[int64(projectInfo.ID)]
		resDetail = s.CopyProjectInfo(projectInfo, Detail)
		resDetail.Id = uint(allProjectsInfoIndex + 1)
		resDetail.BugTotal = projectsBugCount[int64(projectInfo.ID)]
		resDetail.UserList = projectsUserListGroupByProject[int64(projectInfo.ID)]

		//当前项目如果是用户参与的项目，则添加到resUserDetails中
		for userProjectIdsIndex, id := range userProjectIds {
			if int64(projectInfo.ID) == id {
				resDetail.Id = uint(userProjectIdsIndex + 1)
				resDetail.Accessible = 1
				resDetail.Accessible = 1
				resUserDetails = append(resUserDetails, resDetail)
				break
			}
		}

		resAllDetails = append(resAllDetails, resDetail)

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
	ret.ScenarioTotal = detail.ScenarioTotal
	ret.ProjectId = detail.ProjectId
	ret.Coverage = detail.Coverage
	ret.ExecTotal = detail.ExecTotal
	ret.PassRate = detail.PassRate
	ret.InterfaceTotal = detail.InterfaceTotal
	return
}

func (s *SummaryDetailsService) GetAllDetailGroupByProjectId() (ret map[int64]model.SummaryDetails, err error) {

	//查找所有项目id
	projectIds, err := s.FindProjectIds()

	//从biz_scenario表根据projectid,查找场景总数
	ScenariosTotal, err := s.CountAllScenarioTotalProjectId()

	//根据projectid,从biz_scenario_report表,获得所有报告总数,然后计算
	execsTotal, err := s.CountAllExecTotalProjectId()

	//从biz_interface表根据projectid,查找接口总数
	interfacesTotal, err := s.CountAllEndpointTotalProjectId()

	//从biz_scenario_report拿到assertion的相关数据,计算后存储
	passRates, err := s.FindAllPassRateByProjectId()

	//从processor里边，根据project_id，取出来对应的，所有endPointId总数，其中endPointId不重复
	endPointCountOfProcessor, err := s.FindAllProcessEndpointCountGroupByProjectId()

	ret = make(map[int64]model.SummaryDetails, len(projectIds))

	for _, projectId := range projectIds {
		details, _ := s.HandleDetail(projectId, ScenariosTotal[projectId], interfacesTotal[projectId], execsTotal[projectId], passRates[projectId], endPointCountOfProcessor[projectId])
		//返回的数组，需要处理成map形式
		ret[projectId] = details
	}
	return
}

func (s *SummaryDetailsService) HandleDetail(projectId int64, ScenariosTotal int64, interfacesTotal int64, execsTotal int64, passRates float64, endPointCountOfProcessor int64) (ret model.SummaryDetails, err error) {
	ret.ProjectId = projectId
	ret.ScenarioTotal = ScenariosTotal
	ret.InterfaceTotal = interfacesTotal
	ret.ExecTotal = execsTotal
	ret.PassRate, err = strconv.ParseFloat(fmt.Sprintf("%.1f", passRates), 64)

	var coverage float64
	if ret.InterfaceTotal != 0 {
		coverage = float64(endPointCountOfProcessor) / float64(ret.InterfaceTotal) * 100
	} else {
		coverage = 0
	}
	ret.Coverage, err = strconv.ParseFloat(fmt.Sprintf("%.1f", coverage), 64)
	return
}

func (s *SummaryDetailsService) LetUsersGroupByProjectId(projectsInfo []model.SummaryProjectInfo, projectsUsers []model.UserIdAndName) (ret map[int64][]v1.ResUserIdAndName) {
	//将拿到的userList，根据projectid，装入map中，key是projectid，value是userList    []v1.ResUserIdAndName
	var m1 map[int64][]v1.ResUserIdAndName
	m1 = make(map[int64][]v1.ResUserIdAndName, len(projectsInfo))
	for _, projectInfo := range projectsInfo {
		var tmpUsers []v1.ResUserIdAndName
		for _, projectUsers := range projectsUsers {
			if projectUsers.ProjectId == int64(projectInfo.ID) {
				var tmpUser v1.ResUserIdAndName
				copier.CopyWithOption(&tmpUser, projectUsers, copier.Option{DeepCopy: true})
				tmpUsers = append(tmpUsers, tmpUser)
			}
		}
		m1[int64(projectInfo.ID)] = tmpUsers
	}
	return m1
}

func (s *SummaryDetailsService) Count() (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Count()
}

func (s *SummaryDetailsService) CountByUserId(userId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountByUserId(userId)
}

func (s *SummaryDetailsService) CountUserTotal() (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountUserTotal()
}

func (s *SummaryDetailsService) CountProjectUserTotal(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountProjectUserTotal(projectId)
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
		var tmpEndpointIds []int64
		for _, result := range results {
			if result.ProjectId == projectId {
				tmpEndpointIds = append(tmpEndpointIds, result.Id)
			}
		}
		ids[projectId] = tmpEndpointIds
	}
	return
}

func (s *SummaryDetailsService) CoverageByProjectId(projectId int64, interfaceIds []int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CoverageByProjectId(projectId, interfaceIds)
}

func (s *SummaryDetailsService) FindAllProcessEndpointCountGroupByProjectId() (counts map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	result, err := r.FindAllProcessEndpointCountGroupByProjectId()

	counts = make(map[int64]int64, len(result))
	for _, value := range result {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[value.ProjectId] = value.Id
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

func (s *SummaryDetailsService) SummaryCardByProjectId(projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.SummaryCardByProjectId(projectId)
}

func (s *SummaryDetailsService) SummaryCardByDateAndProjectId(startTime string, endTime string, projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.SummaryCardByDateAndProjectId(startTime, endTime, projectId)
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

func (s *SummaryDetailsService) CountBugsGroupByProjectId() (bugsCount map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	result, err := r.CountBugsGroupByProjectId()
	bugsCount = make(map[int64]int64, len(result))
	for _, value := range result {
		//这里写的是id实际是count，都是int64，不做多余的明明
		bugsCount[value.ProjectId] = value.Count
	}
	return
}

func (s *SummaryDetailsService) CountScenarioTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountScenarioTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountAllScenarioTotalProjectId() (scenarioTotal map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	scenariosTotal, err := r.CountAllScenarioTotalProjectId()

	scenarioTotal = make(map[int64]int64, len(scenariosTotal))
	for _, value := range scenariosTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		scenarioTotal[int64(value.ProjectId)] = value.Id
	}
	return
}

func (s *SummaryDetailsService) CountEndpointTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountEndpointTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountAllEndpointTotalProjectId() (counts map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	endpointsTotal, err := r.CountAllEndpointTotalProjectId()
	counts = make(map[int64]int64, len(endpointsTotal))
	for _, value := range endpointsTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[value.ProjectId] = value.Id
	}
	return
}

func (s *SummaryDetailsService) CountExecTotalProjectId(projectId int64) (count int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.CountExecTotalProjectId(projectId)
}

func (s *SummaryDetailsService) CountAllExecTotalProjectId() (counts map[int64]int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	execsTotal, err := r.CountAllExecTotalProjectId()
	counts = make(map[int64]int64, len(execsTotal))
	for _, value := range execsTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[value.ProjectId] = value.Id
	}
	return
}

func (s *SummaryDetailsService) FindPassRateByProjectId(projectId int64) (passRate float64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.FindPassRateByProjectId(projectId)
}

func (s *SummaryDetailsService) FindAllPassRateByProjectId() (passRate map[int64]float64, err error) {
	r := repo.NewSummaryDetailsRepo()
	passRates, err := r.FindAllPassRateByProjectId()

	passRate = make(map[int64]float64, len(passRates))

	for _, rate := range passRates {
		passRate[rate.ProjectId] = rate.Coverage
	}
	return
}

func (s *SummaryDetailsService) Create(req model.SummaryDetails) (err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Create(req)
}

func (s *SummaryDetailsService) UpdateColumnsByDate(id int64, req model.SummaryDetails) (err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.UpdateColumnsByDate(id, req)
}

func (s *SummaryDetailsService) HasDataOfDate(startTime string, endTime string, projectId int64) (id int64, err error) {
	r := repo.NewSummaryDetailsRepo()
	return r.Existed(startTime, endTime, projectId)
}

func (s *SummaryDetailsService) CreateByDate(req model.SummaryDetails) (err error) {
	startTime, endTime := GetTodayStartAndEndTime()
	id, err := s.HasDataOfDate(startTime, endTime, req.ProjectId)
	if id == 0 {
		err = s.Create(req)
	} else {
		err = s.UpdateColumnsByDate(id, req)
	}
	return
}

// SaveDetails 查询今日是否已存在当前projectId对应的数据，没有则create，有则update
func (s *SummaryDetailsService) SaveDetails() (err error) {
	details, err := s.GetAllDetailGroupByProjectId()

	for _, detail := range details {
		newDetail := s.CopyDetailsWithoutBaseModel(detail)
		err = s.CreateByDate(newDetail)
	}

	return
}
