package db

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/saas/tenant"
)

type Factory struct {
}

func (f *Factory) creatDb(tenantId consts.TenantId) {
	ret := tenant.NewTenant()
	ret.GetInfo(tenantId, "")

}
