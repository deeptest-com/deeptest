package dao

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"path/filepath"
	"sync"
)

var (
	once sync.Once
	DB   *gorm.DB
)

func GetDB(app string) *gorm.DB {
	once.Do(func() {
		DB = GormSQLLite(app)

		DB.AutoMigrate(Models...)
	})

	return DB
}

func GormSQLLite(app string) (db *gorm.DB) {
	conn := DBFile(app)

	db, err := gorm.Open(sqlite.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logUtils.Info(err.Error())
		return
	}

	return db
}

func DBFile(app string) string {
	path := filepath.Join(consts.WorkDir, app+".db")
	return path
}
