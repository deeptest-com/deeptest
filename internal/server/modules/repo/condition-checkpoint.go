package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	checkpointHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/checkpoint"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CheckpointRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *CheckpointRepo) Get(id uint) (checkpoint model.DebugConditionCheckpoint, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&checkpoint).Error
	return
}

func (r *CheckpointRepo) GetByName(name string, interfaceId uint) (checkpoint model.DebugConditionCheckpoint, err error) {
	var checkpoints []model.DebugConditionCheckpoint

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

func (r *CheckpointRepo) Save(checkpoint *model.DebugConditionCheckpoint) (err error) {
	r.UpdateDesc(checkpoint)

	err = r.DB.Save(checkpoint).Error
	return
}
func (r *CheckpointRepo) UpdateDesc(po *model.DebugConditionCheckpoint) (err error) {
	desc := checkpointHelpper.GenDesc(po.Type, po.Operator, po.Value, po.Expression, po.ExtractorVariable)
	values := map[string]interface{}{
		"desc": desc,
	}

	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", po.ConditionId).
		Updates(values).Error

	return
}

func (r *CheckpointRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugConditionCheckpoint{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}
func (r *CheckpointRepo) DeleteByCondition(conditionId uint) (err error) {
	err = r.DB.Model(&model.DebugConditionCheckpoint{}).
		Where("condition_id=?", conditionId).
		Update("deleted", true).
		Error

	return
}

func (r *CheckpointRepo) UpdateResult(checkpoint domain.CheckpointBase) (err error) {
	values := map[string]interface{}{
		"actual_result": checkpoint.ActualResult,
		"result_status": checkpoint.ResultStatus,
	}

	err = r.DB.Model(&model.DebugConditionCheckpoint{}).
		Where("id=?", checkpoint.ConditionEntityId).
		Updates(values).
		Error

	return
}

func (r *CheckpointRepo) CreateLog(checkpoint domain.CheckpointBase) (
	log model.ExecLogCheckpoint, err error) {

	copier.CopyWithOption(&log, checkpoint, copier.Option{DeepCopy: true})

	log.ID = 0
	log.ConditionId = checkpoint.ConditionId
	log.ConditionEntityId = checkpoint.ConditionEntityId

	log.InvokeId = checkpoint.InvokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.DB.Save(&log).Error

	return
}

func (r *CheckpointRepo) CreateDefault(conditionId uint) (po model.DebugConditionCheckpoint) {
	po.ConditionId = conditionId

	po = model.DebugConditionCheckpoint{
		CheckpointBase: domain.CheckpointBase{
			ConditionId: conditionId,

			Type:              consts.ResponseStatus,
			Operator:          consts.Equal,
			Expression:        "",
			ExtractorVariable: "",
			Value:             "200",
		},
	}

	r.Save(&po)

	return
}

func (r *CheckpointRepo) GetLog(conditionId, invokeId uint) (ret model.ExecLogCheckpoint, err error) {
	err = r.DB.
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		First(&ret).Error

	ret.ConditionEntityType = consts.ConditionTypeCheckpoint

	return
}

func (r *CheckpointRepo) GetLogFromScriptAssert(conditionId, invokeId uint) (ret []model.ExecLogCheckpoint, err error) {
	err = r.DB.
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		Find(&ret).Error

	for index, _ := range ret {
		ret[index].ConditionEntityType = consts.ConditionTypeCheckpoint
	}

	return
}
