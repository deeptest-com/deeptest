package repo

import (
	"fmt"
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *BaseRepo) GetAllParentIds(id uint, tableName string) (ids []uint, err error) {
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

	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		return
	}

	return
}

func (r *BaseRepo) GetAllChildIds(id uint, tableName string) (ids []uint, err error) {
	sql := `
		WITH RECURSIVE temp AS
		(
			SELECT id, parent_id, name from %s a where a.id = %d AND NOT a.deleted
		
			UNION ALL
		
			SELECT b.id, b.parent_id, b.name 
				from temp c
				inner join %s b on b.parent_id = c.id
				WHERE NOT b.deleted
		) 
		select id from temp e;
`

	sql = fmt.Sprintf(sql, tableName, id, tableName)

	err = r.DB.Raw(sql).Scan(&ids).Error
	if err != nil {
		return
	}

	return
}

func (r *BaseRepo) Save(id uint, entity interface{}) (err error) {
	var count int64
	r.DB.Where("id = ?", id).Count(&count)
	if count == 0 {
		err = r.DB.Create(entity).Error
	} else {
		err = r.DB.Updates(entity).Error
	}
	return
}
