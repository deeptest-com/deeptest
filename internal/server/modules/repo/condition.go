package repo

import (
	"encoding/json"
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ConditionRepo struct {
	DB *gorm.DB `inject:""`

	ExtractorRepo    *ExtractorRepo    `inject:""`
	CheckpointRepo   *CheckpointRepo   `inject:""`
	ScriptRepo       *ScriptRepo       `inject:""`
	DatabaseOptRepo  *DatabaseOptRepo  `inject:""`
	DatabaseConnRepo *DatabaseConnRepo `inject:""`

	ResponseDefineRepo    *ResponseDefineRepo    `inject:""`
	EndpointInterfaceRepo *EndpointInterfaceRepo `inject:""`
}

func (r *ConditionRepo) List(debugInterfaceId, endpointInterfaceId uint, typ consts.ConditionCategory,
	usedBy consts.UsedBy, forAlternativeCase string, src consts.ConditionSrc) (
	pos []model.DebugCondition, err error) {

	db := r.DB.Where("NOT deleted")

	if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", endpointInterfaceId, 0)
	}

	if typ == consts.ConditionCategoryAssert {
		db.Where("entity_type = ?", consts.ConditionTypeCheckpoint)

	} else if typ == consts.PostCondition {
		db.Where("entity_type IN (?)", []consts.ConditionType{
			consts.ConditionTypeExtractor,
			consts.ConditionTypeScript,
			consts.ConditionTypeExtractor,
			consts.ConditionTypeDatabase,
		})

	} else if typ == consts.ConditionCategoryConsole {
		db.Where("entity_type IN (?)", []consts.ConditionType{
			consts.ConditionTypeExtractor,
			consts.ConditionTypeScript,
			consts.ConditionTypeExtractor,
			consts.ConditionTypeDatabase,
			consts.ConditionTypeCheckpoint,
		})

	} else if typ == consts.ConditionCategoryResponse {
		db.Where("entity_type = ?", consts.ConditionTypeResponseDefine)

	} else if typ == consts.ConditionCategoryResult {
		db.Where("entity_type IN (?) ",
			[]consts.ConditionType{
				consts.ConditionTypeResponseDefine,
				consts.ConditionTypeCheckpoint,
				consts.ConditionTypeScript,
			})
	}

	if usedBy != "" {
		db.Where("used_by=?", usedBy)
	}
	if src != "" {
		db.Where("condition_src=?", src)
	}

	if forAlternativeCase == "true" {
		db.Where("is_for_benchmark_case")
		db.Where("entity_type != ?", consts.ConditionTypeResponseDefine)

	} else if forAlternativeCase == "false" {
		db.Where("NOT is_for_benchmark_case")
	}

	db.Order("ordr ASC")

	err = db.Find(&pos).Error

	return
}

