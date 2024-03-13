package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/jinzhu/copier"
	"strconv"
)

type SummaryDetailsService struct {
	SummaryDetailsRepo *repo.SummaryDetailsRepo `inject:""`
	UserRepo           *repo.UserRepo           `inject:""`
}

func (s *SummaryDetailsService) Card(tenantId consts.TenantId, projectId int64) (res v1.ResSummaryCard, err error) {
	var summaryCardTotal, oldSummaryCardTotal model.SummaryCardTotal

	startTime, endTime := GetEarlierDateUntilTodayStartAndEndTime(-30)

	if projectId == 0 {
		summaryCardTotal, err = s.SummaryCard(tenantId)
		res.ProjectTotal, err = s.Count(tenantId)
		oldSummaryCardTotal, err = s.SummaryCardByDate(tenantId, startTime, endTime)
		res.UserTotal, err = s.CountUserTotal(tenantId)
	} else {
		summaryCardTotal, err = s.SummaryCardByProjectId(tenantId, projectId)
		res.ProjectTotal = 1
		oldSummaryCardTotal, err = s.SummaryCardByDateAndProjectId(tenantId, startTime, endTime, projectId)
		res.UserTotal, err = s.CountProjectUserTotal(tenantId, projectId)
	}

	copier.CopyWithOption(&res, summaryCardTotal, copier.Option{DeepCopy: true})

	if oldSummaryCardTotal.Coverage != 0 {
		res.CoverageHb, err = strconv.ParseFloat(fmt.Sprintf("%.2f", res.Coverage-oldSummaryCardTotal.Coverage), 64)
	}

	if oldSummaryCardTotal.InterfaceTotal != 0 {
		res.InterfaceHb, err = strconv.ParseFloat(fmt.Sprintf("%.2f", DecimalHB(float64(res.InterfaceTotal), float64(oldSummaryCardTotal.InterfaceTotal))), 64)
	}

	if oldSummaryCardTotal.ScenarioTotal != 0 {
		res.ScenarioHb, err = strconv.ParseFloat(fmt.Sprintf("%.2f", DecimalHB(float64(res.ScenarioTotal), float64(oldSummaryCardTotal.ScenarioTotal))), 64)
	}

	return
}

func (s *SummaryDetailsService) Details(tenantId consts.TenantId, userId int64) (res v1.ResSummaryDetail, err error) {
	//从project表收集项目总数
	res.ProjectTotal, err = s.Count(tenantId)
	res.UserProjectTotal, err = s.CountByUserId(tenantId, userId)
	//查找所有项目对应的summaryDetail数据，并转为map
	allDetails, err := s.GetAllDetailGroupByProjectId(tenantId)
	//查找用户参与的项目id,并转为map
	userProjectIds, err := s.FindProjectIdsByUserId(tenantId, userId)
	//查询所有项目信息
	allProjectsInfo, err := s.FindAllProjectInfo(tenantId)
	//组装返回的json结构体
	res.ProjectList, res.UserProjectList, err = s.HandleSummaryDetails(tenantId, userId, userProjectIds, allDetails, allProjectsInfo)
	return
}

