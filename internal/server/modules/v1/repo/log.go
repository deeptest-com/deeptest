package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LogRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewLogRepo() *LogRepo {
	return &LogRepo{}
}

func (r *LogRepo) ListByReport(reportId uint) (logs []*model.ExecLogProcessor, err error) {
	err = r.DB.
		Where("report_id=?", reportId).
		Where("NOT deleted").
		Order("parent_id ASC, id ASC").
		Find(&logs).Error
	return
}

func (r *LogRepo) Get(id uint) (scenario model.ExecLogProcessor, err error) {
	err = r.DB.Model(&model.ExecLogProcessor{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *LogRepo) Save(log *model.ExecLogProcessor) (err error) {
	err = r.DB.Save(log).Error

	return
}

func (r *LogRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.ExecLogProcessor{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
