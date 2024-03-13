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
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func (r *ScriptRepo) Get(tenantId consts.TenantId, id uint) (script model.DebugConditionScript, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&script).Error
	return
}

func (r *ScriptRepo) GetByCondition(tenantId consts.TenantId, conditionId uint) (script model.DebugConditionScript, err error) {
	err = r.GetDB(tenantId).
		Where("condition_id=?", conditionId).
		Where("NOT deleted").
		First(&script).Error
	return
}

func (r *ScriptRepo) GetByName(tenantId consts.TenantId, name string, interfaceId uint) (script model.DebugConditionScript, err error) {
	var scripts []model.DebugConditionScript

	db := r.GetDB(tenantId).Model(&script).
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

func (r *ScriptRepo) Save(tenantId consts.TenantId, script *model.DebugConditionScript) (err error) {
	err = r.GetDB(tenantId).Save(script).Error
	return
}

func (r *ScriptRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugConditionScript{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}

func (r *ScriptRepo) DeleteByCondition(tenantId consts.TenantId, conditionId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugConditionScript{}).
		Where("condition_id=?", conditionId).
		Update("deleted", true).
		Error

	return
}

func (r *ScriptRepo) UpdateResult(tenantId consts.TenantId, script domain.ScriptBase) (err error) {
	values := map[string]interface{}{
		"output":        script.Output,
		"result_status": script.ResultStatus,
	}

	err = r.GetDB(tenantId).Model(&model.DebugConditionScript{}).
		Where("id=?", script.ConditionEntityId).
		Updates(values).
		Error

	return
}

func (r *ScriptRepo) CreateLog(tenantId consts.TenantId, script domain.ScriptBase) (
	log model.ExecLogScript, err error) {

	copier.CopyWithOption(&log, script, copier.Option{DeepCopy: true})

	log.ID = 0
	log.ConditionId = script.ConditionId
	log.ConditionEntityId = script.ConditionEntityId
	log.InvokeId = script.InvokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.GetDB(tenantId).Save(&log).Error

	return
}

func (r *ScriptRepo) CreateDefault(tenantId consts.TenantId, conditionId uint, src consts.ConditionSrc) (po model.DebugConditionScript) {
	po = model.DebugConditionScript{
		ScriptBase: domain.ScriptBase{
			ConditionId:  conditionId,
			Content:      scriptHelper.GetScript(scriptHelper.SnippetVariablesGet),
			ConditionSrc: src,
		},
	}

	r.Save(tenantId, &po)

	return
}

func (r *ScriptRepo) GetLog(tenantId consts.TenantId, conditionId, invokeId uint) (ret model.ExecLogScript, err error) {
	err = r.GetDB(tenantId).
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		First(&ret).Error

	ret.ConditionEntityType = consts.ConditionTypeScript

	return
}
