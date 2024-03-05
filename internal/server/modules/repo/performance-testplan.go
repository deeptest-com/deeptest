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

type PerformanceTestPlanRepo struct {
	DB               *gorm.DB `inject:""`
	*BaseRepo        `inject:""`
	ProjectRepo      *ProjectRepo      `inject:""`
	ScenarioRepo     *ScenarioRepo     `inject:""`
	ScenarioNodeRepo *ScenarioNodeRepo `inject:""`
}

func (r *PerformanceTestPlanRepo) Paginate(req v1.PerformanceTestPlanReqPaginate, projectId int) (data _domain.PageData, err error) {
	var count int64
	var categoryIds []uint

	if req.CategoryId > 0 {
		categoryIds, err = r.BaseRepo.GetDescendantIds(uint(req.CategoryId), model.Category{}.TableName(),
			serverConsts.PerformanceTestPlanCategory, projectId)
		if err != nil {
			return
		}
	}

	db := r.DB.Model(&model.PerformanceTestPlan{}).
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

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count performanceTestPlan error", zap.String("error:", err.Error()))
		return
	}

	performanceTestPlans := make([]*model.PerformanceTestPlan, 0)
	req.Order = "desc"
	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&performanceTestPlans).Error
	if err != nil {
		logUtils.Errorf("query performanceTestPlan error", zap.String("error:", err.Error()))
		return
	}

	data.Populate(performanceTestPlans, count, req.Page, req.PageSize)

	return
}

func (r *PerformanceTestPlanRepo) Get(id uint) (performanceTestPlan model.PerformanceTestPlan, err error) {
	err = r.DB.Model(&model.PerformanceTestPlan{}).Where("id = ?", id).First(&performanceTestPlan).Error
	if err != nil {
		logUtils.Errorf("find performanceTestPlan by id error", zap.String("error:", err.Error()))
		return performanceTestPlan, err
	}

	return performanceTestPlan, nil
}

func (r *PerformanceTestPlanRepo) Create(performanceTestPlan model.PerformanceTestPlan) (ret model.PerformanceTestPlan, err error) {
	scenario := model.Scenario{

		Status:    consts.Draft,
		DesignFor: consts.DesignForPerformanceTest,

		ProjectId:      performanceTestPlan.ProjectId,
		CreateUserName: performanceTestPlan.CreateUserName,
		CreateUserId:   performanceTestPlan.CreateUserId,
	}
	scenario, err = r.ScenarioRepo.Create(scenario)
	if err != nil {
		return
	}
	root, err := r.ScenarioNodeRepo.CreateDefault(scenario.ID, scenario.ProjectId, scenario.CreateUserId)
	if err != nil {
		return
	}
	err = r.ScenarioNodeRepo.CreateFoldersForPerformance(root.ID, scenario.ID, scenario.ProjectId, scenario.CreateUserId)
	if err != nil {
		return
	}

	performanceTestPlan.ScenarioId = scenario.ID

	err = r.DB.Model(&model.PerformanceTestPlan{}).Create(&performanceTestPlan).Error
	if err != nil {
		logUtils.Errorf("add performanceTestPlan error", zap.String("error:", err.Error()))

		return
	}

	err = r.UpdateSerialNumber(performanceTestPlan.ID, performanceTestPlan.ProjectId)
	if err != nil {
		logUtils.Errorf("update performanceTestPlan serial number error", zap.String("error:", err.Error()))
		return
	}
	ret = performanceTestPlan

	return
}

func (r *PerformanceTestPlanRepo) Update(req model.PerformanceTestPlan) error {
	err := r.DB.Model(&req).Where("id = ?", req.ID).Updates(req).Error
	if err != nil {
		logUtils.Errorf("update performanceTestPlan error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *PerformanceTestPlanRepo) DeleteById(id uint) (err error) {
	po, _ := r.Get(id)

	r.ScenarioRepo.DeleteById(po.ScenarioId)

	err = r.DB.Model(&model.PerformanceTestPlan{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete performanceTestPlan by id error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *PerformanceTestPlanRepo) DeleteChildren(ids []int, tx *gorm.DB) (err error) {
	err = tx.Model(&model.PerformanceTestPlan{}).Where("id IN (?)", ids).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("batch delete performanceTestPlan error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *PerformanceTestPlanRepo) GetChildrenIds(id uint) (ids []int, err error) {
	tmpl := `
		WITH RECURSIVE performanceTestPlan AS (
			SELECT * FROM biz_performanceTestPlan WHERE id = %d
			UNION ALL
			SELECT child.* FROM biz_performanceTestPlan child, performanceTestPlan WHERE child.parent_id = performanceTestPlan.id
		)
		SELECT id FROM performanceTestPlan WHERE id != %d
    `
	sql := fmt.Sprintf(tmpl, id, id)
	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		logUtils.Errorf("get children performanceTestPlan error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *PerformanceTestPlanRepo) UpdateSerialNumber(id, projectId uint) (err error) {
	var project model.Project
	project, err = r.ProjectRepo.Get(projectId)
	if err != nil {
		return
	}

	err = r.DB.Model(&model.PerformanceTestPlan{}).Where("id=?", id).
		Update("serial_number", project.ShortName+"-PT-"+strconv.Itoa(int(id))).Error
	return
}

func (r *PerformanceTestPlanRepo) UpdateStatus(id uint, status consts.TestStatus, updateUserId uint, updateUserName string) error {
	fields := map[string]interface{}{
		"status":           status,
		"update_user_id":   updateUserId,
		"update_user_name": updateUserName,
	}
	return r.DB.Model(&model.PerformanceTestPlan{}).Where("id = ?", id).Updates(fields).Error
}

func (r *PerformanceTestPlanRepo) GetCategoryCount(result interface{}, projectId uint) (err error) {
	err = r.DB.Raw("select count(id) count, category_id from "+model.PerformanceTestPlan{}.TableName()+" where not deleted and not disabled and project_id=? group by category_id", projectId).Scan(result).Error
	return
}

func (r *PerformanceTestPlanRepo) DeleteByCategoryIds(categoryIds []uint) (err error) {
	err = r.DB.Model(&model.PerformanceTestPlan{}).
		Where("category_id IN (?)", categoryIds).
		Update("deleted", 1).Error

	return
}
