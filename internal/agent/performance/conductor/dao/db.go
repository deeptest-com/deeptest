package dao

import (
	ptlog "github.com/aaronchen2k/deeptest/internal/agent/performance/pkg/log"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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
		ptlog.Logf("sqlite db conn obj = %v", DB)

		DB.AutoMigrate(Models...)
	})

	return DB
}

func GormSQLLite(app string) (db *gorm.DB) {
	pth := DBFile(app)
	ptlog.Logf("sqlite file is %s", pth)

	db, err := gorm.Open(sqlite.Open(pth), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		ptlog.Logf("sqlite db conn err %s", err.Error())
		return
	}

	return db
}

func DBFile(app string) string {
	path := filepath.Join(consts.WorkDir, app+".db")
	return path
}
