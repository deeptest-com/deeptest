package repo

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ShareVariableRepo struct {
	DB        *gorm.DB `inject:""`
	*BaseRepo `inject:""`

	ScenarioProcessorRepo *ScenarioProcessorRepo `inject:""`
}

func (r *ShareVariableRepo) Save(tenantId consts.TenantId, po *model.ShareVariable) (err error) {
	po.ID, _ = r.findExist(tenantId, *po)

	err = r.GetDB(tenantId).Save(po).Error

	return
}

func (r *ShareVariableRepo) findExist(tenantId consts.TenantId, po model.ShareVariable) (id uint, err error) {
	existPo := model.ShareVariable{}

	db := r.GetDB(tenantId).Model(&po).
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

func (r *ShareVariableRepo) GetExistByInterfaceDebug(tenantId consts.TenantId, name string, serveId uint, usedBy consts.UsedBy) (id uint, err error) {
	po := model.ShareVariable{}

	err = r.GetDB(tenantId).Model(&po).
		Where("name = ? AND used_by = ? AND serve_id =? AND not deleted",
			name, usedBy, serveId).
		First(&po).Error

	id = po.ID

	return
}
func (r *ShareVariableRepo) GetExistByScenarioDebug(tenantId consts.TenantId, name string, scenarioId uint) (id uint, err error) {
	po := model.ShareVariable{}

	err = r.GetDB(tenantId).Model(&po).
		Where("name = ? AND scenario_id =? AND not deleted",
			name, scenarioId).
		First(&po).Error

	id = po.ID

	return
}

func (r *ShareVariableRepo) ListForInterfaceDebug(tenantId consts.TenantId, serveId uint, usedBy consts.UsedBy) (pos []model.ShareVariable, err error) {
	err = r.GetDB(tenantId).Model(&model.ShareVariable{}).
		Where("serve_id=?", serveId).
		Where("used_by=?", usedBy).
		Where("NOT deleted AND NOT disabled").
		Find(&pos).Error

	return
}

func (r *ShareVariableRepo) ListForScenarioDebug(tenantId consts.TenantId, processorId uint) (pos []model.ShareVariable, err error) {
	processor, _ := r.ScenarioProcessorRepo.Get(tenantId, processorId)
	scenarioId := processor.ScenarioId

	ancestorProcessorIds, err := r.GetAncestorIds(tenantId, processorId, model.Processor{}.TableName())

	err = r.GetDB(tenantId).Model(&model.ShareVariable{}).
		Where("scenario_processor_id IN ?", ancestorProcessorIds).
		Where("scenario_id=?", scenarioId).
		Where("NOT deleted AND NOT disabled").
		Find(&pos).Error

	return
}

func (r *ShareVariableRepo) Delete(tenantId consts.TenantId, id int) (err error) {
	err = r.GetDB(tenantId).Model(&model.ShareVariable{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ShareVariableRepo) DeleteAllByServeId(tenantId consts.TenantId, serveId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ShareVariable{}).
		Where("serve_id=?", serveId).
		Update("deleted", true).
		Error

	return
}
func (r *ShareVariableRepo) DeleteAllByScenarioId(tenantId consts.TenantId, scenarioId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.ShareVariable{}).
		Where("scenario_id=?", scenarioId).
		Update("disabled", true).
		Error

	return
}