func (r *ConditionRepo) ListExtractor(req domain.DebugInfo) (
	pos []model.DebugCondition, err error) {

	db := r.DB.
		Where("NOT deleted").
		Order("ordr ASC")

	if req.DebugInterfaceId > 0 {
		db.Where("debug_interface_id=?", req.DebugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", req.EndpointInterfaceId, 0)
	}

	if req.UsedBy == consts.CaseDebug {
		db.Where("is_for_benchmark_case = ?", req.IsForBenchmarkCase)
	}

	db.Where("entity_type = ?", consts.ConditionTypeExtractor)

	err = db.Find(&pos).Error

	return
}

func (r *ConditionRepo) ListDbOpt(req domain.DebugInfo) (pos []model.DebugCondition, err error) {
	db := r.DB.
		Where("NOT deleted").
		Order("ordr ASC")

	if req.DebugInterfaceId > 0 {
		db.Where("debug_interface_id=?", req.DebugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", req.EndpointInterfaceId, 0)
	}

	if req.UsedBy == consts.CaseDebug {
		db.Where("is_for_benchmark_case = ?", req.IsForBenchmarkCase)
	}

	db.Where("entity_type = ?", consts.ConditionTypeDatabase)

	err = db.Find(&pos).Error

	return
}

func (r *ConditionRepo) Get(id uint) (po model.DebugCondition, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *ConditionRepo) Save(po *model.DebugCondition) (err error) {
	if po.Ordr == 0 {
		po.Ordr = r.GetMaxOrder(po.DebugInterfaceId, po.EndpointInterfaceId, po.IsForBenchmarkCase)
	}

	err = r.DB.Save(po).Error
	return
}

func (r *ConditionRepo) CloneAll(srcDebugInterfaceId, srcEndpointInterfaceId, distDebugInterfaceId uint,
	dictUsedBy, srcUsedBy consts.UsedBy, forAlternativeCase bool) (err error) {
	srcConditions, err := r.List(srcDebugInterfaceId, srcEndpointInterfaceId, consts.ConditionCategoryAll,
		srcUsedBy, fmt.Sprintf("%t", forAlternativeCase), "")

	for _, srcCondition := range srcConditions {
		// clone condition po
		srcCondition.ID = 0
		srcCondition.DebugInterfaceId = distDebugInterfaceId
		srcCondition.UsedBy = dictUsedBy
		srcCondition.IsForBenchmarkCase = false

		if srcDebugInterfaceId == distDebugInterfaceId { // clone to benchmark
			srcCondition.IsForBenchmarkCase = true
		}

		r.Save(&srcCondition)

		// clone condition entity
		var entityId uint
		if srcCondition.EntityType == consts.ConditionTypeScript {
			srcEntity, _ := r.ScriptRepo.Get(srcCondition.EntityId)
			srcEntity.ID = 0
			srcEntity.ConditionId = srcCondition.ID

			r.ScriptRepo.Save(&srcEntity)
			entityId = srcEntity.ID
		} else if srcCondition.EntityType == consts.ConditionTypeDatabase {
			srcEntity, _ := r.DatabaseOptRepo.Get(srcCondition.EntityId)
			srcEntity.ID = 0
			srcEntity.ConditionId = srcCondition.ID

			r.DatabaseOptRepo.Save(&srcEntity)
			entityId = srcEntity.ID

		} else if srcCondition.EntityType == consts.ConditionTypeExtractor {
			srcEntity, _ := r.ExtractorRepo.Get(srcCondition.EntityId)
			srcEntity.ID = 0
			srcEntity.ConditionId = srcCondition.ID

			r.ExtractorRepo.Save(&srcEntity)
			entityId = srcEntity.ID

		} else if srcCondition.EntityType == consts.ConditionTypeCheckpoint {
			srcEntity, _ := r.CheckpointRepo.Get(srcCondition.EntityId)
			srcEntity.ID = 0
			srcEntity.ConditionId = srcCondition.ID

			r.CheckpointRepo.Save(&srcEntity)
			entityId = srcEntity.ID

		} else if srcCondition.EntityType == consts.ConditionTypeResponseDefine {
			srcEntity, _ := r.ResponseDefineRepo.Get(srcCondition.EntityId)
			srcEntity.ID = 0
			srcEntity.ConditionId = srcCondition.ID
			srcEntity.Disabled = false

			r.ResponseDefineRepo.Save(&srcEntity)
			entityId = srcEntity.ID
		}

		err = r.UpdateEntityId(srcCondition.ID, entityId)
	}

	return
}

func (r *ConditionRepo) ReplaceAll(debugInterfaceId, endpointInterfaceId uint, conditions []domain.InterfaceExecCondition,
	usedBy consts.UsedBy, src consts.ConditionSrc) (err error) {

	r.removeAll(debugInterfaceId, endpointInterfaceId, usedBy, src)

	for _, item := range conditions {
		// clone condition po
		condition := model.DebugCondition{
			EntityType:          item.Type,
			DebugInterfaceId:    debugInterfaceId,
			EndpointInterfaceId: endpointInterfaceId,
			Desc:                item.Desc,
			ConditionSrc:        src,
		}
		r.Save(&condition)

		// clone condition entity
		var entityId uint
		if item.Type == consts.ConditionTypeExtractor {
			extractor := domain.ExtractorBase{}
			json.Unmarshal(item.Raw, &extractor)

			entity := model.DebugConditionExtractor{}

			copier.CopyWithOption(&entity, extractor, copier.Option{DeepCopy: true})
			entity.ID = 0
			entity.ConditionId = condition.ID

			r.ExtractorRepo.Save(&entity)
			entityId = entity.ID

		} else if item.Type == consts.ConditionTypeCheckpoint {
			checkpoint := domain.CheckpointBase{}
			json.Unmarshal(item.Raw, &checkpoint)

			entity := model.DebugConditionCheckpoint{}

			copier.CopyWithOption(&entity, checkpoint, copier.Option{DeepCopy: true})
			entity.ID = 0
			entity.ConditionId = condition.ID

			r.CheckpointRepo.Save(&entity)
			entityId = entity.ID

		} else if item.Type == consts.ConditionTypeScript {
			script := domain.ScriptBase{}
			json.Unmarshal(item.Raw, &script)

			entity := model.DebugConditionScript{}

			copier.CopyWithOption(&entity, script, copier.Option{DeepCopy: true})
			entity.ID = 0
			entity.ConditionId = condition.ID

			r.ScriptRepo.Save(&entity)
			entityId = entity.ID
		}

		err = r.UpdateEntityId(condition.ID, entityId)
	}

	return
}

func (r *ConditionRepo) Delete(id uint) (err error) {
	po, _ := r.Get(id)

	err = r.DB.Model(&model.DebugCondition{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	if po.EntityType == consts.ConditionTypeExtractor {
		r.ExtractorRepo.DeleteByCondition(id)
	} else if po.EntityType == consts.ConditionTypeCheckpoint {
		r.CheckpointRepo.DeleteByCondition(id)
	} else if po.EntityType == consts.ConditionTypeScript {
		r.ScriptRepo.DeleteByCondition(id)
	}

	return
}

func (r *ConditionRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.DebugCondition{}).
		Where("id=?", id).
		Update("disabled", gorm.Expr("NOT disabled")).
		Error

	return
}

func (r *ConditionRepo) UpdateOrders(req serverDomain.ConditionMoveReq) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for index, id := range req.Data {
			sql := fmt.Sprintf("UPDATE %s SET ordr = %d WHERE id = %d",
				model.DebugCondition{}.TableName(), index+1, id)

			err = r.DB.Exec(sql).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *ConditionRepo) UpdateEntityId(id uint, entityId uint) (err error) {
	err = r.DB.Model(&model.DebugCondition{}).
		Where("id=?", id).
		Update("entity_id", entityId).
		Error

	return
}

func (r *ConditionRepo) ListTo(debugInterfaceId, endpointInterfaceId uint,
	usedBy consts.UsedBy, forAlternativeCase string, src consts.ConditionSrc) (ret []domain.InterfaceExecCondition, err error) {
	pos, err := r.List(debugInterfaceId, endpointInterfaceId, consts.ConditionCategoryAll, usedBy, forAlternativeCase, src)

	for _, po := range pos {
		typ := po.EntityType

		if typ == consts.ConditionTypeScript {
			script := domain.ScriptBase{}

			entity, _ := r.ScriptRepo.Get(po.EntityId)
			copier.CopyWithOption(&script, entity, copier.Option{DeepCopy: true})
			script.Output = ""
			script.ConditionId = po.ID
			script.ConditionEntityId = po.EntityId
			script.ConditionEntityType = typ
			script.Disabled = po.Disabled

			raw, _ := json.Marshal(script)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
			}

			ret = append(ret, condition)

		} else if typ == consts.ConditionTypeDatabase {
			opt := domain.DatabaseOptBase{}

			entity, _ := r.DatabaseOptRepo.Get(po.EntityId)
			copier.CopyWithOption(&opt, entity, copier.Option{DeepCopy: true})

			opt.ConditionId = po.ID
			opt.ConditionEntityId = po.EntityId
			opt.ConditionEntityType = typ
			opt.Disabled = po.Disabled

			raw, _ := json.Marshal(opt)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
			}

			ret = append(ret, condition)

		} else if typ == consts.ConditionTypeExtractor {
			extractor := domain.ExtractorBase{}

			entity, _ := r.ExtractorRepo.Get(po.EntityId)
			copier.CopyWithOption(&extractor, entity, copier.Option{DeepCopy: true})
			extractor.ConditionEntityType = typ
			extractor.ConditionId = po.ID
			extractor.ConditionEntityId = po.EntityId
			extractor.Disabled = po.Disabled

			raw, _ := json.Marshal(extractor)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
				Desc: po.Desc,
			}

			ret = append(ret, condition)

		} else if typ == consts.ConditionTypeCheckpoint {
			checkpoint := domain.CheckpointBase{}

			entity, _ := r.CheckpointRepo.Get(po.EntityId)
			copier.CopyWithOption(&checkpoint, entity, copier.Option{DeepCopy: true})
			checkpoint.ConditionEntityType = typ
			checkpoint.ConditionId = po.ID
			checkpoint.ConditionEntityId = po.EntityId
			checkpoint.Disabled = po.Disabled

			raw, _ := json.Marshal(checkpoint)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
				Desc: po.Desc,
			}

			ret = append(ret, condition)

		} else if typ == consts.ConditionTypeResponseDefine {
			responseDefine := domain.ResponseDefineBase{}

			entity, err := r.ResponseDefineRepo.Get(po.EntityId)
			if err != nil {
				logUtils.Infof("响应码校验拿不到数据 %v", po.EntityId)
				continue
			}
			copier.CopyWithOption(&responseDefine, entity, copier.Option{DeepCopy: true})
			responseDefine.ConditionId = po.ID
			responseDefine.ConditionEntityId = po.EntityId
			responseDefine.ConditionEntityType = typ
			responseDefine.Disabled = po.Disabled

			responseBody := r.EndpointInterfaceRepo.GetResponse(endpointInterfaceId, entity.Code)
			if responseBody.ID == 0 {
				logUtils.Infof("响应体拿不到数据 %v", po.EntityId)
				continue
			}
			responseDefine.Schema = responseBody.SchemaItem.Content
			responseDefine.Code = entity.Code
			responseDefine.MediaType = responseBody.MediaType
			components := r.ResponseDefineRepo.Components(endpointInterfaceId)
			responseDefine.Component = commonUtils.JsonEncode(components)
			raw, _ := json.Marshal(responseDefine)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
			}

			ret = append(ret, condition)
		}

	}

	return
}

