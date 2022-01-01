package repo

import (
	"encoding/json"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TestStepRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
}

func NewTestStepRepo() *TestStepRepo {
	return &TestStepRepo{}
}

func (r *TestStepRepo) ListByScript(scriptId uint) (steps []model.TestStep, err error) {
	err = r.DB.Model(&model.TestStep{}).Where("NOT deleted").
		Where("script_id = ?", scriptId).Find(&steps).Error

	if err != nil {
		logUtils.Errorf("query test step error", zap.String("error:", err.Error()))
		return
	}

	return
}

func (r *TestStepRepo) FindById(id uint) (step serverDomain.TestStepResp, err error) {
	err = r.DB.Model(&model.TestStep{}).Where("id = ?", id).First(&step).Error
	if err != nil {
		logUtils.Errorf("find step by id error", zap.String("error:", err.Error()))
		return
	}

	json.Unmarshal(step.CoordinatesJson, step.Coordinates)

	return
}

func (r *TestStepRepo) Create(req serverDomain.TestStepReq) (uint, error) {
	step := model.TestStep{}

	copier.Copy(&step, &req)
	bytes, _ := json.Marshal(req.Coordinates)
	step.CoordinatesJson = bytes

	err := r.DB.Model(&model.TestStep{}).Create(&step).Error
	if err != nil {
		logUtils.Errorf("add step error", zap.String("error:", err.Error()))
		return 0, err
	}

	return step.ID, nil
}

func (r *TestStepRepo) Update(id uint, req serverDomain.TestStepReq) error {
	step := req.TestStep
	err := r.DB.Model(&model.TestStep{}).Where("id = ?", id).Updates(&step).Error
	if err != nil {
		logUtils.Errorf("update step error", zap.String("error:", err.Error()))
		return err
	}

	return nil
}

func (r *TestStepRepo) DeleteById(id uint) (err error) {
	err = r.DB.Model(&model.TestStep{}).Where("id = ?", id).
		Updates(map[string]interface{}{"deleted": true}).Error
	if err != nil {
		logUtils.Errorf("delete step by id error", zap.String("error:", err.Error()))
		return
	}

	return
}
