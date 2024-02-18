package repo

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	databaseOptHelpper "github.com/aaronchen2k/deeptest/internal/pkg/helper/database-opt"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type DatabaseOptRepo struct {
	*BaseRepo        `inject:""`
	DB               *gorm.DB          `inject:""`
	DatabaseConnRepo *DatabaseConnRepo `inject:""`
}

func (r *DatabaseOptRepo) Get(tenantId consts.TenantId, id uint) (databaseOpt model.DebugConditionDatabaseOpt, err error) {
	err = r.GetDB(tenantId).
		Where("id=?", id).
		Where("NOT deleted").
		First(&databaseOpt).Error
	if err != nil {
		return
	}

	dbConn, err := r.DatabaseConnRepo.Get(tenantId, databaseOpt.DbConnId)
	if err != nil || dbConn.Disabled {
		databaseOpt.DatabaseConnIsDisabled = true
		err = nil

		databaseOpt.DbConnId = 0
		databaseOpt.DatabaseConnBase = domain.DatabaseConnBase{} // clear

	} else {
		databaseOpt.Type = dbConn.Type

		databaseOpt.Host = dbConn.Host
		databaseOpt.Port = dbConn.Port
		databaseOpt.DbName = dbConn.DbName
		databaseOpt.Username = dbConn.Username
		databaseOpt.Password = dbConn.Password
	}

	return
}

func (r *DatabaseOptRepo) Save(tenantId consts.TenantId, databaseOpt *model.DebugConditionDatabaseOpt) (err error) {
	conn, _ := r.DatabaseConnRepo.Get(tenantId, databaseOpt.DbConnId)

	if conn.ID > 0 {
		databaseOpt.Type = conn.Type
		databaseOpt.Host = conn.Host
		databaseOpt.Port = conn.Port
		databaseOpt.Username = conn.Username
		databaseOpt.Password = conn.Password
		databaseOpt.DbName = conn.DbName
	}

	err = r.GetDB(tenantId).Save(databaseOpt).Error

	r.UpdateDesc(tenantId, databaseOpt)

	return
}
func (r *DatabaseOptRepo) UpdateDesc(tenantId consts.TenantId, po *model.DebugConditionDatabaseOpt) (err error) {
	desc := databaseOptHelpper.GenDesc(po.Type, po.Sql)
	values := map[string]interface{}{
		"desc": desc,
	}

	err = r.GetDB(tenantId).Model(&model.DebugCondition{}).
		Where("id=?", po.ConditionId).
		Updates(values).Error

	return
}

func (r *DatabaseOptRepo) Delete(tenantId consts.TenantId, id uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugConditionDatabaseOpt{}).
		Where("id=?", id).
		Update("deleted", true).
		Error

	return
}
func (r *DatabaseOptRepo) DeleteByCondition(tenantId consts.TenantId, conditionId uint) (err error) {
	err = r.GetDB(tenantId).Model(&model.DebugConditionDatabaseOpt{}).
		Where("condition_id=?", conditionId).
		Update("deleted", true).
		Error

	return
}

func (r *DatabaseOptRepo) UpdateResult(tenantId consts.TenantId, databaseOpt domain.DatabaseOptBase) (err error) {
	values := map[string]interface{}{
		"result_msg":    databaseOpt.ResultMsg,
		"result_status": databaseOpt.ResultStatus,
	}

	err = r.GetDB(tenantId).Model(&model.DebugConditionDatabaseOpt{}).
		Where("id=?", databaseOpt.ConditionEntityId).
		Updates(values).
		Error

	return
}

func (r *DatabaseOptRepo) CreateLog(tenantId consts.TenantId, databaseOpt domain.DatabaseOptBase) (
	log model.ExecLogDatabaseOpt, err error) {

	copier.CopyWithOption(&log, databaseOpt, copier.Option{DeepCopy: true})

	log.ID = 0
	log.ConditionId = databaseOpt.ConditionId
	log.ConditionEntityId = databaseOpt.ConditionEntityId

	log.InvokeId = databaseOpt.InvokeId
	log.CreatedAt = nil
	log.UpdatedAt = nil

	err = r.GetDB(tenantId).Save(&log).Error

	return
}

func (r *DatabaseOptRepo) CreateDefault(tenantId consts.TenantId, conditionId uint, src consts.ConditionSrc) (po model.DebugConditionDatabaseOpt) {
	po.ConditionId = conditionId

	po = model.DebugConditionDatabaseOpt{
		DatabaseOptBase: domain.DatabaseOptBase{
			ConditionId: conditionId,

			DatabaseConnBase: domain.DatabaseConnBase{
				Type: consts.DbTypeMySql,
			},
			ConditionSrc: src,
		},
	}

	r.Save(tenantId, &po)

	return
}

func (r *DatabaseOptRepo) GetLog(tenantId consts.TenantId, conditionId, invokeId uint) (ret model.ExecLogDatabaseOpt, err error) {
	err = r.GetDB(tenantId).
		Where("condition_id=? AND invoke_id=?", conditionId, invokeId).
		Where("NOT deleted").
		First(&ret).Error

	ret.ConditionEntityType = consts.ConditionTypeDatabase

	return
}
