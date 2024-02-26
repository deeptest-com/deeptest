package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type RelaPlanScenarioRepo struct {
	*BaseRepo `inject:""`
}

func (r *RelaPlanScenarioRepo) Get(tenantId consts.TenantId, id uint) (res model.RelaPlanScenario, err error) {
	err = r.GetDB(tenantId).Where("id = ?", id).First(&res).Error
	return
}

func (r *RelaPlanScenarioRepo) UpdateOrdrById(tenantId consts.TenantId, id uint, ordr int) (err error) {
	err = r.GetDB(tenantId).Model(model.RelaPlanScenario{}).Where("id=?", id).Update("ordr", ordr).Error
	return
}

func (r *RelaPlanScenarioRepo) IncreaseOrderAfter(tenantId consts.TenantId, ordr int, planId uint) (err error) {
	err = r.GetDB(tenantId).Model(model.RelaPlanScenario{}).Where("ordr >= ? and plan_id = ?  and not deleted", ordr, planId).UpdateColumn("ordr", gorm.Expr("ordr + ?", 1)).Error
	return
}

func (r *RelaPlanScenarioRepo) DecreaseOrderBefore(tenantId consts.TenantId, ordr int, planId uint) (err error) {
	err = r.GetDB(tenantId).Model(model.RelaPlanScenario{}).Where("ordr <= ? and plan_id = ?  and not deleted", ordr, planId).UpdateColumn("ordr", gorm.Expr("ordr - ?", 1)).Error
	return
}

func (r *RelaPlanScenarioRepo) GetMaxOrder(tenantId consts.TenantId, planId uint) (order int) {
	res := model.RelaPlanScenario{}

	err := r.GetDB(tenantId).Model(&model.RelaPlanScenario{}).
		Where("plan_id=? AND not deleted", planId).
		Order("ordr DESC").
		First(&res).Error

	if err == nil {
		order = res.Ordr + 1
	}

	return
}
