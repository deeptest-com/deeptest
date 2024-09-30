package db

import (
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/saas/tenant"
)

func GetByTenantId(tenantId consts.TenantId) (m config.Mysql) {
	ret := tenant.NewTenant()
	dbConfig, err := ret.GetDbConfig(tenantId)
	if err != nil {
		panic(err)
	}

	m = config.Mysql{}

	m.Url, m.Username, m.Password, m.Dbname, m.Config, m.LogMode = dbConfig.Path, dbConfig.Username, dbConfig.Password, dbConfig.Dbname, dbConfig.Config, true
	/*
		m.Url = "127.0.0.1:3306"
		m.Username = "root"
		m.Password = "root"
		m.Dbname = "deeptest"
		m.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	*/
	return
}
