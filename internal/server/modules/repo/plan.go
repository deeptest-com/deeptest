package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/pkg/domain"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type PlanRepo struct {
	DB             *gorm.DB        `inject:""`
	BaseRepo       *BaseRepo       `inject:""`
	ProjectRepo    *ProjectRepo    `inject:""`
	UserRepo       *UserRepo       `inject:""`
	PlanReportRepo *PlanReportRepo `inject:""`
}

func NewPlanRepo() *PlanRepo {
	return &PlanRepo{}
}

func (r *PlanRepo) Paginate(req v1.PlanReqPaginate, projectId int) (data _domain.PageData, err error) {
	var count int64
	var categoryIds []uint

	if req.CategoryId > 0 {
		categoryIds, err = r.BaseRepo.GetAllChildIds(uint(req.CategoryId), model.Category{}.TableName(),
			serverConsts.PlanCategory, projectId)
		if err != nil {
			return
		}
	}

	db := r.DB.Model(&model.Plan{}).
		Joins("LEFT JOIN biz_plan_report r ON biz_plan.id=r.plan_id").
		Where("biz_plan.project_id = ? AND NOT biz_plan.deleted",
			projectId)

	if len(categoryIds) > 0 {
		db.Where("biz_plan.category_id IN(?)", categoryIds)
	} else if req.CategoryId == -1 {
		db.Where("biz_plan.category_id IN(?)", -1)
	}

	if req.Keywords != "" {
		db = db.Where("biz_plan.name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("biz_plan.disabled = ?", commonUtils.IsDisable(req.Enabled))
	}

	if req.Status != "" {
		db = db.Where("biz_plan.status = ?", req.Status)
	}
	if req.AdminId != 0 {
		db = db.Where("biz_plan.admin_id = ?", req.AdminId)
	}
	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count scenario error", zap.String("error:", err.Error()))
		return
	}

	plans := make([]*model.Plan, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&plans).Error
	if err != nil {
		logUtils.Errorf("query scenario error", zap.String("error:", err.Error()))
		return
	}
	r.CombinePassRate(plans)
	r.CombineUserName(plans)
	data.Populate(plans, count, req.Page, req.PageSize)

	return
}

func (r *PlanRepo) CombineUserName(data []*model.Plan) {
	userIds := make([]uint, 0)
	for _, v := range data {
		userIds = append(userIds, v.AdminId)
		userIds = append(userIds, v.UpdateUserId)
	}
	userIds = commonUtils.ArrayRemoveUintDuplication(userIds)

	users, _ := r.UserRepo.FindByIds(userIds)

	userIdNameMap := make(map[uint]string)
	for _, v := range users {
		userIdNameMap[v.ID] = v.Name
	}

	for _, v := range data {
		if adminName, ok := userIdNameMap[v.AdminId]; ok {
			v.AdminName = adminName
		}
		if updateUserName, ok := userIdNameMap[v.UpdateUserId]; ok {
			v.UpdateUserName = updateUserName
		}
	}
}

func (r *PlanRepo) CombinePassRate(data []*model.Plan) {
	for _, v := range data {
		planReport, err := r.PlanReportRepo.GetLastByPlanId(v.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			logUtils.Errorf("get plan report err", zap.String("error:", err.Error()))
			continue
		}
		if planReport.ID == 0 {
			v.TestPassRate = "n/a"
			continue
		}
		if planReport.PassScenarioNum == 0 || planReport.TotalScenarioNum == 0 {
			v.TestPassRate = "0%"
		} else {
			v.TestPassRate = strconv.Itoa(planReport.PassScenarioNum*100/planReport.TotalScenarioNum) + "%"
		}
	}
}

