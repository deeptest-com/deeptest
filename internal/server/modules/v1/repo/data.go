package repo

import (
	"database/sql"
	"fmt"
	serverConfig "github.com/aaronchen2k/deeptest/internal/server/config"
	"gorm.io/gorm"
)

type DataRepo struct {
	DB *gorm.DB `inject:""`
}

func NewDataRepo() *DataRepo {
	return &DataRepo{}
}

// CreateMySqlDb 创建数据库(mysql)
func (s *DataRepo) CreateMySqlDb() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/",
		serverConfig.CONFIG.Mysql.Username, serverConfig.CONFIG.Mysql.Password,
		serverConfig.CONFIG.Mysql.Url)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;",
		serverConfig.CONFIG.Mysql.Dbname)

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
