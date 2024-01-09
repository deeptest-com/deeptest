package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type RelaPlanScenarioRepo struct {
	*BaseRepo `inject:""`
}

func (r *RelaPlanScenarioRepo) Get(id uint) (res model.RelaPlanScenario, err error) {
	err = r.DB.Where("id = ?", id).First(&res).Error
	return
}

func (r *RelaPlanScenarioRepo) UpdateOrdrById(id uint, ordr int) (err error) {
	err = r.DB.Update("ordr", ordr).Where("id=?", id).Error
	return
}

func (r *RelaPlanScenarioRepo) IncreaseOrderAfter(id, planId uint) (err error) {
	err = r.DB.UpdateColumn("ordr", gorm.Expr("ordr + ?", 1)).Where("id >= ? and plan_id = ?  and not deleted", id, planId).Error
	return
}
