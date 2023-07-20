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

type PreConditionRepo struct {
	DB *gorm.DB `inject:""`

	ExtractorRepo  *ExtractorRepo  `inject:""`
	CheckpointRepo *CheckpointRepo `inject:""`
	ScriptRepo     *ScriptRepo     `inject:""`
}

func (r *PreConditionRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugPreCondition, err error) {
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

func (r *PreConditionRepo) Get(id uint) (checkpoint model.DebugPreCondition, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&checkpoint).Error
	return
}

func (r *PreConditionRepo) Save(checkpoint *model.DebugPreCondition) (err error) {
	err = r.DB.Save(checkpoint).Error
	return
}

func (r *PreConditionRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugPreCondition{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *PreConditionRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.DebugPreCondition{}).
		Where("id=?", id).
		Update("disabled", gorm.Expr("NOT disabled")).
		Error

	return
}

func (r *PreConditionRepo) UpdateOrders(req serverDomain.ConditionMoveReq) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for index, id := range req.Data {
			sql := fmt.Sprintf("UPDATE %s SET ordr = %d WHERE id = %d",
				model.DebugPreCondition{}.TableName(), index+1, id)

			err = r.DB.Exec(sql).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *PreConditionRepo) UpdateEntityId(id uint, entityId uint) (err error) {
	err = r.DB.Model(&model.DebugPreCondition{}).
		Where("id=?", id).
		Update("entity_id", entityId).
		Error

	return
}

func (r *PreConditionRepo) ListTo(debugInterfaceId, endpointInterfaceId uint) (ret []domain.InterfaceExecCondition, err error) {
	pos, err := r.List(debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		typ := po.EntityType

		if typ == consts.ConditionTypeExtractor {
			extractor := domain.ExtractorBase{}

			entity, _ := r.ExtractorRepo.Get(po.EntityId)
			copier.CopyWithOption(&extractor, entity, copier.Option{DeepCopy: true})

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
