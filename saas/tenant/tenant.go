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
	/*
		url := fmt.Sprintf("%s/api/v1/openApi/getUserDynamicMenuPermission", config.CONFIG.ThirdParty.Url)

		headers :=
		httpReq := domain.BaseRequest{
			Url:      url,
			BodyType: consts.ContentTypeJSON,
			Headers:  &headers,
			QueryParams: &[]domain.Param{
				{
					Name:  "typeStr",
					Value: "[20,30]",
				},
				{
					Name:  "username",
					Value: username,
				},
			},
		}

		resp, err := httpHelper.Get(httpReq)
		if err != nil {
			logUtils.Infof("get UserButtonPermissions failed, error, %s", err.Error())
			return
		}

		if resp.StatusCode != consts.OK.Int() {
			logUtils.Infof("get UserButtonPermissions failed, response %v", resp)
			err = fmt.Errorf("get UserButtonPermissions failed, response %v", resp)
			return
		}

		respContent := struct {
			Code int
			Data []string
			Msg  string
		}{}
		err = json.Unmarshal([]byte(resp.Content), &respContent)
		if err != nil {
			logUtils.Infof(err.Error())
		}

		if respContent.Code != 200 {
			logUtils.Infof("getUserButtonPermissions failed, response %v", resp)
			err = fmt.Errorf("get UserButtonPermissions failed, response %v", resp)
			return
		}

		ret = respContent.Data
	*/
	//	t.GetInfos()
	return
}

func (t *Tenant) GetDbConfig(tenantId consts.TenantId) (config domain.DbConfig, err error) {
	res := t.GetInfo(tenantId)
	config = res.DbConfig
	return
}

func (t *Tenant) GetInfos() (tenants []domain.Tenant) {
	if config.CONFIG.Saas.Switch == "ON" {
		tenants = new(remote.Remote).GetTenants()
	}
	return
}
