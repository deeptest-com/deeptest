package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TestLogRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewTestLogRepo() *TestLogRepo {
	return &TestLogRepo{}
}

func (r *TestLogRepo) Get(id uint) (scenario model.Log, err error) {
	err = r.DB.Model(&model.Log{}).Where("id = ?", id).First(&scenario).Error
	if err != nil {
		logUtils.Errorf("find scenario by id error", zap.String("error:", err.Error()))
		return scenario, err
	}

	return scenario, nil
}

func (r *TestLogRepo) Save(log *model.Log) (err error) {
	err = r.DB.Save(log).Error

	return
}

func (r *TestLogRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.Log{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete scenario by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
