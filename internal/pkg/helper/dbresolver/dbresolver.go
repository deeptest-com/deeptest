package dbresolver

import (
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

type DBResolver struct {
	*gorm.DB
	Context *context.Context
}

func (dr *DBResolver) Name() string {
	return "gorm:db_resolver"
}

func (dr *DBResolver) Initialize(db *gorm.DB) error {
	dr.DB = db
	dr.registerCallbacks(db)
	//return dr.compile()
	return nil
}
