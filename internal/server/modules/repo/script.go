package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ScriptRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ScriptRepo) List(debugInterfaceId, endpointInterfaceId uint) (pos []model.DebugConditionScript, err error) {
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

//func (r *ScriptRepo) ListTo(debugInterfaceId, endpointInterfaceId uint) (ret []domain.ScriptBase, err error) {
//	pos, err := r.List(debugInterfaceId, endpointInterfaceId)
//
//	for _, po := range pos {
//		script := domain.ScriptBase{}
//		copier.CopyWithOption(&script, po, copier.Option{DeepCopy: true})
//
//		ret = append(ret, script)
//	}
//
//	return
//}

func (r *ScriptRepo) Get(id uint) (script model.DebugConditionScript, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&script).Error
	return
}

func (r *ScriptRepo) GetByName(name string, interfaceId uint) (script model.DebugConditionScript, err error) {
	var scripts []model.DebugConditionScript

	db := r.DB.Model(&script).
		Where("name = ? AND endpoint_interface_id =? AND not deleted", name, interfaceId)

	err = db.Find(&scripts).Error

	if err != nil {
		return
	}

	if len(scripts) > 0 {
		script = scripts[0]
	}

	return
}

func (r *ScriptRepo) Save(script *model.DebugConditionScript) (err error) {
	err = r.DB.Save(script).Error
	return
}

func (r *ScriptRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugConditionScript{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ScriptRepo) UpdateResult(script model.DebugConditionScript, usedBy consts.UsedBy) (err error) {
	values := map[string]interface{}{
		"output":        script.Output,
		"result_status": script.ResultStatus,
	}

	err = r.DB.Model(&script).
		Where("id=? AND used_by=?", script.ID, usedBy).
		Updates(values).
		Error

	return
}

func (r *ScriptRepo) CreateLog(script model.DebugConditionScript, invokeId uint, usedBy consts.UsedBy) (
	log model.ExecLogScript, err error) {

	copier.CopyWithOption(&log, script, copier.Option{DeepCopy: true})

	log.ID = 0
	log.InvokeId = invokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.DB.Save(&log).Error

	return
}

//func (r *ScriptRepo) UpdateResultToExecLog(script model.DebugConditionScript, log *model.ExecLogProcessor) (
//	logScript model.ExecLogCheckpoint, err error) {
//
//	copier.CopyWithOption(&logScript, script, copier.Option{DeepCopy: true})
//
//	logScript.ID = 0
//	logScript.LogId = log.ID
//	logScript.CreatedAt = nil
//	logScript.UpdatedAt = nil
//
//	err = r.DB.Save(&logScript).Error
//
//	return
//}

func (r *ScriptRepo) CloneFromEndpointInterfaceToDebugInterface(endpointInterfaceId, debugInterfaceId uint,
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

func (r *ScriptRepo) CreateDefault(conditionId uint, src consts.ConditionSrc) (po model.DebugConditionScript) {
	po = model.DebugConditionScript{
		ConditionId:  conditionId,
		Content:      scriptHelper.GetScript(scriptHelper.ScriptVariablesGet),
		ConditionSrc: src,
	}

	r.Save(&po)

	return
}
