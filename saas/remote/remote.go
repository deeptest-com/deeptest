package remote

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/integration/leyan/pkg"
	"github.com/aaronchen2k/deeptest/internal/pkg/config"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	v1 "github.com/aaronchen2k/deeptest/internal/pkg/domain"
	httpHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/http"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/aaronchen2k/deeptest/saas/domain"
)

type Remote struct {
}

func (s *Remote) GetTenant(tenantId consts.TenantId) (ret domain.Tenant) {
	url := fmt.Sprintf("%s/api/v1/tenant", config.CONFIG.Saas.Url)

	headers := pkg.GetHeaders("")
	httpReq := v1.BaseRequest{
		Url:         url,
		BodyType:    consts.ContentTypeJSON,
		Headers:     &headers,
		QueryParams: &[]v1.Param{{Name: "id", Value: string(tenantId)}},
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
	type Data struct {
		Id       int64           `json:"id"`
		DbConfig domain.DbConfig `json:"pjtDB"`
	}
	respContent := struct {
		Code int
		Data Data
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

	ret.Id, ret.DbConfig = consts.TenantId(fmt.Sprintf("%d", respContent.Data.Id)), respContent.Data.DbConfig
	return
}

func (s *Remote) GetTenants() (ret []domain.Tenant) {

	url := fmt.Sprintf("%s/api/v1/tenant/list", config.CONFIG.Saas.Url)

	headers := pkg.GetHeaders("")
	httpReq := v1.BaseRequest{
		Url:      url,
		BodyType: consts.ContentTypeJSON,
		Headers:  &headers,
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
