package db

import "github.com/aaronchen2k/deeptest/saas/tenant"

type Factory struct {
}

func (f *Factory) creatDb(tenantId string) {
	ret := tenant.NewTenant()
	ret.GetInfo(tenantId)

}