func (r *ConditionRepo) removeAll(debugInterfaceId, endpointInterfaceId uint, usedBy consts.UsedBy, src consts.ConditionSrc) (err error) {
	pos, _ := r.List(debugInterfaceId, endpointInterfaceId, "", usedBy, "false", src)

	for _, po := range pos {
		r.Delete(po.ID)
	}

	return
}

func (r *ConditionRepo) RemoveAllForBenchmarkCase(debugInterfaceId, endpointInterfaceId uint, usedBy consts.UsedBy, entityType consts.ConditionCategory) (err error) {
	pos, _ := r.List(debugInterfaceId, endpointInterfaceId, entityType, usedBy, "true", "")

	for _, po := range pos {
		if po.IsForBenchmarkCase {
			r.Delete(po.ID)
		}
	}

	return
}

func (r *ConditionRepo) CreateDefaultResponseDefine(debugInterfaceId, endpointInterfaceId uint, usedBy consts.UsedBy) (condition domain.Condition) {
	if endpointInterfaceId == 0 {
		return
	}

	codes := r.EndpointInterfaceRepo.GetResponseCodes(endpointInterfaceId)
	if len(codes) == 0 {
		return
	}

	po, err := r.GetByDebugInterfaceId(debugInterfaceId, endpointInterfaceId, usedBy)
	if err == gorm.ErrRecordNotFound {
		po, err = r.saveDefault(debugInterfaceId, endpointInterfaceId, codes, usedBy)
		if err != nil {
			return
		}
	}

	copier.CopyWithOption(&condition, po, copier.Option{DeepCopy: true})

	entityData, _ := r.ResponseDefineRepo.Get(po.EntityId)
	entityData.Codes = codes
	//entityData.Component = r.ResponseDefineRepo.Components(endpointInterfaceId)
	condition.EntityData = entityData

	return
}

