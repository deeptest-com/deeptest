package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	extractorHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/extractor"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type ExtractorRepo struct {
	DB                *gorm.DB           `inject:""`
	PostConditionRepo *PostConditionRepo `inject:""`
}

func (r *ExtractorRepo) Get(id uint) (extractor model.DebugConditionExtractor, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&extractor).Error
	return
}

func (r *ExtractorRepo) GetByInterfaceVariable(variable string, id, debugInterfaceId uint) (extractor model.DebugConditionExtractor, err error) {
	db := r.DB.Model(&extractor).
		Where("variable = ? AND debug_interface_id =? AND not deleted",
			variable, debugInterfaceId)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&extractor)

	return
}

func (r *ExtractorRepo) Save(extractor *model.DebugConditionExtractor) (id uint, err error) {
	err = r.DB.Save(extractor).Error
	if err != nil {
		return
	}

	id = extractor.ID

	return
}

func (r *ExtractorRepo) Update(extractor *model.DebugConditionExtractor) (err error) {
	r.UpdateDesc(extractor)

	err = r.DB.Updates(extractor).Error
	if err != nil {
		return
	}

	return
}

func (r *ExtractorRepo) UpdateDesc(po *model.DebugConditionExtractor) (err error) {
	desc := extractorHelper.GenDesc(po.Variable, po.Src, po.Key, po.Type, po.Expression, po.BoundaryStart, po.BoundaryEnd)
	values := map[string]interface{}{
		"desc": desc,
	}

	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", po.ConditionId).
		Updates(values).Error

	return
}

func (r *ExtractorRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugConditionExtractor{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}
func (r *ExtractorRepo) DeleteByCondition(conditionId uint) (err error) {
	err = r.DB.Model(&model.DebugConditionExtractor{}).
		Where("condition_id=?", conditionId).
		Update("deleted", true).
		Error

	return
}

func (r *ExtractorRepo) ListLogByInvoke(invokeId uint) (pos []model.ExecLogExtractor, err error) {
	err = r.DB.
		Where("NOT deleted").
		Where("invoke_id=?", invokeId).
		Order("created_at ASC").Error

	return
}

func (r *ExtractorRepo) UpdateResult(extractor domain.ExtractorBase) (err error) {
	extractor.Result = strings.TrimSpace(extractor.Result)
	values := map[string]interface{}{}
	if extractor.Result != "" {
		values["result"] = extractor.Result
		values["result_type"] = extractor.ResultType
	}
	if extractor.Scope != "" {
		values["scope"] = extractor.Scope
	}

	err = r.DB.Model(&model.DebugConditionExtractor{}).
		Where("id = ?", extractor.ConditionEntityId).
		Updates(values).Error

	if err != nil {
		logUtils.Errorf("update DebugConditionExtractor error", zap.String("error:", err.Error()))
		return err
	}

	return
}

func (r *ExtractorRepo) CreateLog(extractor domain.ExtractorBase) (
	log model.ExecLogExtractor, err error) {

	copier.CopyWithOption(&log, extractor, copier.Option{DeepCopy: true})

	log.ID = 0
	log.ConditionId = extractor.ConditionId
	log.ConditionEntityId = extractor.ConditionEntityId
	log.InvokeId = extractor.InvokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.DB.Save(&log).Error

	return
}

func (r *ExtractorRepo) ListExtractorVariableByConditions(conditionIds []uint) (ret []domain.Variable, err error) {
	err = r.DB.Model(&model.DebugConditionExtractor{}).
		Select("id, variable AS name, result AS value").
		Where("condition_id IN (?)", conditionIds).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&ret).Error

	return
}
func (r *ExtractorRepo) ListDbOptVariableByConditions(conditionIds []uint) (ret []domain.Variable, err error) {
	err = r.DB.Model(&model.DebugConditionDatabaseOpt{}).
		Select("id, variable AS name, result AS value").
		Where("condition_id IN (?)", conditionIds).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&ret).Error

	return
}

func (r *ExtractorRepo) GetParentIds(processorId uint, ids *[]uint) {
	var po model.Processor

	r.DB.Where("id = ?", processorId).
		Where("NOT deleted AND NOT disabled").
		First(&po)

	if po.ID > 0 {
		*ids = append(*ids, processorId)
	}

	if po.ParentId > 0 {
		r.GetParentIds(po.ParentId, ids)
	}

	return
}

func (r *ExtractorRepo) CreateDefault(conditionId uint) (po model.DebugConditionExtractor) {
	po = model.DebugConditionExtractor{
		ExtractorBase: domain.ExtractorBase{
			ConditionId: conditionId,

			Src:           consts.Body,
			Type:          consts.Boundary,
			Expression:    "",
			Variable:      "x",
			BoundaryStart: "_prefix",
			BoundaryEnd:   "_postfix",
			Scope:         consts.Public,
		},
	}

	_, err := r.Save(&po)
	if err != nil {
		return
	}

	return
}

func (r *ExtractorRepo) GetLog(conditionId, invokeId uint) (ret model.ExecLogExtractor, err error) {
	err = r.DB.
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		First(&ret).Error

	ret.ConditionEntityType = consts.ConditionTypeExtractor

	return
}
