package repo

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"gorm.io/gorm"
)

type CheckpointRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *CheckpointRepo) List(interfaceId uint) (pos []model.InterfaceCheckpoint, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Order("created_at ASC").
		Find(&pos).Error
	return
}

func (r *CheckpointRepo) Get(id uint) (checkpoint model.InterfaceCheckpoint, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&checkpoint).Error
	return
}

func (r *CheckpointRepo) GetByName(name string, interfaceId uint) (checkpoint model.InterfaceCheckpoint, err error) {
	var checkpoints []model.InterfaceCheckpoint

	db := r.DB.Model(&checkpoint).
		Where("name = ? AND interface_id =? AND not deleted", name, interfaceId)

	err = db.Find(&checkpoints).Error

	if err != nil {
		return
	}

	if len(checkpoints) > 0 {
		checkpoint = checkpoints[0]
	}

	return
}

func (r *CheckpointRepo) Save(checkpoint *model.InterfaceCheckpoint) (err error) {
	err = r.DB.Save(checkpoint).Error
	return
}

func (r *CheckpointRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.InterfaceCheckpoint{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *CheckpointRepo) UpdateResult(checkpoint model.InterfaceCheckpoint) (err error) {
	err = r.DB.Model(&checkpoint).
		Where("id=?", checkpoint.ID).
		Update("result", checkpoint.ResultStatus).
		Error

	return
}
