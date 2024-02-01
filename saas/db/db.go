package db

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/saas/tenant"
)

func GetByTenantId(tenantId string) (m config.Mysql) {
	ret := tenant.NewTenant()
	ret.GetDbConfig(tenantId)
	m = config.Mysql{}
	m.Url = "127.0.0.1:3306"
	m.Username = "root"
	m.Password = "root"
	m.Dbname = "deeptest"
	m.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	return
}
