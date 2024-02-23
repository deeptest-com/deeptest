package agentExec

import (
	"database/sql"
	"encoding/json"
	"fmt"
	queryUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/query"
	valueUtils "github.com/aaronchen2k/deeptest/internal/agent/exec/utils/value"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_netUtils "github.com/aaronchen2k/deeptest/pkg/lib/net"
	driverMysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

func ExecDbOpt(opt *domain.DatabaseOptBase) (err error) {
	if opt.Type == "" || opt.DbConnId == 0 || opt.Sql == "" {
		opt.ResultStatus = consts.Fail
		opt.Result = "数据库操作未初始化"

		return
	}

	if opt.DatabaseConnIsDisabled {
		opt.ResultStatus = consts.Fail
		opt.Result = "数据库操作已禁用"
		return
	}

	ok := _netUtils.Ping(opt.Host, opt.Port)
	if !ok {
		opt.ResultStatus = consts.Fail
		opt.Result = "数据库连接超时"
		return
	}

	opt.ResultStatus = consts.Pass

	if opt.Type == consts.DbTypeOracle {
		orclDb, err1 := OpenOracle(opt)
		if err1 != nil {
			err = err1
			opt.ResultStatus = consts.Fail
			return
		}

		err1 = queryOracle(orclDb, opt)
		if err1 != nil {
			err = err1
			opt.ResultStatus = consts.Fail
			return
		}

		return
	}

	var db *gorm.DB

	if opt.Type == consts.DbTypeMySql {
		db, err = OpenMySqlDb(opt)
	} else if opt.Type == consts.DbTypeSqlServer {
		db, err = OpenSqlServer(opt)
	} else if opt.Type == consts.DbTypePostgreSql {
		db, err = OpenPostgreSQL(opt)
	} else if opt.Type == consts.DbTypeOracle {
		db, err = OpenPostgreSQL(opt)
	}

	if err != nil {
		opt.Result = err.Error()
		opt.ResultStatus = consts.Fail
		return
	}

	queryResult, err := query(db, opt)
	if err != nil {
		opt.Result = err.Error()
		opt.ResultStatus = consts.Fail
		return
	}

	if opt.JsonPath != "" {
		var result interface{}

		result, opt.ResultType, err = queryUtils.JsonPath(string(queryResult), opt.JsonPath)
		if err != nil {
			opt.Result = err.Error()
			opt.ResultStatus = consts.Fail
			return
		}

		opt.Result = valueUtils.InterfaceToStr(result)
	}

	return
}

func OpenMySqlDb(opt *domain.DatabaseOptBase) (db *gorm.DB, err error) {
	params := "charset=utf8mb4&parseTime=True&loc=Local"

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", opt.Username, opt.Password,
		opt.Host, opt.Port, opt.DbName, params)

	dsnConfig := driverMysql.Config{
		Timeout: 1 * time.Second,
	}
	config := mysql.Config{
		DSN:       connStr,
		DSNConfig: &dsnConfig,
	}

	db, err = gorm.Open(mysql.New(config), &gorm.Config{})

	return
}

func OpenSqlServer(opt *domain.DatabaseOptBase) (db *gorm.DB, err error) {
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		opt.Username, opt.Password,
		opt.Host, opt.Port, opt.DbName)

	config := sqlserver.Config{
		DSN: connStr,
	}

	db, err = gorm.Open(sqlserver.New(config), &gorm.Config{})

	return
}

func OpenPostgreSQL(opt *domain.DatabaseOptBase) (db *gorm.DB, err error) {
	params := "sslmode=disable TimeZone=Asia/Shanghai"

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s %s",
		opt.Username, opt.Password,
		opt.Host, opt.Port, opt.DbName, params)

	config := postgres.Config{
		DSN: connStr,
	}

	db, err = gorm.Open(postgres.New(config), &gorm.Config{})

	return
}

func OpenOracle(opt *domain.DatabaseOptBase) (db *sql.DB, err error) {
	connStr := fmt.Sprintf("%s/%s@%s:%s/%s",
		opt.Username, opt.Password,
		opt.Host, opt.Port, opt.DbName)

	db, err = sql.Open("goracle", connStr)

	return
}

func query(db *gorm.DB, opt *domain.DatabaseOptBase) (result []byte, err error) {
	data := []map[string]interface{}{}

	err = db.Raw(opt.Sql).
		Scan(&data).Error

	result, err = json.Marshal(data)

	return
}

func queryOracle(db *sql.DB, opt *domain.DatabaseOptBase) (err error) {
	rows, err := db.Query(opt.Sql, 100)
	if err != nil {
		return
	}
	defer rows.Close()

	cols, _ := rows.Columns()

	data := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err = rows.Scan(columnPointers...); err != nil {
			return
		}

		mp := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			mp[colName] = *val
		}

		data = append(data, mp)
	}

	bytes, err := json.Marshal(data)
	opt.ResultMsg = string(bytes)

	return
}
