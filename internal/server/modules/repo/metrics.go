package repo

import (
	"fmt"
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	model "github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MetricsRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *MetricsRepo) List(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (
	pos []model.AiMetrics, err error) {

	db := r.GetDB(tenantId).Where("NOT deleted")

	if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", endpointInterfaceId, 0)
	}

	db.Order("ordr ASC")

	err = db.Find(&pos).Error

	return
}

func (r *MetricsRepo) Get(tenantId consts.TenantId, id uint) (po model.AiMetrics, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *MetricsRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.AiMetrics{}).
		Where("id=?", id).
		Update("deleted", true).Error

	return
}

func (r *MetricsRepo) Disable(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.AiMetrics{}).
		Where("id=?", id).
		Update("disabled", gorm.Expr("NOT disabled")).
		Error

	return
}

func (r *MetricsRepo) UpdateOrders(tenantId consts.TenantId, req serverDomain.ConditionMoveReq) (err error) {
	return r.GetDB(tenantId).Transaction(func(tx *gorm.DB) error {
		for index, id := range req.Data {
			sql := fmt.Sprintf("UPDATE %s SET ordr = %d WHERE id = %d",
				model.AiMetrics{}.TableName(), index+1, id)

			err = r.GetDB(tenantId).Exec(sql).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *MetricsRepo) ListTo(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (
	ret []domain.AiMetrics, err error) {

	pos, err := r.List(tenantId, debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		to := domain.AiMetrics{}
		copier.CopyWithOption(&to, po, copier.Option{DeepCopy: true})

		ret = append(ret, to)
	}

	return
}

func (r *MetricsRepo) removeAll(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (err error) {
	pos, _ := r.List(tenantId, debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		r.Delete(tenantId, po.ID)
	}

	return
}

func (r *MetricsRepo) GetMaxOrder(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (order int) {
	postMetrics := model.AiMetrics{}

	db := r.GetDB(tenantId).Model(&postMetrics)

	if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", endpointInterfaceId, 0)
	}

	err := db.Order("ordr DESC").
		First(&postMetrics).Error

	if err == nil {
		order = postMetrics.Ordr + 1
	}

	return
}