func (s *SummaryDetailsService) HandleSummaryDetails(tenantId consts.TenantId, userId int64, userProjectIds []int64, allDetails map[int64]model.SummaryDetails, allProjectsInfo []model.SummaryProjectInfo) (resAllDetails []v1.ResSummaryDetails, resUserDetails []v1.ResSummaryDetails, err error) {

	isAdminUser, err := s.UserRepo.IsAdminUser(tenantId, uint(userId))
	if err != nil {
		return
	}

	projectsBugCount, err := s.CountBugsGroupByProjectId(tenantId)
	projectsUsers, err := s.FindAllUserIdAndNameOfProject(tenantId)
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
		if isAdminUser { //管理员所有项目都能访问
			resDetail.Accessible = 1
		}

		//当前项目如果是用户参与的项目，则添加到resUserDetails中
		for userProjectIdsIndex, id := range userProjectIds {
			if int64(projectInfo.ID) == id {
				resDetail.Id = uint(userProjectIdsIndex + 1)
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

func (s *SummaryDetailsService) GetAllDetailGroupByProjectId(tenantId consts.TenantId) (ret map[int64]model.SummaryDetails, err error) {

	//查找所有项目id
	projectIds, err := s.FindProjectIds(tenantId)

	//从biz_scenario表根据projectid,查找场景总数
	ScenariosTotal, err := s.CountAllScenarioTotalProjectId(tenantId)

	//根据projectid,从biz_scenario_report表,获得所有报告总数,然后计算
	execsTotal, err := s.CountAllExecTotalProjectId(tenantId)

	//从biz_interface表根据projectid,查找接口总数
	interfacesTotal, err := s.CountAllEndpointTotalProjectId(tenantId)

	//从biz_scenario_report拿到assertion的相关数据,计算后存储
	passRates, err := s.FindAllPassRateByProjectId(tenantId)

	//通过processorInterface、biz_scenario_report、biz_exec_log_processor联合查询，取出来所有被测试过的接口数量，根据project_id分组
	execLogProcessorInterfaceTotal, err := s.FindAllExecLogProcessorInterfaceTotalGroupByProjectId(tenantId)

	ret = make(map[int64]model.SummaryDetails, len(projectIds))

	for _, projectId := range projectIds {
		details, _ := s.HandleDetail(projectId, ScenariosTotal[projectId], interfacesTotal[projectId], execsTotal[projectId], passRates[projectId], execLogProcessorInterfaceTotal[projectId])
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
	ret.PassRate, err = strconv.ParseFloat(fmt.Sprintf("%.2f", passRates), 64)

	//通过processorInterface、biz_scenario_report、biz_exec_log_processor联合查询，取出来所有被测试过的接口数量，根据project_id分组（跳过0值）
	//然后除以通过processorInterface中对应项目的接口总数
	var coverage float64
	if ret.InterfaceTotal != 0 {
		coverage = float64(endPointCountOfProcessor) / float64(ret.InterfaceTotal) * 100
	} else {
		coverage = 0
	}
	ret.Coverage, err = strconv.ParseFloat(fmt.Sprintf("%.2f", coverage), 64)
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

func (s *SummaryDetailsService) HandlerSummaryDetailsRepo() *repo.SummaryDetailsRepo {
	return s.SummaryDetailsRepo
}

func (s *SummaryDetailsService) Count(tenantId consts.TenantId) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().Count(tenantId)
}

func (s *SummaryDetailsService) CountByUserId(tenantId consts.TenantId, userId int64) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountByUserId(tenantId, userId)
}

func (s *SummaryDetailsService) CountUserTotal(tenantId consts.TenantId) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountUserTotal(tenantId)
}

func (s *SummaryDetailsService) CountProjectUserTotal(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountProjectUserTotal(tenantId, projectId)
}

func (s *SummaryDetailsService) FindAllProjectInfo(tenantId consts.TenantId) (projectDetails []model.SummaryProjectInfo, err error) {
	return s.HandlerSummaryDetailsRepo().FindAllProjectInfo(tenantId)

}

func (s *SummaryDetailsService) FindAllAdminNameByAdminId(tenantId consts.TenantId, adminId int64) (adminName string, err error) {
	return s.HandlerSummaryDetailsRepo().FindAdminNameByAdminId(tenantId, adminId)
}

func (s *SummaryDetailsService) FindProjectIdsByUserId(tenantId consts.TenantId, userId int64) (count []int64, err error) {
	return s.HandlerSummaryDetailsRepo().FindProjectIdsByUserId(tenantId, userId)
}

func (s *SummaryDetailsService) FindEndpointIdsByProjectId(tenantId consts.TenantId, projectId int64) (ids []int64, err error) {
	return s.HandlerSummaryDetailsRepo().FindEndpointIdsByProjectId(tenantId, projectId)
}

func (s *SummaryDetailsService) FindAllEndpointIdsGroupByProjectId(tenantId consts.TenantId, projectIds []int64) (ids map[int64][]int64, err error) {
	results, err := s.HandlerSummaryDetailsRepo().FindAllEndpointIdsGroupByProjectId(tenantId)
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

func (s *SummaryDetailsService) FindExecLogProcessorInterfaceTotalGroupByProjectId(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	count, err = s.HandlerSummaryDetailsRepo().FindExecLogProcessorInterfaceTotalGroupByProjectId(tenantId, projectId)
	return
}

func (s *SummaryDetailsService) FindAllExecLogProcessorInterfaceTotal(tenantId consts.TenantId) (count int64, err error) {
	count, err = s.HandlerSummaryDetailsRepo().FindAllExecLogProcessorInterfaceTotal(tenantId)
	return
}

func (s *SummaryDetailsService) FindAllExecLogProcessorInterfaceTotalGroupByProjectId(tenantId consts.TenantId) (counts map[int64]int64, err error) {
	result, err := s.HandlerSummaryDetailsRepo().FindAllExecLogProcessorInterfaceTotalGroupByProjectId(tenantId)

	counts = make(map[int64]int64, len(result))
	for _, value := range result {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[value.ProjectId] = value.Id
	}
	return
}

func (s *SummaryDetailsService) FindByProjectId(tenantId consts.TenantId, projectId int64) (summaryDetail model.SummaryDetails, err error) {
	return s.HandlerSummaryDetailsRepo().FindByProjectId(tenantId, projectId)
}

func (s *SummaryDetailsService) Find(tenantId consts.TenantId) (details []model.SummaryDetails, err error) {
	return s.HandlerSummaryDetailsRepo().Find(tenantId)
}

func (s *SummaryDetailsService) FindByProjectIds(tenantId consts.TenantId, projectIds []int64) (details []model.SummaryDetails, err error) {
	return s.HandlerSummaryDetailsRepo().FindByProjectIds(tenantId, projectIds)
}

func (s *SummaryDetailsService) FindProjectIds(tenantId consts.TenantId) (ids []int64, err error) {
	return s.HandlerSummaryDetailsRepo().FindProjectIds(tenantId)
}

func (s *SummaryDetailsService) SummaryCard(tenantId consts.TenantId) (summaryCardTotal model.SummaryCardTotal, err error) {
	summaryCardTotal.ScenarioTotal, err = s.CountAllScenarioTotal(tenantId)
	summaryCardTotal.ExecTotal, err = s.CountAllExecTotal(tenantId)
	summaryCardTotal.InterfaceTotal, err = s.CountAllEndpointTotal(tenantId)
	summaryCardTotal.PassRate, err = s.FindAllPassRate(tenantId)

	endPointCountOfProcessor, err := s.FindAllExecLogProcessorInterfaceTotal(tenantId)
	var coverage float64
	if summaryCardTotal.InterfaceTotal != 0 {
		coverage = float64(endPointCountOfProcessor) / float64(summaryCardTotal.InterfaceTotal) * 100
	} else {
		coverage = 0
	}
	summaryCardTotal.Coverage, err = strconv.ParseFloat(fmt.Sprintf("%.2f", coverage), 64)

	return
}

func (s *SummaryDetailsService) SummaryCardByDate(tenantId consts.TenantId, startTime string, endTime string) (summaryCardTotal model.SummaryCardTotal, err error) {

	return s.HandlerSummaryDetailsRepo().SummaryCardByDate(tenantId, startTime, endTime)
}

func (s *SummaryDetailsService) SummaryCardByProjectId(tenantId consts.TenantId, projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	summaryCardTotal.ScenarioTotal, err = s.CountScenarioTotalProjectId(tenantId, projectId)
	summaryCardTotal.ExecTotal, err = s.CountExecTotalProjectId(tenantId, projectId)
	summaryCardTotal.InterfaceTotal, err = s.CountEndpointTotalProjectId(tenantId, projectId)
	endPointCountOfProcessor, err := s.FindExecLogProcessorInterfaceTotalGroupByProjectId(tenantId, projectId)

	summaryCardTotal.PassRate, err = s.FindPassRateByProjectId(tenantId, projectId)
	var coverage float64
	if summaryCardTotal.InterfaceTotal != 0 {
		coverage = float64(endPointCountOfProcessor) / float64(summaryCardTotal.InterfaceTotal) * 100
	} else {
		coverage = 0
	}
	summaryCardTotal.Coverage, err = strconv.ParseFloat(fmt.Sprintf("%.2f", coverage), 64)

	return
}

func (s *SummaryDetailsService) SummaryCardByDateAndProjectId(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (summaryCardTotal model.SummaryCardTotal, err error) {
	return s.HandlerSummaryDetailsRepo().SummaryCardByDateAndProjectId(tenantId, startTime, endTime, projectId)
}

func (s *SummaryDetailsService) FindByProjectIdAndDate(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (summaryDetails model.SummaryDetails, err error) {
	return s.HandlerSummaryDetailsRepo().FindByProjectIdAndDate(tenantId, startTime, endTime, projectId)
}

func (s *SummaryDetailsService) FindAllUserIdAndNameOfProject(tenantId consts.TenantId) (users []model.UserIdAndName, err error) {
	return s.HandlerSummaryDetailsRepo().FindAllUserIdAndNameOfProject(tenantId)
}

func (s *SummaryDetailsService) FindCreateUserNameByProjectId(tenantId consts.TenantId, projectId int64) (userName string, err error) {
	return s.HandlerSummaryDetailsRepo().FindCreateUserNameByProjectId(tenantId, projectId)
}

func (s *SummaryDetailsService) CountBugsGroupByProjectId(tenantId consts.TenantId) (bugsCount map[int64]int64, err error) {
	result, err := s.HandlerSummaryDetailsRepo().CountBugsGroupByProjectId(tenantId)
	bugsCount = make(map[int64]int64, len(result))
	for _, value := range result {
		//这里写的是id实际是count，都是int64，不做多余的明明
		bugsCount[value.ProjectId] = value.Count
	}
	return
}

func (s *SummaryDetailsService) CountScenarioTotalProjectId(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountScenarioTotalProjectId(tenantId, projectId)
}

func (s *SummaryDetailsService) CountAllScenarioTotal(tenantId consts.TenantId) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountAllScenarioTotal(tenantId)
}

func (s *SummaryDetailsService) CountAllScenarioTotalProjectId(tenantId consts.TenantId) (scenarioTotal map[int64]int64, err error) {
	scenariosTotal, err := s.HandlerSummaryDetailsRepo().CountAllScenarioTotalProjectId(tenantId)

	scenarioTotal = make(map[int64]int64, len(scenariosTotal))
	for _, value := range scenariosTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		scenarioTotal[int64(value.ProjectId)] = value.Id
	}
	return
}

func (s *SummaryDetailsService) CountEndpointTotalProjectId(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountEndpointInterfaceTotalProjectId(tenantId, projectId)
}

func (s *SummaryDetailsService) CountAllEndpointTotal(tenantId consts.TenantId) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountAllEndpointTotal(tenantId)
}

func (s *SummaryDetailsService) CountAllEndpointTotalProjectId(tenantId consts.TenantId) (counts map[int64]int64, err error) {
	endpointsTotal, err := s.HandlerSummaryDetailsRepo().CountAllEndpointInterfaceTotalProjectId(tenantId)
	counts = make(map[int64]int64, len(endpointsTotal))
	for _, value := range endpointsTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[value.ProjectId] = value.Id
	}
	return
}

func (s *SummaryDetailsService) CountExecTotalProjectId(tenantId consts.TenantId, projectId int64) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountExecTotalProjectId(tenantId, projectId)
}
func (s *SummaryDetailsService) CountAllExecTotal(tenantId consts.TenantId) (count int64, err error) {
	return s.HandlerSummaryDetailsRepo().CountAllExecTotal(tenantId)
}

