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

	ScriptRepo *ScriptRepo `inject:""`
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

	err = db.Find(&pos).Error

	return
}

func (r *PreConditionRepo) Get(id uint) (condition model.DebugPreCondition, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&condition).Error
	return
}

func (r *PreConditionRepo) Save(condition *model.DebugPreCondition) (err error) {
	err = r.DB.Save(condition).Error
	return
}

func (r *PreConditionRepo) CloneAll(srcDebugInterfaceId, srcEndpointInterfaceId, distDebugInterfaceId uint) (err error) {
	srcConditions, err := r.List(srcDebugInterfaceId, srcEndpointInterfaceId)

	for _, srcCondition := range srcConditions {
		// clone condition po
		srcCondition.ID = 0
		srcCondition.DebugInterfaceId = distDebugInterfaceId

		r.Save(&srcCondition)

		// clone condition entity
		var entityId uint
		if srcCondition.EntityType == consts.ConditionTypeScript {
			srcEntity, _ := r.ScriptRepo.Get(srcCondition.EntityId)
			srcEntity.ID = 0
			srcEntity.ConditionId = srcCondition.ID

			r.ScriptRepo.Save(&srcEntity)
			entityId = srcEntity.ID
		}

		err = r.UpdateEntityId(srcCondition.ID, entityId)
	}

	return
}

func (r *PreConditionRepo) ReplaceAll(debugInterfaceId, endpointInterfaceId uint, preConditions []domain.InterfaceExecCondition) (err error) {
	r.removeAll(debugInterfaceId, endpointInterfaceId)

	for _, item := range preConditions {
		// clone condition po
		condition := model.DebugPreCondition{
			EntityType:          consts.ConditionTypeScript,
			DebugInterfaceId:    debugInterfaceId,
			EndpointInterfaceId: endpointInterfaceId,
			Desc:                item.Desc,
		}
		r.Save(&condition)

		// clone condition entity
		var entityId uint
		if item.Type == consts.ConditionTypeScript {
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

func (r *PreConditionRepo) Delete(id uint) (err error) {
	po, _ := r.Get(id)

	err = r.DB.Model(&model.DebugPreCondition{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	if po.EntityType == consts.ConditionTypeScript {
		r.ScriptRepo.DeleteByCondition(id)
	}

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

		if typ == consts.ConditionTypeScript {
			script := domain.ScriptBase{}

			entity, _ := r.ScriptRepo.Get(po.EntityId)
			copier.CopyWithOption(&script, entity, copier.Option{DeepCopy: true})
			script.ConditionEntityType = typ
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

func (r *PreConditionRepo) removeAll(debugInterfaceId, endpointInterfaceId uint) (err error) {
	pos, _ := r.List(debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		r.Delete(po.ID)
	}

	return
}
