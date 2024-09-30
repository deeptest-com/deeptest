package repo

import (
	"fmt"
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	"github.com/deeptest-com/deeptest/internal/server/core/dao"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/pkg/domain"
	commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ScenarioRepo struct {
	DB                   *gorm.DB `inject:""`
	*BaseRepo            `inject:""`
	ProjectRepo          *ProjectRepo          `inject:""`
	PlanRepo             *PlanRepo             `inject:""`
	RelaPlanScenarioRepo *RelaPlanScenarioRepo `inject:""`
}

func (r *ScenarioRepo) ListByProject(tenantId consts.TenantId, projectId int) (pos []model.Scenario, err error) {
	err = r.GetDB(tenantId).
		Where("project_id=?", projectId).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *ScenarioRepo) Paginate(tenantId consts.TenantId, req v1.ScenarioReqPaginate, projectId int) (data _domain.PageData, err error) {
	var count int64
	var categoryIds []uint

	if req.CategoryId > 0 {
		categoryIds, err = r.BaseRepo.GetDescendantIds(tenantId, uint(req.CategoryId), model.Category{}.TableName(),
			serverConsts.ScenarioCategory, projectId)
		if err != nil {
			return
		}
	}

	db := r.GetDB(tenantId).Model(&model.Scenario{}).
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
		db = db.Where("status in ?", strings.Split(req.Status, ","))
	}
	if req.Priority != "" {
		db = db.Where("priority in ?", strings.Split(req.Priority, ","))
	}
	if req.Type != "" {
		db = db.Where("type in ?", strings.Split(req.Type, ","))
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

func (r *ScenarioRepo) Get(tenantId consts.TenantId, id uint) (scenario model.Scenario, err error) {
	err = r.GetDB(tenantId).Model(&model.Scenario{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *ScenarioRepo) FindByName(tenantId consts.TenantId, scenarioName string, id uint) (scenario model.Scenario, err error) {
	db := r.GetDB(tenantId).Model(&model.Scenario{}).
		Where("name = ? AND NOT deleted", scenarioName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&scenario)

	return
}

func (r *ScenarioRepo) Create(tenantId consts.TenantId, scenario model.Scenario) (ret model.Scenario, err error) {
	//po, err := r.FindExpressionByName(scenario.Name, 0)
	//if po.Name != "" {
	//	bizErr = &_domain.BizErr{Code: _domain.ErrNameExist.Code}
	//	return
	//}

	err = r.GetDB(tenantId).Model(&model.Scenario{}).Create(&scenario).Error
	if err != nil {
		logUtils.Errorf("add scenario error", zap.String("error:", err.Error()))

		return
	}

	err = r.UpdateSerialNumber(tenantId, scenario.ID, scenario.ProjectId)
	if err != nil {
		logUtils.Errorf("update scenario serial number error", zap.String("error:", err.Error()))
		return
	}
	ret = scenario

	return
}

func (r *ScenarioRepo) Update(tenantId consts.TenantId, req model.Scenario) error {
	/*
		values := map[string]interface{}{
			"name":             req.Name,
			"desc":             req.Desc,
			"disabled":         req.Disabled,
			"create_user_id":   req.CreateUserId,
			"create_user_name": req.CreateUserName,
			"priority":         req.Priority,
			"type":             req.Type,
			"status":           req.Status,
		}
	*/
	err := r.GetDB(tenantId).Model(&req).Where("id = ?", req.ID).Updates(req).Error
	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ScenarioRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Scenario{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ScenarioRepo) DeleteChildren(tenantId consts.TenantId, ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.Scenario{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *ScenarioRepo) GetChildrenIds(tenantId consts.TenantId, id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE scenario AS (
			SELECT * FROM biz_scenario WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_scenario child, scenario WHERE child.parent_id = scenario.id
		)
		SELECT id FROM scenario WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.GetDB(tenantId).Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children scenario error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *ScenarioRepo) UpdateSerialNumber(tenantId consts.TenantId, id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(tenantId, projectId)
	if err != nil {
		return
	}

	err = r.GetDB(tenantId).Model(&model.Scenario{}).Where("id=?", id).Update("serial_number", project.ShortName+"-TS-"+strconv.Itoa(int(id))).Error
	return
}

func (r *ScenarioRepo) ListScenarioRelation(tenantId consts.TenantId, id uint) (pos []model.RelaPlanScenario, err error) {
	err = r.GetDB(tenantId).
		Where("scenario_id=?", id).
		Where("NOT deleted").
		Find(&pos).Error
	return
}

func (r *ScenarioRepo) AddPlans(tenantId consts.TenantId, scenarioId uint, planIds []int) (err error) {
	relations, _ := r.ListScenarioRelation(tenantId, scenarioId)
	existMap := map[uint]bool{}
	for _, item := range relations {
		existMap[item.PlanId] = true
	}

	var pos []model.RelaPlanScenario

	for _, id := range planIds {
		if existMap[uint(id)] {
			continue
		}

		po := model.RelaPlanScenario{
			PlanId:     uint(id),
			ScenarioId: scenarioId,
		}
		pos = append(pos, po)
	}

	for _, po := range pos {
		po.Ordr = r.RelaPlanScenarioRepo.GetMaxOrder(tenantId, po.PlanId)
		err = r.GetDB(tenantId).Create(&po).Error
		if err != nil {
			return
		}
	}

	return
}

func (r *ScenarioRepo) PlanList(tenantId consts.TenantId, req v1.ScenarioPlanReqPaginate, scenarioId int) (data _domain.PageData, err error) {
	relations, _ := r.ListScenarioRelation(tenantId, uint(scenarioId))
	var planIds []uint
	for _, item := range relations {
		planIds = append(planIds, item.PlanId)
	}

	db := r.GetDB(tenantId).Model(&model.Plan{}).Where("not deleted and project_id=?", req.ProjectId)

	if req.Ref && len(planIds) == 0 {
		return
	}

	if len(planIds) > 0 {
		if req.Ref {
			db = db.Where(" id in (?)", planIds)
		} else {
			db = db.Where(" id not in (?)", planIds)
		}
	}

	var count int64

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}

	if req.AdminId != 0 {
		db = db.Where("admin_id = ?", req.AdminId)
	}

	if req.Keywords != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count plan error", zap.String("error:", err.Error()))
		return
	}

	plans := make([]*model.Plan, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&plans).Error
	if err != nil {
		logUtils.Errorf("query plan error", zap.String("error:", err.Error()))
		return
	}

	r.PlanRepo.CombinePassRate(tenantId, plans)
	r.PlanRepo.CombineUserName(tenantId, plans)
	data.Populate(plans, count, req.Page, req.PageSize)

	return
}

func (r *ScenarioRepo) UpdateStatus(tenantId consts.TenantId, id uint, status consts.TestStatus, updateUserId uint, updateUserName string) error {
	fields := map[string]interface{}{
		"status":           status,
		"update_user_id":   updateUserId,
		"update_user_name": updateUserName,
	}
	return r.GetDB(tenantId).Model(&model.Scenario{}).Where("id = ?", id).Updates(fields).Error
}

func (r *ScenarioRepo) UpdatePriority(tenantId consts.TenantId, id uint, priority string, updateUserId uint, updateUserName string) error {
	fields := map[string]interface{}{
		"priority":         priority,
		"update_user_id":   updateUserId,
		"update_user_name": updateUserName,
	}
	return r.GetDB(tenantId).Model(&model.Scenario{}).Where("id = ?", id).Updates(fields).Error
}

func (r *ScenarioRepo) GetByIds(tenantId consts.TenantId, ids []uint) (scenarios []model.Scenario, err error) {
	err = r.GetDB(tenantId).Model(&model.Scenario{}).Where("id IN (?)", ids).Find(&scenarios).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenarios, err
	}

	return
}

func (r *ScenarioRepo) RemovePlans(tenantId consts.TenantId, scenarioId uint, planIds []int) (err error) {
	err = r.GetDB(tenantId).Model(&model.RelaPlanScenario{}).Where("scenario_id=? and plan_id in (?)", scenarioId, planIds).Update("deleted", true).Error
	return
}

func (r *ScenarioRepo) GetCategoryCount(tenantId consts.TenantId, result interface{}, projectId uint) (err error) {
	err = r.GetDB(tenantId).Raw("select count(id) count, category_id from "+model.Scenario{}.TableName()+" where not deleted and not disabled and project_id=? group by category_id", projectId).Scan(result).Error
	return
}

func (r *ScenarioRepo) DeleteByCategoryIds(tenantId consts.TenantId, categoryIds []uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Scenario{}).
		Where("category_id IN (?)", categoryIds).
		Update("deleted", 1).Error

	return
}

func (r *ScenarioRepo) UpdateCurrEnvId(tenantId consts.TenantId, id, currEnvId uint) error {
	return r.GetDB(tenantId).Model(&model.Scenario{}).Where("id = ?", id).Update("curr_env_id", currEnvId).Error
}
