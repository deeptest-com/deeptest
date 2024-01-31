package tenant

import (
	"github.com/aaronchen2k/deeptest/saas/domain"
)

type tenant struct {
	Id       string          `json:"id"`
	DbConfig domain.DbConfig `json:"dbConfig"`
}

func NewTenant() *tenant {
	return new(tenant)
}

func (t *tenant) GetInfo(tenantId string) {
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
}

func (t *tenant) GetDbConfig(tenantId string) (config domain.DbConfig, err error) {
	return
}
