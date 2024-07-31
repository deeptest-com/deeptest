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
	DB                   *gorm.DB `inject:""`
	*BaseRepo            `inject:""`
	ProjectRepo          *ProjectRepo          `inject:""`
	UserRepo             *UserRepo             `inject:""`
	PlanReportRepo       *PlanReportRepo       `inject:""`
	RelaPlanScenarioRepo *RelaPlanScenarioRepo `inject:""`
}

func (r *PlanRepo) Paginate(tenantId consts.TenantId, req v1.PlanReqPaginate, projectId int) (data _domain.PageData, err error) {
	var count int64
	var categoryIds []uint

	if req.CategoryId > 0 {
		categoryIds, err = r.BaseRepo.GetDescendantIds(tenantId, uint(req.CategoryId), model.Category{}.TableName(),
			serverConsts.PlanCategory, projectId)
		if err != nil {
			return
		}
	}

	db := r.GetDB(tenantId).Model(&model.Plan{}).
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
	r.CombinePassRate(tenantId, plans)
	r.CombineUserName(tenantId, plans)
	data.Populate(plans, count, req.Page, req.PageSize)

	return
}

func (r *PlanRepo) CombineUserName(tenantId consts.TenantId, data []*model.Plan) {
	userIds := make([]uint, 0)
	for _, v := range data {
		userIds = append(userIds, v.AdminId)
		userIds = append(userIds, v.UpdateUserId)
		userIds = append(userIds, v.CreateUserId)
	}
	userIds = commonUtils.ArrayRemoveUintDuplication(userIds)

	users, _ := r.UserRepo.FindByIds(tenantId, userIds)

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

func (r *PlanRepo) CombinePassRate(tenantId consts.TenantId, data []*model.Plan) {
	for _, v := range data {
		planReport, err := r.PlanReportRepo.GetLastByPlanId(tenantId, v.ID)
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

func (r *PlanRepo) Get(tenantId consts.TenantId, id uint) (scenario model.Plan, err error) {
	err = r.GetDB(tenantId).Model(&model.Plan{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *PlanRepo) FindByName(tenantId consts.TenantId, scenarioName string, id uint) (scenario model.Plan, err error) {
	db := r.GetDB(tenantId).Model(&model.Plan{}).
		Where("name = ? AND NOT deleted", scenarioName)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&scenario)

	return
}

func (r *PlanRepo) Create(tenantId consts.TenantId, scenario model.Plan) (ret model.Plan, bizErr *_domain.BizErr) {
	//po, err := r.FindExpressionByName(scenario.Name, 0)
	//if po.Name != "" {
	//	bizErr = &_domain.BizErr{Code: _domain.ErrNameExist.Code}
	//	return
	//}

	err := r.GetDB(tenantId).Model(&model.Plan{}).Create(&scenario).Error
	if err != nil {
		logUtils.Errorf("add scenario error", zap.String("error:", err.Error()))
		bizErr = &_domain.BizErr{Code: _domain.SystemErr.Code}

		return
	}

	scenario.SerialNumber, err = r.UpdateSerialNumber(tenantId, scenario.ID, scenario.ProjectId)
	if err != nil {
		logUtils.Errorf("add plan serial number error", zap.String("error:", err.Error()))
		bizErr = &_domain.BizErr{Code: _domain.SystemErr.Code}

		return
	}

	ret = scenario

	return
}

func (r *PlanRepo) Update(tenantId consts.TenantId, req model.Plan) error {
	values := map[string]interface{}{
		"name":   req.Name,
		"desc":   req.Desc,
		"status": req.Status,
		//		"admin_id":       req.AdminId,  负责人不支持修改
		"category_id":    req.CategoryId,
		"test_stage":     req.TestStage,
		"update_user_id": req.UpdateUserId,
		"disabled":       req.Disabled,
		"deleted":        req.Deleted,
	}
	err := r.GetDB(tenantId).Model(&req).Where("id = ?", req.ID).Updates(values).Error
	if err != nil {
		logUtils.Errorf("update scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *PlanRepo) DeleteById(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Plan{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *PlanRepo) DeleteChildren(tenantId consts.TenantId, ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.Plan{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete scenario error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *PlanRepo) GetChildrenIds(tenantId consts.TenantId, id uint) (ids []int, err error) {
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

func (r *PlanRepo) AddScenarios(tenantId consts.TenantId, planId uint, scenarioIds []uint) (err error) {
	relations, _ := r.ListScenarioRelation(tenantId, planId)
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

	for _, po := range pos {
		po.Ordr = r.RelaPlanScenarioRepo.GetMaxOrder(tenantId, planId)
		err = r.GetDB(tenantId).Create(&po).Error
		if err != nil {
			return
		}
	}

	return
}

func (r *PlanRepo) ListScenario(tenantId consts.TenantId, id uint) (pos []model.Scenario, err error) {
	relations, _ := r.ListScenarioRelation(tenantId, id)
	var scenarioIds []uint
	for _, item := range relations {
		scenarioIds = append(scenarioIds, item.ScenarioId)
	}

	if len(scenarioIds) == 0 {
		return
	}

	err = r.GetDB(tenantId).Model(model.Scenario{}).
		Where("id IN (?)", scenarioIds).
		Where("NOT deleted").Order(fmt.Sprintf("field(id,%s)", commonUtils.UintArrToStr(scenarioIds))).
		Find(&pos).Error

	return
}

func (r *PlanRepo) ListScenarioRelation(tenantId consts.TenantId, id uint) (pos []model.RelaPlanScenario, err error) {
	err = r.GetDB(tenantId).Model(model.RelaPlanScenario{}).
		Where("plan_id=?", id).
		Where("NOT deleted").Order("ordr asc").
		Find(&pos).Error
	return
}

func (r *PlanRepo) RemoveScenario(tenantId consts.TenantId, planId int, scenarioId int) (err error) {
	r.GetDB(tenantId).Where("plan_id = ? && scenario_id = ?", planId, scenarioId).Delete(&model.RelaPlanScenario{})

	return
}

func (r *PlanRepo) RemoveScenarios(tenantId consts.TenantId, planId int, scenarioIds []uint) (err error) {
	r.GetDB(tenantId).Where("plan_id = ? && scenario_id In (?)", planId, scenarioIds).Delete(&model.RelaPlanScenario{})

	return
}

func (r *PlanRepo) UpdateSerialNumber(tenantId consts.TenantId, id, projectId uint) (number string, err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(tenantId, projectId)
	if err != nil {
		return
	}

	number = project.ShortName + "-TP-" + strconv.Itoa(int(id))

	err = r.GetDB(tenantId).Model(&model.Plan{}).Where("id=?", id).Update("serial_number", number).Error

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

func (r *PlanRepo) PlanScenariosPaginate(tenantId consts.TenantId, req v1.PlanScenariosReqPaginate, planId uint) (data _domain.PageData, err error) {
	var count int64

	db := r.GetDB(tenantId).Model(&model.Scenario{}).
		Select("biz_scenario.*, if(c.name  is not null,c.name,'未分类') category_name,r.id ref_id").
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
	req.Field, req.Order = "r.ordr", "asc"
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

func (r *PlanRepo) GetScenarioNumByPlan(tenantId consts.TenantId, planId uint) (num int64, err error) {
	err = r.GetDB(tenantId).Model(model.RelaPlanScenario{}).
		Where("plan_id = ? AND NOT deleted", planId).
		Count(&num).Error
	return
}

func (r *PlanRepo) NotRelationScenarioList(tenantId consts.TenantId, req v1.NotRelationScenarioReqPaginate, projectId int) (data _domain.PageData, err error) {
	relations, _ := r.ListScenarioRelation(tenantId, req.PlanId)
	var scenarioIds []uint
	for _, v := range relations {
		scenarioIds = append(scenarioIds, v.ScenarioId)
	}

	var count int64

	db := r.GetDB(tenantId).Model(&model.Scenario{}).
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

func (r *PlanRepo) GetCategoryCount(tenantId consts.TenantId, result interface{}, projectId uint) (err error) {
	err = r.GetDB(tenantId).Raw("select count(id) count, category_id from "+model.Plan{}.TableName()+" where not deleted and not disabled and project_id=? group by category_id", projectId).Scan(result).Error
	return
}

func (r *PlanRepo) DeleteByCategoryIds(tenantId consts.TenantId, categoryIds []uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.Plan{}).
		Where("category_id IN (?)", categoryIds).
		Update("deleted", 1).Error

	return
}

func (r *PlanRepo) UpdateCurrEnvId(tenantId consts.TenantId, id, currEnvId uint) error {
	return r.GetDB(tenantId).Model(&model.Plan{}).Where("id = ?", id).UpdateColumn("curr_env_id", currEnvId).Error
}

func (r *PlanRepo) MoveScenario(tenantId consts.TenantId, req v1.MoveReq) (err error) {
	destination, err := r.RelaPlanScenarioRepo.Get(tenantId, req.DestinationId)
	if err != nil {
		return
	}

	souurce, err := r.RelaPlanScenarioRepo.Get(tenantId, req.SourceId)
	if err != nil {
		return
	}

	if destination.Ordr < souurce.Ordr {
		err = r.RelaPlanScenarioRepo.IncreaseOrderAfter(tenantId, destination.Ordr, req.PlanId)
	} else {
		err = r.RelaPlanScenarioRepo.DecreaseOrderBefore(tenantId, destination.Ordr, req.PlanId)
	}

	if err != nil {
		return
	}

	err = r.RelaPlanScenarioRepo.UpdateOrdrById(tenantId, req.SourceId, destination.Ordr)

	return
}
