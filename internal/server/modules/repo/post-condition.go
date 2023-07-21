package repo

import (
	"encoding/json"
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type PostConditionRepo struct {
	DB *gorm.DB `inject:""`

	ExtractorRepo  *ExtractorRepo  `inject:""`
	CheckpointRepo *CheckpointRepo `inject:""`
	ScriptRepo     *ScriptRepo     `inject:""`
}

func (r *PostConditionRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugPostCondition, err error) {
	db := r.DB.
		Where("NOT deleted").
		Order("ordr ASC")

	if debugInterfaceId > 0 {
		db.Where("debug_interface_id=?", debugInterfaceId)
	} else {
		db.Where("endpoint_interface_id=? AND debug_interface_id=?", endpointInterfaceId, 0)
	}

	err = db.
		Find(&pos).Error

	return
}

func (r *PostConditionRepo) Get(id uint) (po model.DebugPostCondition, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&po).Error
	return
}

func (r *PostConditionRepo) Save(po *model.DebugPostCondition) (err error) {
	err = r.DB.Save(po).Error
	return
}

func (r *PostConditionRepo) Delete(id uint) (err error) {
	po, _ := r.Get(id)

	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	if po.EntityType == consts.ConditionTypeExtractor {
		r.ScriptRepo.DeleteByCondition(id)
	} else if po.EntityType == consts.ConditionTypeCheckpoint {
		r.CheckpointRepo.DeleteByCondition(id)
	} else if po.EntityType == consts.ConditionTypeScript {
		r.ScriptRepo.DeleteByCondition(id)
	}

	return
}

func (r *PostConditionRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", id).
		Update("disabled", gorm.Expr("NOT disabled")).
		Error

	return
}

func (r *PostConditionRepo) UpdateOrders(req serverDomain.ConditionMoveReq) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for index, id := range req.Data {
			sql := fmt.Sprintf("UPDATE %s SET ordr = %d WHERE id = %d",
				model.DebugPostCondition{}.TableName(), index+1, id)

			err = r.DB.Exec(sql).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *PostConditionRepo) UpdateEntityId(id uint, entityId uint) (err error) {
	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", id).
		Update("entity_id", entityId).
		Error

	return
}

func (r *PostConditionRepo) ListTo(debugInterfaceId, endpointInterfaceId uint) (ret []domain.InterfaceExecCondition, err error) {
	pos, err := r.List(debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		typ := po.EntityType

		if typ == consts.ConditionTypeExtractor {
			extractor := domain.ExtractorBase{}

			entity, _ := r.ExtractorRepo.Get(po.EntityId)
			copier.CopyWithOption(&extractor, entity, copier.Option{DeepCopy: true})
			extractor.ConditionId = po.ID
			extractor.ConditionEntityId = po.EntityId

			raw, _ := json.Marshal(extractor)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
			}

			ret = append(ret, condition)

		} else if typ == consts.ConditionTypeCheckpoint {
			checkpoint := domain.CheckpointBase{}

			entity, _ := r.CheckpointRepo.Get(po.EntityId)
			copier.CopyWithOption(&checkpoint, entity, copier.Option{DeepCopy: true})
			checkpoint.ConditionId = po.ID
			checkpoint.ConditionEntityId = po.EntityId

			raw, _ := json.Marshal(checkpoint)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
			}

			ret = append(ret, condition)

		} else if typ == consts.ConditionTypeScript {
			script := domain.ScriptBase{}

			entity, _ := r.ScriptRepo.Get(po.EntityId)
			copier.CopyWithOption(&script, entity, copier.Option{DeepCopy: true})
			script.ConditionId = po.ID
			script.ConditionEntityId = po.EntityId

			raw, _ := json.Marshal(script)
			condition := domain.InterfaceExecCondition{
				Type: typ,
				Raw:  raw,
			}

			ret = append(ret, condition)
		}

	}

	return
}
