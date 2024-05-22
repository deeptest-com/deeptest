package tenant

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/saas/domain"
	"github.com/aaronchen2k/deeptest/saas/remote"
	"strings"
)

type Tenant struct {
	Id       consts.TenantId `json:"id"`
	DbConfig domain.DbConfig `json:"dbConfig"`
	SpecCode string          `json:"specCode"`
	SkuCode  string          `json:"skuCode"`
}

func NewTenant() *Tenant {
	return new(Tenant)
}

func (t *Tenant) GetInfo(tenantId consts.TenantId, prefix string) (tenant domain.Tenant) {
	tenant = new(remote.Remote).GetTenant(tenantId, prefix)
	return
}

func (t *Tenant) GetDbConfig(tenantId consts.TenantId) (config domain.DbConfig, err error) {
	res := t.GetInfo(tenantId, "")
	config = res.DbConfig
	return
}

func (t *Tenant) GetInfos() (tenants []domain.Tenant) {
	if config.CONFIG.Saas.Switch {
		tenants = new(remote.Remote).GetTenants()
	}
	return
}

func (t *Tenant) ForFree(tenantId consts.TenantId) bool {
	tenant := new(remote.Remote).GetTenant(tenantId, "")
	version := strings.ReplaceAll(tenant.SpecCode, tenant.SkuCode, "")
	return version == "-01"

}
