package repo

import (
	"github.com/aaronchen2k/deeptest/internal/agent/exec/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CheckpointRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *CheckpointRepo) List(interfaceId uint, usedBy consts.UsedBy) (pos []model.InterfaceCheckpoint, err error) {
	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("used_by = ? AND NOT deleted", usedBy).
		Order("created_at ASC").
		Find(&pos).Error
	return
}

func (r *CheckpointRepo) ListTo(interfaceId uint) (ret []domain.Checkpoint, err error) {
	pos := make([]model.InterfaceCheckpoint, 0)

	err = r.DB.
		Where("interface_id=?", interfaceId).
		Where("NOT deleted").
		Find(&pos).Error

	for _, po := range pos {
		checkpoint := domain.Checkpoint{}
		copier.CopyWithOption(&checkpoint, po, copier.Option{DeepCopy: true})

		ret = append(ret, checkpoint)
	}

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

func (r *CheckpointRepo) UpdateResult(checkpoint model.InterfaceCheckpoint, usedBy consts.UsedBy) (err error) {
	values := map[string]interface{}{
		"actual_result": checkpoint.ActualResult,
		"result_status": checkpoint.ResultStatus,
		"used_by":       usedBy,
	}

	err = r.DB.Model(&checkpoint).
		Where("id=?", checkpoint.ID).
		Updates(values).
		Error

	return
}

func (r *CheckpointRepo) UpdateResultToExecLog(checkpoint model.InterfaceCheckpoint, log *model.ExecLogProcessor) (
	logCheckpoint model.ExecLogCheckpoint, err error) {

	copier.CopyWithOption(&logCheckpoint, checkpoint, copier.Option{DeepCopy: true})
	logCheckpoint.ID = 0
	logCheckpoint.LogId = log.ID
	logCheckpoint.CreatedAt = nil
	logCheckpoint.UpdatedAt = nil

	err = r.DB.Save(&logCheckpoint).Error

	return
}