func (r *ConditionRepo) GetByDebugInterfaceId(debugInterfaceId, endpointInterfaceId uint, by consts.UsedBy) (po model.DebugCondition, err error) {
	err = r.DB.
		Where("debug_interface_id=? and endpoint_interface_id=? and entity_type=?", debugInterfaceId, endpointInterfaceId, consts.ConditionTypeResponseDefine).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *ConditionRepo) saveDefault(debugInterfaceId, endpointInterfaceId uint, codes []string, by consts.UsedBy) (
	po model.DebugCondition, err error) {

	responseDefine := model.DebugConditionResponseDefine{}
	responseDefine.Code = "200"
	if len(codes) > 0 {
		responseDefine.Code = codes[0]
	}

	err = r.ResponseDefineRepo.Save(&responseDefine)
	if err != nil {
		return
	}

	po.EntityType = consts.ConditionTypeResponseDefine
	po.EndpointInterfaceId = endpointInterfaceId
	po.DebugInterfaceId = debugInterfaceId
	po.UsedBy = by
	po.EntityId = responseDefine.ID
	err = r.Save(&po)

	return
}

func (r *ConditionRepo) GetMaxOrder(debugInterfaceId, endpointInterfaceId uint, isForBenchmarkCase bool) (order int) {
	postCondition := model.DebugCondition{}

	db := r.DB.Model(&postCondition).
		Where("is_for_benchmark_case", isForBenchmarkCase)

	if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", endpointInterfaceId, 0)
	}

	err := db.Order("ordr DESC").
		First(&postCondition).Error

	if err == nil {
		order = postCondition.Ordr + 1
	}

	return
}
