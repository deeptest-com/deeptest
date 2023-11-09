package dbresolver

import (
	"fmt"
	"gorm.io/gorm"
)

func (dr *DBResolver) registerCallbacks(db *gorm.DB) {
	dr.Callback().Create().Before("*").Register("gorm:db_resolver", dr.selectDB)
	dr.Callback().Query().Before("*").Register("gorm:db_resolver", dr.selectDB)
	dr.Callback().Update().Before("*").Register("gorm:db_resolver", dr.selectDB)
	dr.Callback().Delete().Before("*").Register("gorm:db_resolver", dr.selectDB)
	dr.Callback().Row().Before("*").Register("gorm:db_resolver", dr.selectDB)
	dr.Callback().Raw().Before("*").Register("gorm:db_resolver", dr.selectDB)
}

func (dr *DBResolver) selectDB(db *gorm.DB) {
	if dr.Context == nil {
		return
	}
	fmt.Println("selectDB....", dr.Context.Path())
	//dbname := fmt.Sprintf("%s.", db.Statement.Context.Value("dbName"))
	//db.Statement.Table = dbname + db.Statement.Table

}
