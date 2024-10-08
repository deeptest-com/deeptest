package tenant

import (
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/saas/domain"
	"github.com/deeptest-com/deeptest/saas/remote"
	"strings"
)

type Tenant struct {
	Id            consts.TenantId `json:"id"`
	DbConfig      domain.DbConfig `json:"dbConfig"`
	SpecCode      string          `json:"specCode"`
	SkuCode       string          `json:"skuCode"`
	ManagerId     uint64          `json:"managerId"`
	ManagerMobile string          `json:"managerMobile"`
	managerMail   string          `json:"managerMail"`
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
