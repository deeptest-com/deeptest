package dao

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
)

// GetDB 数据库单例
func GetDB(args ...interface{}) *gorm.DB {
	once.Do(func() {
		if consts.RunFrom == consts.FromServer && config.CONFIG.System.DbType == "mysql" {
			db = GormMySQL()
		} else {
			db = GormSQLLite()
		}
	})
	/*
		db.Use(dbresolver.Register(dbresolver.Config{
			// use `db2` as sources, `db3`, `db4` as replicas
			Sources:  []gorm.Dialector{mysql.Open("db2_dsn")},
			Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
			// sources/replicas load balancing policy
			Policy: dbresolver.RandomPolicy{},
			// print sources/replicas mode in logger
			TraceResolverMode: true,
		}))
	*/
	/*
		if len(args) > 0 {
			//ctx := context.Background()
			//x := context.WithValue(ctx, "dbName", "deeptest")
			db.Statement.Table = "xxxxxxxxxx"
			fmt.Println(db, "+++++")
		}

	*/

	return db
}

// MysqlTables 注册数据库表专用
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		logUtils.Errorf("注册数据表错误", zap.Any("err", err))
		os.Exit(0)
	}
	logUtils.Infof("注册数据表成功")
}

func GormSQLLite() *gorm.DB {
	conn := DBFile()
	dialector := sqlite.Open(conn)

	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
		},
	})

	if err != nil {
		logUtils.Info(err.Error())
	}

	_ = db.Use(
		dbresolver.Register(
			dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	db.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})

	//err = db.AutoMigrate(
	//	modelRef.Models...,
	//)
	//if err != nil {
	//	logUtils.Info(err.SendErrorMsg())
	//}

	return db
}

func DBFile() string {
	path := filepath.Join(consts.WorkDir, consts.RunFrom.String()+".db")
	return path
}

// GormMySQL 初始化Mysql数据库
func GormMySQL() *gorm.DB {
	m := config.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	var gormConf = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch config.CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		gormConf.Logger = Default.LogMode(logger.Silent)
	case "error", "SendErrorMsg":
		gormConf.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		gormConf.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		gormConf.Logger = Default.LogMode(logger.Info)
	case "zap", "Zap":
		gormConf.Logger = Default.LogMode(logger.Info)
	default:
		if mod {
			gormConf.Logger = Default.LogMode(logger.Info)
			break
		}
		gormConf.Logger = Default.LogMode(logger.Silent)
	}
	return gormConf
}