func (s *SummaryDetailsService) CountAllExecTotalProjectId(tenantId consts.TenantId) (counts map[int64]int64, err error) {
	execsTotal, err := s.HandlerSummaryDetailsRepo().CountAllExecTotalProjectId(tenantId)
	counts = make(map[int64]int64, len(execsTotal))
	for _, value := range execsTotal {
		//这里写的是id实际是count，都是int64，不做多余的明明
		counts[value.ProjectId] = value.Id
	}
	return
}

func (s *SummaryDetailsService) FindPassRateByProjectId(tenantId consts.TenantId, projectId int64) (passRate float64, err error) {
	result, err := s.HandlerSummaryDetailsRepo().FindAssertionCountByProjectId(tenantId, projectId)

	totalCount := result.TotalAssertionNum + result.CheckpointPass + result.CheckpointFail
	passCount := result.PassAssertionNum + result.CheckpointPass

	if totalCount > 0 {
		passRate = float64(passCount) / float64(totalCount) * 100.0
	} else {
		passRate = 0.0
	}
	return
}

func (s *SummaryDetailsService) FindAllPassRate(tenantId consts.TenantId) (passRate float64, err error) {
	result, err := s.HandlerSummaryDetailsRepo().FindAllAssertionCount(tenantId)

	totalCount := result.TotalAssertionNum + result.CheckpointPass + result.CheckpointFail
	passCount := result.PassAssertionNum + result.CheckpointPass

	if totalCount > 0 {
		passRate = float64(passCount) / float64(totalCount) * 100.0
	} else {
		passRate = 0.0
	}
	return

}

