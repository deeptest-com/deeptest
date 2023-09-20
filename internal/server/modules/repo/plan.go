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
	"strings"
)

type PlanRepo struct {
	DB             *gorm.DB `inject:""`
	*BaseRepo      `inject:""`
	ProjectRepo    *ProjectRepo    `inject:""`
	UserRepo       *UserRepo       `inject:""`
	PlanReportRepo *PlanReportRepo `inject:""`
}

func (r *PlanRepo) Paginate(req v1.PlanReqPaginate, projectId int) (data _domain.PageData, err error) {
	var count int64
	var categoryIds []uint

	if req.CategoryId > 0 {
		categoryIds, err = r.BaseRepo.GetDescendantIds(uint(req.CategoryId), model.Category{}.TableName(),
			serverConsts.PlanCategory, projectId)
		if err != nil {
			return
		}
	}

	db := r.DB.Model(&model.Plan{}).
		Where("project_id = ? AND NOT deleted",
			projectId)

	if len(categoryIds) > 0 {
		db.Where("category_id IN(?)", categoryIds)
	} else if req.CategoryId == -1 {
		db.Where("category_id IN(?)", -1)
	}

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", commonUtils.IsDisable(req.Enabled))
	}
	if req.Status != "" {
		db = db.Where("status IN (?)", strings.Split(req.Status, ","))
	}
	if req.AdminId != "" {
		db = db.Where("admin_id IN (?)", strings.Split(req.AdminId, ","))
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
		userIds = append(userIds, v.CreateUserId)
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
		if createUserName, ok := userIdNameMap[v.CreateUserId]; ok {
			v.CreateUserName = createUserName
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
	//po, err := r.FindExpressionByName(scenario.Name, 0)
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

func (r *PlanRepo) AddScenarios(planId uint, scenarioIds []uint) (err error) {
	relations, _ := r.ListScenarioRelation(planId)
	existMap := map[uint]bool{}
	for _, item := range relations {
		existMap[item.ScenarioId] = true
	}

	var pos []model.RelaPlanScenario

	for _, id := range scenarioIds {
		if existMap[id] {
			continue
		}

		po := model.RelaPlanScenario{
			PlanId:     planId,
			ScenarioId: id,
		}
		pos = append(pos, po)
	}
	if len(pos) == 0 {
		return
	}
	err = r.DB.Create(&pos).Error

	return
}

func (r *PlanRepo) ListScenario(id uint) (pos []model.Scenario, err error) {
	relations, _ := r.ListScenarioRelation(id)
	var scenarioIds []uint
	for _, item := range relations {
		scenarioIds = append(scenarioIds, item.ScenarioId)
	}

	err = r.DB.Model(model.Scenario{}).
		Where("id IN (?)", scenarioIds).
		Where("NOT deleted").
		Find(&pos).Error

	return
}

func (r *PlanRepo) ListScenarioRelation(id uint) (pos []model.RelaPlanScenario, err error) {
	err = r.DB.Model(model.RelaPlanScenario{}).
		Where("plan_id=?", id).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *PlanRepo) RemoveScenario(planId int, scenarioId int) (err error) {
	r.DB.Where("plan_id = ? && scenario_id = ?", planId, scenarioId).Delete(&model.RelaPlanScenario{})

	return
}

func (r *PlanRepo) RemoveScenarios(planId int, scenarioIds []uint) (err error) {
	r.DB.Where("plan_id = ? && scenario_id In (?)", planId, scenarioIds).Delete(&model.RelaPlanScenario{})

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

func (r *PlanRepo) PlanScenariosPaginate(req v1.PlanScenariosReqPaginate, planId uint) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.Scenario{}).
		Select("biz_scenario.*, c.name category_name").
		Joins("LEFT JOIN biz_plan_scenario_r r ON biz_scenario.id=r.scenario_id").
		Joins("LEFT JOIN biz_category c ON biz_scenario.category_id=c.id").
		Where("r.plan_id = ? AND NOT biz_scenario.deleted AND NOT r.deleted", planId)

	if req.Keywords != "" {
		db = db.Where("biz_scenario.name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("biz_scenario.disabled = ?", commonUtils.IsDisable(req.Enabled))
	}
	if req.Priority != "" {
		db = db.Where("biz_scenario.priority = ?", req.Priority)
	}
	if req.CreateUserId != 0 {
		db = db.Where("biz_scenario.create_user_id = ?", req.CreateUserId)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count scenario error", zap.String("error:", err.Error()))
		return
	}

	scenarios := make([]*model.ScenarioDetail, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&scenarios).Error
	if err != nil {
		logUtils.Errorf("query scenario error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(scenarios, count, req.Page, req.PageSize)

	return
}

func (r *PlanRepo) GetScenarioNumByPlan(planId uint) (num int64, err error) {
	err = r.DB.Model(model.RelaPlanScenario{}).
		Where("plan_id = ? AND NOT deleted", planId).
		Count(&num).Error
	return
}

func (r *PlanRepo) NotRelationScenarioList(req v1.NotRelationScenarioReqPaginate, projectId int) (data _domain.PageData, err error) {
	relations, _ := r.ListScenarioRelation(req.PlanId)
	var scenarioIds []uint
	for _, v := range relations {
		scenarioIds = append(scenarioIds, v.ScenarioId)
	}

	var count int64

	db := r.DB.Model(&model.Scenario{}).
		Where("project_id = ? AND NOT deleted",
			projectId)

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Enabled != "" {
		db = db.Where("disabled = ?", commonUtils.IsDisable(req.Enabled))
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.Priority != "" {
		db = db.Where("priority = ?", req.Priority)
	}
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}
	if req.CreateUserId != 0 {
		db = db.Where("create_user_id = ?", req.CreateUserId)
	}
	if len(scenarioIds) > 0 {
		db = db.Where("id NOT IN (?)", scenarioIds)

	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count scenario error", zap.String("error:", err.Error()))
		return
	}

	scenarios := make([]*model.Scenario, 0)
	req.Order = "desc"
	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&scenarios).Error
	if err != nil {
		logUtils.Errorf("query scenario error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(scenarios, count, req.Page, req.PageSize)

	return
}

func (r *PlanRepo) GetCategoryCount(result interface{}, projectId uint) (err error) {
	err = r.DB.Raw("select count(id) count, category_id from "+model.Plan{}.TableName()+" where not deleted and not disabled and project_id=? group by category_id", projectId).Scan(result).Error
	return
}

func (r *PlanRepo) DeleteByCategoryIds(categoryIds []uint) (err error) {
	err = r.DB.Model(&model.Plan{}).
		Where("category_id IN (?)", categoryIds).
		Update("deleted", 1).Error

	return
}

func (r *PlanRepo) UpdateCurrEnvId(id, currEnvId uint) error {
	return r.DB.Model(&model.Plan{}).Where("id = ?", id).UpdateColumn("curr_env_id", currEnvId).Error
}
