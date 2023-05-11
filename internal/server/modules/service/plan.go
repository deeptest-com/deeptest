package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	"gorm.io/gorm"
	"strconv"
)

type PlanService struct {
	PlanRepo       *repo.PlanRepo       `inject:""`
	PlanReportRepo *repo.PlanReportRepo `inject:""`
	UserRepo       *repo.UserRepo       `inject:""`
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (s *PlanService) Paginate(req v1.PlanReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.PlanRepo.Paginate(req, projectId)

	if err != nil {
		return
	}

	return
}

func (s *PlanService) GetById(id uint, detail bool) (ret v1.PlanAndReportDetail, err error) {
	userIds := make([]uint, 0)
	plan, err := s.PlanRepo.Get(id)
	if err != nil {
		return
	}
	userIds = append(userIds, plan.AdminId)
	userIds = append(userIds, plan.UpdateUserId)

	lastPlanReport, err := s.PlanReportRepo.GetLastByPlanId(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	planExecTimes, err := s.PlanReportRepo.GetPlanExecNumber(id)
	if err != nil {
		return
	}

	testPassRate := ""
	if lastPlanReport.ID == 0 {
		testPassRate = "n/a"
	} else {
		userIds = append(userIds, lastPlanReport.CreateUserId)

		if lastPlanReport.TotalScenarioNum == 0 {
			testPassRate = "0%"
		} else {
			testPassRate = strconv.Itoa(lastPlanReport.PassScenarioNum*100/lastPlanReport.TotalScenarioNum) + "%"
		}
	}

	userIdNameMap := s.UserRepo.GetUserIdNameMap(userIds)

	if adminName, ok := userIdNameMap[plan.AdminId]; ok {
		ret.AdminName = adminName
	}
	if updateUserName, ok := userIdNameMap[plan.UpdateUserId]; ok {
		ret.UpdateUserName = updateUserName
	}
	ret.Id = plan.ID
	ret.CreatedAt = plan.CreatedAt
	ret.UpdatedAt = plan.UpdatedAt
	ret.Status = plan.Status
	ret.TestPassRate = testPassRate
	ret.ExecTimes = planExecTimes
	if lastPlanReport.ID != 0 {
		if executorName, ok := userIdNameMap[lastPlanReport.CreateUserId]; ok {
			ret.ExecutorName = executorName
		}
		ret.ExecTime = lastPlanReport.CreatedAt
		ret.ExecEnv = lastPlanReport.ExecEnv
	}

	//if detail {
	//	ret.Scenarios, _ = s.PlanRepo.ListScenario(id)
	//}

	return
}

func (s *PlanService) Create(req model.Plan) (po model.Plan, bizErr *_domain.BizErr) {
	po, bizErr = s.PlanRepo.Create(req)

	return
}

func (s *PlanService) Update(req model.Plan) error {
	return s.PlanRepo.Update(req)
}

func (s *PlanService) DeleteById(id uint) error {
	return s.PlanRepo.DeleteById(id)
}

func (s *PlanService) AddScenarios(planId int, scenarioIds []int) (err error) {
	err = s.PlanRepo.AddScenarios(planId, scenarioIds)
	return
}

func (s *PlanService) RemoveScenario(planId int, scenarioId int) (err error) {
	err = s.PlanRepo.RemoveScenario(planId, scenarioId)
	return
}

func (s *PlanService) StatusDropDownOptions() map[consts.TestStatus]string {
	return s.PlanRepo.StatusDropDownOptions()
}

func (s *PlanService) TestStageDropDownOptions() map[consts.TestStage]string {
	return s.PlanRepo.TestStageDropDownOptions()
}
