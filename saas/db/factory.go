package db

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/saas/tenant"
)

type Factory struct {
}

func (f *Factory) creatDb(tenantId consts.TenantId) {
	ret := tenant.NewTenant()
	ret.GetInfo(tenantId, "")

}
