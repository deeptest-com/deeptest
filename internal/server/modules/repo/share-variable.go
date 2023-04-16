package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type ShareVariableRepo struct {
	DB       *gorm.DB  `inject:""`
	RoleRepo *RoleRepo `inject:""`
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

	err = r.DB.Model(&po).
		Where("name=?, interfaceId=?, serveId=?, scenarioId=?, scope=?",
			po.Name, po.InterfaceId, po.ServeId, po.ScenarioId, po.Scope).
		Where("NOT deleted AND NOT disabled").
		First(&existPo).Error

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

func (r *ShareVariableRepo) ListByScenarioDebug(scenarioId uint) (pos []model.ShareVariable, err error) {
	err = r.DB.Model(&model.ShareVariable{}).
		Where("scenario_id=?", scenarioId).
		Where("NOT deleted AND NOT disabled").
		Find(&pos).Error

	return

	return
}
