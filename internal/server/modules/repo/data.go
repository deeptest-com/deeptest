package repo

import (
	"database/sql"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"gorm.io/gorm"
)

type DataRepo struct {
	*BaseRepo `inject:""`
	DB        *gorm.DB `inject:""`
}

func NewDataRepo(db *gorm.DB) *DataRepo {
	return &DataRepo{DB: db}
}

// CreateMySqlDb 创建数据库(mysql)
func (s *DataRepo) CreateMySqlDb(tenantId consts.TenantId) error {
	if tenantId != "" {
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/",
		config.CONFIG.Mysql.Username, config.CONFIG.Mysql.Password,
		config.CONFIG.Mysql.Url)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;",
		config.CONFIG.Mysql.Dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec(createSql)
	return err
}

func (s *DataRepo) SetSqlMode() (err error) {
	sql := "SET sql_mode = '';"
	err = s.DB.Raw(sql).Error

	return
}
