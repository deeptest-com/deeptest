package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	cookieHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/cookie"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type CookieRepo struct {
	DB                *gorm.DB           `inject:""`
	PostConditionRepo *PostConditionRepo `inject:""`
}

func (r *CookieRepo) Get(id uint) (cookie model.DebugConditionCookie, err error) {
	err = r.DB.
		Where("id=?", id).
		Where("NOT deleted").
		First(&cookie).Error
	return
}

func (r *CookieRepo) GetByInterfaceVariable(variable string, id, debugInterfaceId uint) (cookie model.DebugConditionCookie, err error) {
	db := r.DB.Model(&cookie).
		Where("variable = ? AND debug_interface_id =? AND not deleted",
			variable, debugInterfaceId)

	if id > 0 {
		db.Where("id != ?", id)
	}

	db.First(&cookie)

	return
}

func (r *CookieRepo) Save(cookie *model.DebugConditionCookie) (id uint, err error) {
	err = r.DB.Save(cookie).Error
	if err != nil {
		return
	}

	id = cookie.ID

	return
}

func (r *CookieRepo) Update(cookie *model.DebugConditionCookie) (err error) {
	r.UpdateDesc(cookie)

	err = r.DB.Updates(cookie).Error
	if err != nil {
		return
	}

	return
}

func (r *CookieRepo) UpdateDesc(po *model.DebugConditionCookie) (err error) {
	desc := cookieHelper.GenDesc(po.CookieName, po.VariableName)
	values := map[string]interface{}{
		"desc": desc,
	}

	err = r.DB.Model(&model.DebugPostCondition{}).
		Where("id=?", po.ConditionId).
		Updates(values).Error

	return
}

func (r *CookieRepo) Delete(id uint) (err error) {
	err = r.DB.Model(&model.DebugConditionCookie{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}
func (r *CookieRepo) DeleteByCondition(conditionId uint) (err error) {
	err = r.DB.Model(&model.DebugConditionCookie{}).
		Where("condition_id=?", conditionId).
		Update("deleted", true).
		Error

	return
}

func (r *CookieRepo) ListLogByInvoke(invokeId uint) (pos []model.ExecLogCookie, err error) {
	err = r.DB.
		Where("NOT deleted").
		Where("invoke_id=?", invokeId).
		Order("created_at ASC").Error

	return
}

func (r *CookieRepo) UpdateResult(cookie domain.CookieBase) (err error) {
	cookie.Result = strings.TrimSpace(cookie.Result)
	values := map[string]interface{}{}
	if cookie.Result != "" {
		values["result"] = cookie.Result
	}

	err = r.DB.Model(&model.DebugConditionCookie{}).
		Where("id = ?", cookie.ConditionEntityId).
		Updates(values).Error

	if err != nil {
		logUtils.Errorf("update DebugConditionCookie error", zap.String("error:", err.Error()))
		return err
	}

	return
}

func (r *CookieRepo) CreateLog(cookie domain.CookieBase) (
	log model.ExecLogCookie, err error) {

	copier.CopyWithOption(&log, cookie, copier.Option{DeepCopy: true})

	log.ID = 0
	log.ConditionId = cookie.ConditionId
	log.ConditionEntityId = cookie.ConditionEntityId
	log.InvokeId = cookie.InvokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.DB.Save(&log).Error

	return
}

func (r *CookieRepo) ListCookieVariableByInterface(conditionIds []uint) (variables []domain.Variable, err error) {
	err = r.DB.Model(&model.DebugConditionCookie{}).
		Select("id, variable AS name, result AS value").
		Where("condition_id IN (?)", conditionIds).
		Where("NOT deleted AND NOT disabled").
		Order("created_at ASC").
		Find(&variables).Error

	return
}

func (r *CookieRepo) ListValidCookieVariableForInterface(interfaceId, projectId uint, usedBy consts.UsedBy) (
	variables []domain.Variable, err error) {

	q := r.DB.Model(&model.DebugConditionCookie{}).
		Select("id, variable AS name, result AS value, " +
			"endpoint_interface_id AS endpointInterfaceId, scope AS scope").
		Where("NOT deleted AND NOT disabled")

	//if usedBy == consts.InterfaceDebug {
	//	q.Where("project_id=?", projectId)
	//
	//} else {
	//	processorInterface, _ := r.ProcessorInterfaceRepo.Get(interfaceId)
	//
	//	var parentIds []uint
	//	r.GetParentIds(processorInterface.ProcessorId, &parentIds)
	//
	//	q.Where("scenario_id=?", processorInterface.ScenarioId).
	//		Where("scope = ? OR scenario_processor_id IN(?)", consts.Public, parentIds)
	//}

	err = q.Order("created_at ASC").
		Find(&variables).Error

	return
}

func (r *CookieRepo) GetParentIds(processorId uint, ids *[]uint) {
	var po model.Processor

	r.DB.Where("id = ?", processorId).
		Where("NOT deleted AND NOT disabled").
		First(&po)

	if po.ID > 0 {
		*ids = append(*ids, processorId)
	}

	if po.ParentId > 0 {
		r.GetParentIds(po.ParentId, ids)
	}

	return
}

func (r *CookieRepo) CreateDefault(conditionId uint) (po model.DebugConditionCookie) {
	po = model.DebugConditionCookie{
		CookieBase: domain.CookieBase{
			ConditionId: conditionId,

			CookieName:   "cookie_name",
			VariableName: "variable_name",
		},
	}

	r.Save(&po)

	return
}

func (r *CookieRepo) GetLog(conditionId, invokeId uint) (ret model.ExecLogCookie, err error) {
	err = r.DB.
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		First(&ret).Error

	ret.ConditionEntityType = consts.ConditionTypeCookie

	return
}
