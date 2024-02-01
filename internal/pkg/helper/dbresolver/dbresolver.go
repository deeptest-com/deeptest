package dbresolver

import (
	"fmt"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"gorm.io/gorm"
	"sync"
)

type DBResolver struct {
	connPools sync.Map
}

func NewDBResolver() *DBResolver {
	return &DBResolver{connPools: sync.Map{}}
}

func (dr *DBResolver) Apply(poolName string, handler func() (*gorm.DB, error)) *DBResolver {
	if poolName == "" {
		return dr
	}
	_, ok := dr.connPools.Load(poolName)
	if !ok {
		connPool, err := handler()
		if err != nil {
			panic(err)
		}

		dr.connPools.Store(poolName, connPool)
	}

	return dr
}

func (dr *DBResolver) GetConnPool(poolName string) *gorm.DB {
	if poolName == "" {
		return nil
	}
	ret, ok := dr.connPools.Load(poolName)
	if !ok {
		panic(fmt.Errorf("connPool %s is not initialize", poolName))
	}

	db := ret.(*gorm.DB)
	if db == nil {
		dr.connPools.Delete(poolName)
		return nil
	}

	if pinger, ok := db.ConnPool.(interface{ Ping() error }); ok {
		err := pinger.Ping()
		if err != nil {
			dr.connPools.Delete(poolName)
			logUtils.Errorf("poolName%s,ping failed", poolName)
		}
	}

	return db
}
