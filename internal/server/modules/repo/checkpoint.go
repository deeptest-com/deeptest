package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_i118Utils "github.com/aaronchen2k/deeptest/pkg/lib/i118"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CheckpointRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *CheckpointRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugConditionCheckpoint, err error) {
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

//func (r *CheckpointRepo) ListTo(debugInterfaceId, endpointInterfaceId uint) (ret []domain.CheckpointBase, err error) {
//	pos, err := r.List(debugInterfaceId, endpointInterfaceId)
//
//	for _, po := range pos {
//		checkpoint := domain.CheckpointBase{}
//		copier.CopyWithOption(&checkpoint, po, copier.Option{DeepCopy: true})
//
//		ret = append(ret, checkpoint)
//	}
//
//	return
//}

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
	name := ""

	opt := fmt.Sprintf("%v", po.Operator)
	optName := _i118Utils.Sprintf(opt)
	if po.Type == consts.ResponseStatus {
		name = _i118Utils.Sprintf("usage")
		name = fmt.Sprintf("状态码检查点 %s \"%s\"", optName, po.Value)
	} else if po.Type == consts.ResponseHeader {
		name = fmt.Sprintf("响应头检查点 %s \"%s\"", optName, po.Expression)
	} else if po.Type == consts.ResponseBody {
		name = fmt.Sprintf("响应体检查点 %s \"%s\"", optName, po.Value)
	} else if po.Type == consts.Extractor {
		name = fmt.Sprintf("提取器检查点 %s %s \"%s\"", po.ExtractorVariable, optName, po.Value)
	} else if po.Type == consts.Judgement {
		name = fmt.Sprintf("表达式检查点 \"%s\"", po.Expression)
	}

	desc := name
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

func (r *CheckpointRepo) UpdateResult(checkpoint model.DebugConditionCheckpoint, usedBy consts.UsedBy) (err error) {
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

func (r *CheckpointRepo) CreateLog(checkpoint model.DebugConditionCheckpoint, invokeId uint, usedBy consts.UsedBy) (
	log model.ExecLogCheckpoint, err error) {

	copier.CopyWithOption(&log, checkpoint, copier.Option{DeepCopy: true})

	log.ID = 0
	log.InvokeId = invokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.DB.Save(&log).Error

	return
}

//func (r *CheckpointRepo) UpdateResultToExecLog(checkpoint model.DebugConditionCheckpoint, log *model.ExecLogProcessor) (
//	logCheckpoint model.ExecLogCheckpoint, err error) {
//
//	copier.CopyWithOption(&logCheckpoint, checkpoint, copier.Option{DeepCopy: true})
//
//	logCheckpoint.ID = 0
//	logCheckpoint.InvokeId = log.ID
//	logCheckpoint.CreatedAt = nil
//	logCheckpoint.UpdatedAt = nil
//
//	err = r.DB.Save(&logCheckpoint).Error
//
//	return
//}

func (r *CheckpointRepo) CloneFromEndpointInterfaceToDebugInterface(endpointInterfaceId, debugInterfaceId uint,
	usedBy consts.UsedBy) (
	err error) {

	srcPos, _ := r.List(0, endpointInterfaceId)

	for _, po := range srcPos {
		po.ID = 0
		//po.EndpointInterfaceId = endpointInterfaceId
		//po.DebugInterfaceId = debugInterfaceId
		//po.UsedBy = usedBy

		r.Save(&po)
	}

	return
}

func (r *CheckpointRepo) CreateDefault(conditionId uint) (po model.DebugConditionCheckpoint) {
	po.ConditionId = conditionId

	po = model.DebugConditionCheckpoint{
		ConditionId: conditionId,

		CheckpointBase: domain.CheckpointBase{
			Type:              consts.ResponseStatus,
			Operator:          consts.Equal,
			Expression:        "",
			ExtractorVariable: "",
			Value:             "",
		},
	}

	r.Save(&po)

	return
}