func (r *PlanRepo) Get(id uint) (scenario model.Plan, err error) {
	err = r.DB.Model(&model.Plan{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *PlanRepo) FindByName(scenarioName string, id uint) (scenario model.Plan, err error) {
	db := r.DB.Model(&model.Plan{}).
		Where("name = ? AND NOT deleted", scenarioName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&scenario)

	return
}

func (r *PlanRepo) Create(scenario model.Plan) (ret model.Plan, bizErr *_domain.BizErr) {
	//po, err := r.FindByName(scenario.Name, 0)
	//if po.Name != "" {
	//	bizErr = &_domain.BizErr{Code: _domain.ErrNameExist.Code}
	//	return
	//}

	err := r.DB.Model(&model.Plan{}).Create(&scenario).Error
	if err != nil {
		logUtils.Errorf("add scenario error", zap.String("error:", err.Error()))
		bizErr = &_domain.BizErr{Code: _domain.SystemErr.Code}

		return
	}

	err = r.UpdateSerialNumber(scenario.ID, scenario.ProjectId)
	if err != nil {
		logUtils.Errorf("add plan serial number error", zap.String("error:", err.Error()))
		bizErr = &_domain.BizErr{Code: _domain.SystemErr.Code}

		return
	}

	ret = scenario

	return
}

func (r *PlanRepo) Update(req model.Plan) error {
	values := map[string]interface{}{
		"name":           req.Name,
		"desc":           req.Desc,
		"status":         req.Status,
		"admin_id":       req.AdminId,
		"category_id":    req.CategoryId,
		"test_stage":     req.TestStage,
		"update_user_id": req.UpdateUserId,
		"disabled":       req.Disabled,
	}
	err := r.DB.Model(&req).Where("id = ?", req.ID).Updates(values).Error
	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *PlanRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.Plan{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *PlanRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.Plan{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *PlanRepo) GetChildrenIds(id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE scenario AS (
			SELECT * FROM biz_scenario WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_scenario child, scenario WHERE child.parent_id = scenario.id
		)
		SELECT id FROM scenario WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children scenario error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *PlanRepo) AddScenarios(planId int, scenarioIds []int) (err error) {
	relations, _ := r.ListScenarioRelation(uint(planId))
	existMap := map[uint]bool{}
	for _, item := range relations {
		existMap[item.ScenarioId] = true
	}

	var pos []model.RelaPlanScenario

	for _, id := range scenarioIds {
		if existMap[uint(id)] {
			continue
		}

		po := model.RelaPlanScenario{
			PlanId:     uint(planId),
			ScenarioId: uint(id),
		}
		pos = append(pos, po)
	}

	err = r.DB.Create(&pos).Error

	return
}

func (r *PlanRepo) ListScenario(id uint) (pos []model.Scenario, err error) {
	relations, _ := r.ListScenarioRelation(uint(id))
	var scenarioIds []uint
	for _, item := range relations {
		scenarioIds = append(scenarioIds, item.ScenarioId)
	}

	err = r.DB.
		Where("id IN (?)", scenarioIds).
		Where("NOT deleted").
		Find(&pos).Error

	return
}

func (r *PlanRepo) ListScenarioRelation(id uint) (pos []model.RelaPlanScenario, err error) {
	err = r.DB.
		Where("plan_id=?", id).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *PlanRepo) RemoveScenario(planId int, scenarioId int) (err error) {
	r.DB.Where("plan_id = ? && scenario_id = ?", planId, scenarioId).Delete(&model.RelaPlanScenario{})

	return
}

func (r *PlanRepo) UpdateSerialNumber(id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.Plan{}).Where("id=?", id).Update("serial_number", project.ShortName+"-TP-"+strconv.Itoa(int(id))).Error
	return
}

func (r *PlanRepo) StatusDropDownOptions() map[consts.TestStatus]string {
	return map[consts.TestStatus]string{
		consts.Draft:     "草稿",
		consts.Disabled:  "已禁用",
		consts.ToExecute: "待执行",
		consts.Executed:  "已执行",
	}
}

func (r *PlanRepo) TestStageDropDownOptions() map[consts.TestStage]string {
	return map[consts.TestStage]string{
		consts.UintTest:        "单元测试",
		consts.IntegrationTest: "集成测试",
		consts.SystemTest:      "系统测试",
		consts.AcceptanceTest:  "验收测试",
	}
}
