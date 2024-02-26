package repo

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"

	"github.com/aaronchen2k/deeptest/internal/server/core/dao"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"gorm.io/gorm"
)

type IRepo interface {
	Save(tenantId consts.TenantId, id uint, entity interface{}) error
	GetCategoryCount(tenantId consts.TenantId, result interface{}, projectId uint) (err error)
	SaveEntity(tenantId consts.TenantId, category *model.Category) (err error)
}

type BaseRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *BaseRepo) GetAncestorIds(tenantId consts.TenantId, id uint, tableName string) (ids []uint, err error) {
	sql := `
		WITH RECURSIVE temp AS
		(
			SELECT id, parent_id, name from %s a where a.id = %d
		
			UNION ALL
		
			SELECT b.id, b.parent_id, b.name 
				from temp c
				inner join %s b on b.id = c.parent_id
		) 
		select id from temp e;
`

	sql = fmt.Sprintf(sql, tableName, id, tableName)

	err = r.GetDB(tenantId).Raw(sql).Scan(&ids).Error
	if err != nil {
		return
	}

	return
}

func (r *BaseRepo) GetDescendantIds(tenantId consts.TenantId, id uint, tableName string, typ serverConsts.CategoryDiscriminator, projectId int) (
	ids []uint, err error) {
	sql := `
		WITH RECURSIVE temp AS
		(
			SELECT id, parent_id from %s a 
				WHERE a.id = %d AND type='%s' AND project_id=%d AND NOT a.deleted
		
			UNION ALL
		
			SELECT b.id, b.parent_id 
				from temp c
				inner join %s b on b.parent_id = c.id
				WHERE type='%s' AND project_id=%d AND NOT b.deleted
		) 
		select id from temp e;
`
	sql = fmt.Sprintf(sql, tableName,
		id, typ, projectId,
		tableName,
		typ, projectId)

	err = r.GetDB(tenantId).Raw(sql).Scan(&ids).Error
	if err != nil {
		return
	}

	return
}

func (r *BaseRepo) GetAllChildIdsSimple(tenantId consts.TenantId, id uint, tableName string) (
	ids []uint, err error) {
	sql := `
		WITH RECURSIVE temp AS
		(
			SELECT id, parent_id from %s a 
				WHERE a.id = %d AND NOT a.deleted
		
			UNION ALL
		
			SELECT b.id, b.parent_id 
				from temp c
				inner join %s b on b.parent_id = c.id
				WHERE NOT b.deleted
		) 
		select id from temp e;
`
	sql = fmt.Sprintf(sql, tableName, id, tableName)

	err = r.GetDB(tenantId).Raw(sql).Scan(&ids).Error
	if err != nil {
		return
	}

	return
}

func (r *BaseRepo) Save(tenantId consts.TenantId, id uint, entity interface{}) (err error) {
	var count int64

	err = r.GetDB(tenantId).Model(&entity).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return
	}

	if count == 0 {
		err = r.GetDB(tenantId).Create(entity).Error
	} else {
		err = r.GetDB(tenantId).Updates(entity).Error
	}
	return
}

func (r *BaseRepo) SaveEntity(tenantId consts.TenantId, category *model.Category) (err error) {
	return
}

func (r *BaseRepo) GetAdminRoleName() (roleName consts.RoleType) {
	roleName = consts.Admin
	if config.CONFIG.System.SysEnv == "ly" {
		roleName = consts.IntegrationAdmin
	}
	return

}

func (r *BaseRepo) GetDB(tenantId consts.TenantId) (db *gorm.DB) {
	//return default db
	if tenantId == "" {
		return r.DB
	}
	handler := func() (db *gorm.DB, err error) {
		return dao.InitSaasDBHandler(tenantId)
	}
	db = dao.GetDBResolver().Apply(tenantId, handler).GetConnPool(tenantId)
	if db == nil {
		panic(fmt.Errorf("tenantId:%s,db is empty", tenantId))
	}
	return
}
