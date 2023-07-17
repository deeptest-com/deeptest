package repo

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type PostConditionRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *PostConditionRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugPostCondition, err error) {
	db := r.DB.
		Where("NOT deleted").
		Order("created_at ASC")

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

func (r *PostConditionRepo) Save(checkpoint *model.DebugPostCondition) (err error) {
	err = r.DB.Save(checkpoint).Error
	return
}

func (r *PostConditionRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *PostConditionRepo) Disable(id uint) (err error) {
	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", id).
		Update("disabled", gorm.Expr("NOT disabled")).
		Error

	return
}

func (r *PostConditionRepo) AddOrder(req serverDomain.ConditionMoveReq) (ret int, err error) {
	dist, _ := r.Get(req.DropId)

	db := r.DB.Model(&model.DebugPostCondition{}).
		Where("NOT deleted")

	if req.DebugInterfaceId > 0 {
		db.Where("debug_interface_id = ?", req.DebugInterfaceId)

	} else if req.EndpointInterfaceId > 0 {
		db.Where("endpoint_interface_id = ?", req.EndpointInterfaceId)

	}

	if req.Position == serverConsts.Before {
		db.Where("ordr >= ?", dist.Ordr)

		ret = dist.Ordr

	} else if req.Position == serverConsts.After {
		r.DB.Model(&model.DiagnoseInterface{}).
			Where("ordr > ?", dist.Ordr)

		ret = dist.Ordr + 1

	}

	err = db.Update("ordr", gorm.Expr("ordr + 1")).Error

	return
}

func (r *PostConditionRepo) UpdateOrder(node model.DebugPostCondition) (err error) {
	err = r.DB.Model(&node).
		Updates(model.DebugPostCondition{Ordr: node.Ordr}).
		Error

	return
}

func (r *PostConditionRepo) UpdateEntityId(id uint, entityId uint) (err error) {
	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", id).
		Update("entity_id", entityId).
		Error

	return
}
