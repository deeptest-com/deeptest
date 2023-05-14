package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ShareVariableRepo struct {
	DB        *gorm.DB `inject:""`
	*BaseRepo `inject:""`

	ScenarioProcessorRepo *ScenarioProcessorRepo `inject:""`
}

func NewShareVariableRepo() *ShareVariableRepo {
	return &ShareVariableRepo{}
}

func (r *ShareVariableRepo) Save(po *model.ShareVariable) (err error) {
	po.ID, _ = r.findExist(*po)

	err = r.DB.Save(po).Error

	return
}

func (r *ShareVariableRepo) findExist(po model.ShareVariable) (id uint, err error) {
	existPo := model.ShareVariable{}

	db := r.DB.Model(&po).
		Where("name=?", po.Name).
		Where("NOT deleted AND NOT disabled")

	if po.ServeId > 0 {
		db.Where("serve_id=?", po.ServeId)
	}

	if po.ScenarioId > 0 {
		db.Where("scenario_id=?", po.ScenarioId)
	}

	err = db.First(&existPo).Error

	id = po.ID

	return
}

func (r *ShareVariableRepo) GetExistByInterfaceDebug(name string, serveId uint) (id uint, err error) {
	po := model.ShareVariable{}

	err = r.DB.Model(&po).
		Where("name = ? AND serve_id =? AND not deleted",
			name, serveId).
		First(&po).Error

	id = po.ID

	return
}
func (r *ShareVariableRepo) GetExistByScenarioDebug(name string, scenarioId uint) (id uint, err error) {
	po := model.ShareVariable{}

	err = r.DB.Model(&po).
		Where("name = ? AND scenario_id =? AND not deleted",
			name, scenarioId).
		First(&po).Error

	id = po.ID

	return
}

func (r *ShareVariableRepo) ListByInterfaceDebug(serveId uint) (pos []model.ShareVariable, err error) {
	err = r.DB.Model(&model.ShareVariable{}).
		Where("serve_id=?", serveId).
		Where("NOT deleted AND NOT disabled").
		Find(&pos).Error

	return
}

func (r *ShareVariableRepo) ListByScenarioDebug(processorId uint) (pos []model.ShareVariable, err error) {
	processor, _ := r.ScenarioProcessorRepo.Get(processorId)
	scenarioId := processor.ScenarioId

	parentIds, err := r.GetAllParentIds(processorId, model.Processor{}.TableName())

	err = r.DB.Model(&model.ShareVariable{}).
		Where("processor_id IN ?", parentIds).
		Where("scenario_id=?", scenarioId).
		Where("NOT deleted AND NOT disabled").
		Find(&pos).Error

	return
}

func (r *ShareVariableRepo) Delete(id int) (err error) {
	err = r.DB.Model(&model.ShareVariable{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ShareVariableRepo) DeleteAllByServeId(serveId uint) (err error) {
	err = r.DB.Model(&model.ShareVariable{}).
		Where("serve_id=?", serveId).
		Update("deleted", true).
		Error

	return
}
func (r *ShareVariableRepo) DeleteAllByScenarioId(scenarioId uint) (err error) {
	err = r.DB.Model(&model.DebugInterfaceExtractor{}).
		Where("scenario_id=?", scenarioId).
		Update("disable_share", true).
		Error

	return
}
