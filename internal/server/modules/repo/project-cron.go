package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ProjectCronRepo struct {
	DB           *gorm.DB      `inject:""`
	UserRepo     *UserRepo     `inject:""`
	CategoryRepo *CategoryRepo `inject:""`
	*BaseRepo    `inject:""`
}

func (r *ProjectCronRepo) Paginate(tenantId consts.TenantId, req v1.ProjectCronReqPaginate) (data _domain.PageData, err error) {
	baseSql := " FROM biz_project_cron t1 LEFT JOIN biz_project_cron_config_lecang t2 ON t1.config_id = t2.id LEFT JOIN biz_project_serve_swagger_sync t3 ON t1.config_id = t3.id WHERE NOT t1.deleted"

	if req.Name != "" {
		baseSql = baseSql + fmt.Sprintf("where t1.name LIKE %s", fmt.Sprintf("%%%s%%", req.Name))
	}
	if req.Source != "" {
		baseSql = baseSql + fmt.Sprintf("where t1.source = %s", req.Source)
	}

	if req.Switch != 0 {
		baseSql = baseSql + fmt.Sprintf("where t1.switch = %d", req.Switch)
	}

	selectSql := "SELECT t1.*,CASE WHEN t1.source = 'lecang' THEN t2.category_id WHEN t1.source = 'swagger' THEN t3.category_id ELSE 0 END AS category_id" + baseSql
	countSql := "SELECT count(*)" + baseSql

	var count int64

	db := r.GetDB(tenantId).Model(&model.ProjectCronList{})

	err = db.Raw(countSql).Count(&count).Error

	cronList := make([]*model.ProjectCronList, 0)
	err = db.Raw(selectSql + " limit " + strconv.Itoa((req.Page-1)*req.PageSize) + "," + strconv.Itoa(req.PageSize)).Find(&cronList).Error
	if err != nil {
		logUtils.Errorf("query project cron list error", zap.String("error:", err.Error()))
		return
	}

	r.CombineUserName(tenantId, cronList)
	r.CombineCategory(tenantId, cronList)
	data.Populate(cronList, count, req.Page, req.PageSize)

	return
}

func (r *ProjectCronRepo) CombineUserName(tenantId consts.TenantId, data []*model.ProjectCronList) {
	userIds := make([]uint, 0)
	for _, v := range data {
		userIds = append(userIds, v.CreateUserId)
	}
	userIds = _commUtils.ArrayRemoveUintDuplication(userIds)

	users, _ := r.UserRepo.FindByIds(tenantId, userIds)

	userIdNameMap := make(map[uint]string)
	for _, v := range users {
		userIdNameMap[v.ID] = v.Name
	}

	for _, v := range data {
		if name, ok := userIdNameMap[v.CreateUserId]; ok {
			v.CreateUserName = name
		}
	}
}

func (r *ProjectCronRepo) CombineCategory(tenantId consts.TenantId, configs []*model.ProjectCronList) {
	categoryIds := make([]int, 0)
	for _, v := range configs {
		categoryIds = append(categoryIds, v.CategoryId)
	}

	categoryIds = _commUtils.ArrayRemoveIntDuplication(categoryIds)
	categories, _ := r.CategoryRepo.BatchGetByIds(tenantId, categoryIds)

	categoryIdNameMap := make(map[int]string)
	for _, v := range categories {
		categoryIdNameMap[int(v.ID)] = v.Name
	}
	categoryIdNameMap[-1] = "未分类"

	for _, v := range configs {
		if name, ok := categoryIdNameMap[v.CategoryId]; ok {
			v.CategoryName = name
		}
	}
}

func (r *ProjectCronRepo) ListAllCron(tenantId consts.TenantId) (res []model.ProjectCron, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectCron{}).
		Where("switch = ? AND NOT deleted", consts.SwitchON).
		Find(&res).Error

	return
}
func (r *ProjectCronRepo) Create(tenantId consts.TenantId, config model.ProjectCron) (id uint, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectCron{}).Create(&config).Error
	if err != nil {
		return
	}

	id = config.ID
	return
}

func (r *ProjectCronRepo) Update(tenantId consts.TenantId, config model.ProjectCron) error {
	return r.GetDB(tenantId).Save(&config).Error
	//return r.DB.Model(&model.ProjectCron{}).Where("id = ?", config.ID).Updates(&config).Error
}

func (r *ProjectCronRepo) Save(tenantId consts.TenantId, config model.ProjectCron) (ret model.ProjectCron, err error) {
	err = r.GetDB(tenantId).Save(&config).Error
	if err != nil {
		return
	}

	ret = config

	return
}

func (r *ProjectCronRepo) DeleteById(tenantId consts.TenantId, id uint) error {
	return r.GetDB(tenantId).Model(&model.ProjectCron{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
}

func (r *ProjectCronRepo) UpdateSwitchById(tenantId consts.TenantId, id uint, switchStatus consts.SwitchStatus) error {
	return r.GetDB(tenantId).Model(&model.ProjectCron{}).
		Where("id = ?", id).
		Update("switch", switchStatus).Error
}

func (r *ProjectCronRepo) GetById(tenantId consts.TenantId, id uint) (config model.ProjectCron, err error) {
	err = r.GetDB(tenantId).Model(&model.ProjectCron{}).
		Where("id = ?", id).
		Find(&config).Error

	return
}

func (r *ProjectCronRepo) UpdateExecResult(tenantId consts.TenantId, configId uint, source consts.CronSource, execStatus consts.CronExecStatus, execErr string) (err error) {
	updateColumns := make(map[string]interface{})
	updateColumns["exec_status"] = execStatus
	updateColumns["exec_err"] = execErr
	updateColumns["exec_time"] = time.Now()

	err = r.GetDB(tenantId).Model(&model.ProjectCron{}).
		Where("config_id = ?", configId).
		Where("source = ?", source).
		Updates(updateColumns).Error

	return
}
