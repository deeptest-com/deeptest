package repo

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type ProjectCronRepo struct {
	DB       *gorm.DB  `inject:""`
	UserRepo *UserRepo `inject:""`
}

func (r *ProjectCronRepo) Paginate(req v1.ProjectCronReqPaginate) (data _domain.PageData, err error) {
	var count int64

	db := r.DB.Model(&model.ProjectCron{}).
		Where("project_id = ? AND NOT deleted",
			req.ProjectId)

	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}
	if req.Source != "" {
		db = db.Where("source = ?", req.Source)
	}
	if req.Switch != 0 {
		db = db.Where("switch = ?", req.Switch)
	}

	err = db.Count(&count).Error
	if err != nil {
		logUtils.Errorf("count scenario error", zap.String("error:", err.Error()))
		return
	}

	cronList := make([]*model.ProjectCron, 0)

	err = db.
		Scopes(dao.PaginateScope(req.Page, req.PageSize, req.Order, req.Field)).
		Find(&cronList).Error
	if err != nil {
		logUtils.Errorf("query project cron list error", zap.String("error:", err.Error()))
		return
	}

	r.CombineUserName(cronList)
	r.CombineCategory(cronList)
	data.Populate(cronList, count, req.Page, req.PageSize)

	return
}

func (r *ProjectCronRepo) CombineUserName(data []*model.ProjectCron) {
	userIds := make([]uint, 0)
	for _, v := range data {
		userIds = append(userIds, v.CreateUserId)
	}
	userIds = _commUtils.ArrayRemoveUintDuplication(userIds)

	users, _ := r.UserRepo.FindByIds(userIds)

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

func (r *ProjectCronRepo) CombineCategory(configs []*model.ProjectCron) {
	// TODO

}

func (r *ProjectCronRepo) ListAllCron() (res []model.ProjectCron, err error) {
	err = r.DB.Model(&model.ProjectCron{}).
		Where("switch = ? AND NOT deleted", consts.SwitchON).
		Find(&res).Error

	return
}
func (r *ProjectCronRepo) Create(config model.ProjectCron) (id uint, err error) {
	err = r.DB.Model(&model.ProjectCron{}).Create(&config).Error
	if err != nil {
		return
	}

	id = config.ID
	return
}

func (r *ProjectCronRepo) Update(config model.ProjectCron) error {
	return r.DB.Save(&config).Error
	//return r.DB.Model(&model.ProjectCron{}).Where("id = ?", config.ID).Updates(&config).Error
}

func (r *ProjectCronRepo) Save(config model.ProjectCron) (id uint, err error) {
	err = r.DB.Save(&config).Error
	if err != nil {
		return
	}

	id = config.ID

	return
}

func (r *ProjectCronRepo) DeleteById(id uint) error {
	return r.DB.Model(&model.ProjectCron{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
}

func (r *ProjectCronRepo) UpdateSwitchById(id uint, switchStatus consts.SwitchStatus) error {
	return r.DB.Model(&model.ProjectCron{}).
		Where("id = ?", id).
		Update("switch", switchStatus).Error
}

func (r *ProjectCronRepo) GetById(id uint) (config model.ProjectCron, err error) {
	err = r.DB.Model(&model.ProjectCron{}).
		Where("id = ?", id).
		Find(&config).Error

	return
}

func (r *ProjectCronRepo) UpdateExecResult(configId uint, source consts.CronSource, execStatus consts.CronExecStatus, execErr string) (err error) {
	updateColumns := make(map[string]interface{})
	updateColumns["exec_status"] = execStatus
	updateColumns["exec_err"] = execErr
	updateColumns["exec_time"] = time.Now()

	err = r.DB.Model(&model.ProjectCron{}).
		Where("config_id = ?", configId).
		Where("source = ?", source).
		Updates(updateColumns).Error

	return
}
