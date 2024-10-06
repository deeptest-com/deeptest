package repo

import (
	"encoding/json"
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

func (r *MetricsRepo) Save(tenantId consts.TenantId, po *model.AiMetrics) (err error) {
	if po.Ordr == 0 {
		po.Ordr = r.GetMaxOrder(tenantId, po.DebugInterfaceId, po.EndpointInterfaceId)
	}

	err = r.GetDB(tenantId).Save(po).Error
	return
}

func (r *MetricsRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	po, _ := r.Get(tenantId, id)

	err = r.GetDB(tenantId).Model(&model.AiMetrics{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	deleteDb := r.GetDB(tenantId)
	if po.EntityType == consts.Summarization {
		deleteDb.Model(&model.AiMetricsSummarization{})
	} else if po.EntityType == consts.AnswerRelevancy {
		deleteDb.Model(&model.AiMetricsAnswerRelevancy{})
	} else if po.EntityType == consts.Faithfulness {
		deleteDb.Model(&model.AiMetricsFaithfulness{})
	} else if po.EntityType == consts.ContextualPrecision {
		deleteDb.Model(&model.AiMetricsContextualPrecision{})
	} else if po.EntityType == consts.ContextualRecall {
		deleteDb.Model(&model.AiMetricsContextualRecall{})
	} else if po.EntityType == consts.ContextualRelevancy {
		deleteDb.Model(&model.AiMetricsContextualRelevancy{})
	} else if po.EntityType == consts.Hallucination {
		deleteDb.Model(&model.AiMetricsHallucination{})
	} else if po.EntityType == consts.Bias {
		deleteDb.Model(&model.AiMetricsBias{})
	} else if po.EntityType == consts.Toxicity {
		deleteDb.Model(&model.AiMetricsToxicity{})
	}

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

func (r *MetricsRepo) UpdateEntityId(tenantId consts.TenantId, id uint, entityId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.AiMetrics{}).
		Where("id=?", id).
		Update("entity_id", entityId).
		Error

	return
}

func (r *MetricsRepo) ListTo(tenantId consts.TenantId, debugInterfaceId, endpointInterfaceId uint) (
	ret []domain.InterfaceExecMetrics, err error) {
	pos, err := r.List(tenantId, debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		typ := po.EntityType

		if typ == consts.Summarization {
			to := domain.AiMetricsSummarizationBase{}

			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.AnswerRelevancy {
			to := domain.AiMetricsAnswerRelevancyBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.Faithfulness {
			to := domain.AiMetricsFaithfulnessBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.ContextualPrecision {
			to := domain.AiMetricsContextualPrecisionBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.ContextualRecall {
			to := domain.AiMetricsContextualRecallBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.ContextualRelevancy {
			to := domain.AiMetricsContextualRelevancyBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.Hallucination {
			to := domain.AiMetricsHallucinationBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.Bias {
			to := domain.AiMetricsBiasBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		} else if typ == consts.Toxicity {
			to := domain.AiMetricsToxicityBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)
		}
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

func (r *MetricsRepo) GetEntity(tenantId consts.TenantId, id uint, typ consts.MetricsType) (entity interface{}, err error) {
	if typ == consts.Summarization {
		entity = model.AiMetricsSummarization{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.AnswerRelevancy {
		entity = model.AiMetricsAnswerRelevancy{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.Faithfulness {
		entity = model.AiMetricsFaithfulness{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.ContextualPrecision {
		entity = model.AiMetricsContextualPrecision{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.ContextualRecall {
		entity = model.AiMetricsContextualRecall{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.ContextualRelevancy {
		entity = model.AiMetricsContextualRelevancy{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.Hallucination {
		entity = model.AiMetricsHallucination{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.Bias {
		entity = model.AiMetricsBias{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error

	} else if typ == consts.Toxicity {
		entity = model.AiMetricsToxicity{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&entity).Error
	}

	return
}

func (r *MetricsRepo) CopyEntity(to domain.EntityToInterface, po model.AiMetrics, tenantId consts.TenantId) (
	ret domain.InterfaceExecMetrics, err error) {

	entity, _ := r.GetEntity(tenantId, po.EntityId, po.EntityType)

	copier.CopyWithOption(&to, entity, copier.Option{DeepCopy: true})

	to.SetInfo(po.ID, po.EntityId, po.EntityType, po.Disabled)

	raw, _ := json.Marshal(to)
	ret = domain.InterfaceExecMetrics{
		Type: po.EntityType,
		Raw:  raw,
	}

	return
}
