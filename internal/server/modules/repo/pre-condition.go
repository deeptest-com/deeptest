package repo

import (
	"fmt"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
	"strings"
)

type PreConditionRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *PreConditionRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugPreCondition, err error) {
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
		Update("disabled", true).
		Error

	return
}

func (r *PreConditionRepo) UpdateOrders(req serverDomain.ConditionMoveReq) (err error) {
	var arr []string
	for index, id := range req.Data {
		str := fmt.Sprintf("UPDATE %s SET ordr = %d WHERE id = %d",
			model.DebugPreCondition{}.TableName(), index+1, id)
		arr = append(arr, str)
	}

	sql := strings.Join(arr, ";")

	err = r.DB.Exec(sql).Error

	return
}

//func (r *PreConditionRepo) UpdateOrder(node model.DebugPreCondition) (err error) {
//	err = r.DB.Model(&node).
//		Updates(model.DebugPreCondition{Ordr: node.Ordr}).
//		Error
//
//	return
//}
//
//func (r *PreConditionRepo) AddOrder(req serverDomain.ConditionMoveReq) (ret int, err error) {
//	dist, _ := r.Get(req.DropId)
//
//	db := r.DB.Model(&model.DebugPreCondition{}).
//		Where("NOT deleted")
//
//	if req.DebugInterfaceId > 0 {
//		db.Where("debug_interface_id = ?", req.DebugInterfaceId)
//
//	} else if req.EndpointInterfaceId > 0 {
//		db.Where("endpoint_interface_id = ?", req.EndpointInterfaceId)
//
//	}
//
//	if req.Position == serverConsts.Before {
//		db.Where("ordr >= ?", dist.Ordr)
//
//		ret = dist.Ordr
//
//	} else if req.Position == serverConsts.After {
//		r.DB.Model(&model.DiagnoseInterface{}).
//			Where("ordr > ?", dist.Ordr)
//
//		ret = dist.Ordr + 1
//
//	}
//
//	err = db.Update("ordr", gorm.Expr("ordr + 1")).Error
//
//	return
//}