func (s *SummaryDetailsService) FindAllPassRateByProjectId(tenantId consts.TenantId) (ret map[int64]float64, err error) {
	result, err := s.HandlerSummaryDetailsRepo().FindAllAssertionCountGroupByProjectId(tenantId)

	ret = make(map[int64]float64, len(result))

	for _, value := range result {
		var passRate float64

		totalCount := value.TotalAssertionNum + value.CheckpointPass + value.CheckpointFail
		passCount := value.PassAssertionNum + value.CheckpointPass

		if totalCount > 0 {
			passRate = float64(passCount) / float64(totalCount) * 100.0
		} else {
			passRate = 0.0
		}
		ret[value.ProjectId] = passRate
	}
	return
}

func (s *SummaryDetailsService) Create(tenantId consts.TenantId, req model.SummaryDetails) (err error) {
	return s.HandlerSummaryDetailsRepo().Create(tenantId, req)
}

func (s *SummaryDetailsService) UpdateColumnsByDate(tenantId consts.TenantId, id int64, req model.SummaryDetails) (err error) {
	return s.HandlerSummaryDetailsRepo().UpdateColumnsByDate(tenantId, id, req)
}

func (s *SummaryDetailsService) HasDataOfDate(tenantId consts.TenantId, startTime string, endTime string, projectId int64) (id int64, err error) {
	return s.HandlerSummaryDetailsRepo().Existed(tenantId, startTime, endTime, projectId)
}

func (s *SummaryDetailsService) CreateByDate(tenantId consts.TenantId, req model.SummaryDetails) (err error) {
	startTime, endTime := GetTodayStartAndEndTime()
	id, err := s.HasDataOfDate(tenantId, startTime, endTime, req.ProjectId)
	if id == 0 {
		err = s.Create(tenantId, req)
	} else {
		err = s.UpdateColumnsByDate(tenantId, id, req)
	}
	return
}

// SaveDetails 查询今日是否已存在当前projectId对应的数据，没有则create，有则update
func (s *SummaryDetailsService) SaveDetails(tenantId consts.TenantId) (err error) {
	details, err := s.GetAllDetailGroupByProjectId(tenantId)

	for _, detail := range details {
		newDetail := s.CopyDetailsWithoutBaseModel(detail)
		err = s.CreateByDate(tenantId, newDetail)
	}

	return
}
