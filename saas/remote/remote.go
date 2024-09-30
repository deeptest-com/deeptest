package remote

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest/integration/thirdparty/pkg"
	"github.com/deeptest-com/deeptest/internal/pkg/config"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	_http "github.com/deeptest-com/deeptest/pkg/lib/http"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
	"github.com/deeptest-com/deeptest/saas/domain"
)

type Remote struct {
}

func (s *Remote) GetTenant(tenantId consts.TenantId, prefix string) (ret domain.Tenant) {
	url := fmt.Sprintf("%s/api/v1/tenant", config.CONFIG.Saas.Url)

	headers := pkg.GetHeaders("")

	resp, code, err := _http.Get(fmt.Sprintf("%s?id=%s&customDomainPrefix=%s", url, tenantId, prefix), headers)
	if err != nil {
		logUtils.Errorf("get tenant failed, error, %s", err.Error())
		return
	}

	if code != consts.OK.Int() {
		logUtils.Errorf("get tenant failed, response %v", resp)
		err = fmt.Errorf("get tenant/list failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data domain.Tenant
		Msg  string
	}{}
	err = json.Unmarshal(resp, &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 0 {
		logUtils.Errorf("tenant failed, response %v", resp)
		err = fmt.Errorf("get tenant/list failed, response %v", resp)
		return
	}

	ret = respContent.Data
	return
}

func (s *Remote) GetTenants() (ret []domain.Tenant) {

	url := fmt.Sprintf("%s/api/v1/tenant/list", config.CONFIG.Saas.Url)

	headers := pkg.GetHeaders("")

	resp, code, err := _http.Get(fmt.Sprintf("%s?page=1&pageSize=9999999999", url), headers)
	if err != nil {
		logUtils.Errorf("get tenant/list failed, error, %s", err.Error())
		return
	}

	if code != consts.OK.Int() {
		logUtils.Errorf("get tenant/list failed, response %v", resp)
		err = fmt.Errorf("get tenant/list failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []domain.Tenant
		Msg  string
	}{}
	err = json.Unmarshal(resp, &respContent)
	if err != nil {
		logUtils.Infof(err.Error())
	}

	if respContent.Code != 0 {
		logUtils.Errorf("tenant/list failed, response %v", resp)
		err = fmt.Errorf("get tenant/list failed, response %v", resp)
		return
	}

	ret = respContent.Data
	return

}
