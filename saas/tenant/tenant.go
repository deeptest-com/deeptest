package tenant

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/saas/domain"
	"github.com/aaronchen2k/deeptest/saas/remote"
)

type Tenant struct {
	Id       consts.TenantId `json:"id"`
	DbConfig domain.DbConfig `json:"dbConfig"`
}

func NewTenant() *Tenant {
	return new(Tenant)
}

func (t *Tenant) GetInfo(tenantId consts.TenantId) (tenant domain.Tenant) {
	tenant = new(remote.Remote).GetTenant(tenantId)
	return
}

func (t *Tenant) GetDbConfig(tenantId consts.TenantId) (config domain.DbConfig, err error) {
	res := t.GetInfo(tenantId)
	config = res.DbConfig
	return
}

func (t *Tenant) GetInfos() (tenants []domain.Tenant) {
	if config.CONFIG.Saas.Switch {
		tenants = new(remote.Remote).GetTenants()
	}
	return
}
