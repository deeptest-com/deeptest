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
	ret, ok := dr.connPools.Load(poolName)
	if !ok {
		panic(fmt.Errorf("connPool %s is not initialize", poolName))
	}
	db := ret.(*gorm.DB)
	return db
}
