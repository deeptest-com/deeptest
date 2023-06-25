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

func (r *CheckpointRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugInterfaceCheckpoint, err error) {
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

func (r *CheckpointRepo) ListTo(debugInterfaceId, endpointInterfaceId uint) (ret []agentDomain.Checkpoint, err error) {
	pos, err := r.List(debugInterfaceId, endpointInterfaceId)

	for _, po := range pos {
		checkpoint := agentDomain.Checkpoint{}
		copier.CopyWithOption(&checkpoint, po, copier.Option{DeepCopy: true})

		ret = append(ret, checkpoint)
	}

	return
}

func (r *CheckpointRepo) Get(id uint) (checkpoint model.DebugInterfaceCheckpoint, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&checkpoint).Error
	return
}

func (r *CheckpointRepo) GetByName(name string, interfaceId uint) (checkpoint model.DebugInterfaceCheckpoint, err error) {
	var checkpoints []model.DebugInterfaceCheckpoint

	db := r.DB.Model(&checkpoint).
		Where("name = ? AND endpoint_interface_id =? AND not deleted", name, interfaceId)

	err = db.Find(&checkpoints).Error

	if err != nil {
		return
	}

	if len(checkpoints) > 0 {
		checkpoint = checkpoints[0]
	}

	return
}

func (r *CheckpointRepo) Save(checkpoint *model.DebugInterfaceCheckpoint) (err error) {
	err = r.DB.Save(checkpoint).Error
	return
}

func (r *CheckpointRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugInterfaceCheckpoint{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *CheckpointRepo) UpdateResult(checkpoint model.DebugInterfaceCheckpoint, usedBy consts.UsedBy) (err error) {
	values := map[string]interface{}{
		"actual_result": checkpoint.ActualResult,
		"result_status": checkpoint.ResultStatus,
	}

	err = r.DB.Model(&checkpoint).
		Where("id=? AND used_by=?", checkpoint.ID, usedBy).
		Updates(values).
		Error

	return
}

func (r *CheckpointRepo) UpdateResultToExecLog(checkpoint model.DebugInterfaceCheckpoint, log *model.ExecLogProcessor) (
	logCheckpoint model.ExecLogCheckpoint, err error) {

	copier.CopyWithOption(&logCheckpoint, checkpoint, copier.Option{DeepCopy: true})

	logCheckpoint.ID = 0
	logCheckpoint.LogId = log.ID
	logCheckpoint.CreatedAt = nil
	logCheckpoint.UpdatedAt = nil

	err = r.DB.Save(&logCheckpoint).Error

	return
}

func (r *CheckpointRepo) CloneFromEndpointInterfaceToDebugInterface(endpointInterfaceId, debugInterfaceId uint,
	usedBy consts.UsedBy) (
	err error) {

	srcPos, _ := r.List(0, endpointInterfaceId)

	for _, po := range srcPos {
		po.ID = 0
		po.EndpointInterfaceId = endpointInterfaceId
		po.DebugInterfaceId = debugInterfaceId
		po.UsedBy = usedBy

		r.Save(&po)
	}

	return
}
