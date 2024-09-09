package remote

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/integration/thirdparty/pkg"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	v1 "github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/saas/domain"
)

type Remote struct {
}

func (s *Remote) GetTenant(tenantId consts.TenantId, prefix string) (ret domain.Tenant) {
	url := fmt.Sprintf("%s/api/v1/tenant", config.CONFIG.Saas.Url)

	headers := pkg.GetHeaders("")
	httpReq := v1.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
		//QueryParams: &[]v1.Param{{Name: "id", Value: string(tenantId)}, {Name: "customDomainPrefix", Value: prefix}, {Name: "env", Value: "local"}},
		QueryParams: &[]v1.Param{{Name: "id", Value: string(tenantId)}, {Name: "customDomainPrefix", Value: prefix}},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Errorf("get tenant failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Errorf("get tenant failed, response %v", resp)
		err = fmt.Errorf("get tenant/list failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data domain.Tenant
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
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
	httpReq := v1.BaseRequest{
		Url:         url,
		BodyType:    consts.ContentTypeJSON,
		Headers:     &headers,
		QueryParams: &[]v1.Param{{Name: "page", Value: "1"}, {Name: "pageSize", Value: "9999999999"}},
	}

	resp, err := httpHelper.Get(httpReq)
	if err != nil {
		logUtils.Errorf("get tenant/list failed, error, %s", err.Error())
		return
	}

	if resp.StatusCode != consts.OK.Int() {
		logUtils.Errorf("get tenant/list failed, response %v", resp)
		err = fmt.Errorf("get tenant/list failed, response %v", resp)
		return
	}

	respContent := struct {
		Code int
		Data []domain.Tenant
		Msg  string
	}{}
	err = json.Unmarshal([]byte(resp.Content), &respContent)
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
