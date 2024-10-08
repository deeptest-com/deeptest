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

		switch typ {
		case consts.Summarization:
			to := domain.AiMetricsSummarizationBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.AnswerRelevancy:
			to := domain.AiMetricsAnswerRelevancyBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.Faithfulness:
			to := domain.AiMetricsFaithfulnessBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.ContextualPrecision:
			to := domain.AiMetricsContextualPrecisionBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.ContextualRecall:
			to := domain.AiMetricsContextualRecallBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.ContextualRelevancy:
			to := domain.AiMetricsContextualRelevancyBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.Hallucination:
			to := domain.AiMetricsHallucinationBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.Bias:
			to := domain.AiMetricsBiasBase{}
			item, _ := r.CopyEntity(to, po, tenantId)

			ret = append(ret, item)

		case consts.Toxicity:
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
	switch typ {
	case consts.Summarization:
		po := model.AiMetricsSummarization{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.AnswerRelevancy:
		po := model.AiMetricsAnswerRelevancy{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.Faithfulness:
		po := model.AiMetricsFaithfulness{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.ContextualPrecision:
		po := model.AiMetricsContextualPrecision{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.ContextualRecall:
		po := model.AiMetricsContextualRecall{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.ContextualRelevancy:
		po := model.AiMetricsContextualRelevancy{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.Hallucination:
		po := model.AiMetricsHallucination{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.Bias:
		po := model.AiMetricsBias{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po

	case consts.Toxicity:
		po := model.AiMetricsToxicity{}
		err = r.GetDB(tenantId).Where("id = ?", id).First(&po).Error
		entity = po
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

func (r *MetricsRepo) CreateDefault(tenantId consts.TenantId, metricsId uint, typ consts.MetricsType) (
	entityId uint, err error) {

	switch typ {
	case consts.Summarization:
		entity := model.AiMetricsSummarization{
			AiMetricsSummarizationBase: domain.AiMetricsSummarizationBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "摘要",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.AnswerRelevancy:
		entity := model.AiMetricsAnswerRelevancy{
			AiMetricsAnswerRelevancyBase: domain.AiMetricsAnswerRelevancyBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "回答相关性",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.Faithfulness:
		entity := model.AiMetricsFaithfulness{
			AiMetricsFaithfulnessBase: domain.AiMetricsFaithfulnessBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "忠实度",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.ContextualPrecision:
		entity := model.AiMetricsContextualPrecision{
			AiMetricsContextualPrecisionBase: domain.AiMetricsContextualPrecisionBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "检索查准率",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.ContextualRecall:
		entity := model.AiMetricsContextualRecall{
			AiMetricsContextualRecallBase: domain.AiMetricsContextualRecallBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "检索查全率",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.ContextualRelevancy:
		entity := model.AiMetricsContextualRelevancy{
			AiMetricsContextualRelevancyBase: domain.AiMetricsContextualRelevancyBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "检索相关性",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.Hallucination:
		entity := model.AiMetricsHallucination{
			AiMetricsHallucinationBase: domain.AiMetricsHallucinationBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "幻觉",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.Bias:
		entity := model.AiMetricsBias{
			AiMetricsBiasBase: domain.AiMetricsBiasBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "歧视",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID

	case consts.Toxicity:
		entity := model.AiMetricsToxicity{
			AiMetricsToxicityBase: domain.AiMetricsToxicityBase{
				AiMetricsEntityBase: domain.AiMetricsEntityBase{
					Name:      "毒性",
					MetricsId: metricsId,
				},
			},
		}
		err = r.GetDB(tenantId).Save(&entity).Error
		entityId = entity.ID
	}

	return
}
