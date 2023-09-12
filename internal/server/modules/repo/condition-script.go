package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	scriptHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/script"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ScriptRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ScriptRepo) Get(id uint) (script model.DebugConditionScript, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&script).Error
	return
}

func (r *ScriptRepo) GetByCondition(conditionId uint) (script model.DebugConditionScript, err error) {
	err = r.DB.
		Where("condition_id=?", conditionId).
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

func (r *ScriptRepo) DeleteByCondition(conditionId uint) (err error) {
	err = r.DB.Model(&model.DebugConditionScript{}).
		Where("condition_id=?", conditionId).
		Update("deleted", true).
		Error

	return
}

func (r *ScriptRepo) UpdateResult(script domain.ScriptBase) (err error) {
	values := map[string]interface{}{
		"output":        script.Output,
		"result_status": script.ResultStatus,
	}

	err = r.DB.Model(&model.DebugConditionScript{}).
		Where("id=?", script.ConditionEntityId).
		Updates(values).
		Error

	return
}

func (r *ScriptRepo) CreateLog(script domain.ScriptBase) (
	log model.ExecLogScript, err error) {

	copier.CopyWithOption(&log, script, copier.Option{DeepCopy: true})

	log.ID = 0
	log.ConditionId = script.ConditionId
	log.ConditionEntityId = script.ConditionEntityId
	log.InvokeId = script.InvokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.DB.Save(&log).Error

	return
}

func (r *ScriptRepo) CreateDefault(conditionId uint, src consts.ConditionSrc) (po model.DebugConditionScript) {
	po = model.DebugConditionScript{
		ScriptBase: domain.ScriptBase{
			ConditionId:  conditionId,
			Content:      scriptHelper.GetScript(scriptHelper.SnippetVariablesGet),
			ConditionSrc: src,
		},
	}

	r.Save(&po)

	return
}

func (r *ScriptRepo) GetLog(conditionId, invokeId uint) (ret model.ExecLogScript, err error) {
	err = r.DB.
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		First(&ret).Error

	ret.ConditionEntityType = consts.ConditionTypeScript

	return
}
