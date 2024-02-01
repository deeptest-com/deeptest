package dbresolver

import (
	"fmt"
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
		poolName = "default"
	}
	_, ok := dr.connPools.Load(poolName)
	if !ok {
		connPool, err := handler()
		if err != nil {
			panic(err)
		}
		dr.connPools.Store(poolName, connPool)
		//	dr.connPools.Store(poolName, "xxx")
	}
	return dr
}

func (dr *DBResolver) GetConnPool(poolName string) *gorm.DB {
	if poolName == "" {
		poolName = "default"
	}
	ret, ok := dr.connPools.Load(poolName)
	if !ok {
		panic(fmt.Errorf("connPool %s is not initialize", poolName))
	}
	db := ret.(*gorm.DB)
	return db
}
