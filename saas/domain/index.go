package domain

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"strconv"
)

type DbConfig struct {
	Path            string `json:"path"`
	Config          string `json:"Config"`
	Dbname          string `json:"dbname"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	SchemaType      string `json:"schemaType"`
	Maxidleconns    int64  `json:"maxidleconns"`
	Maxopenconns    int64  `json:"maxopenconns"`
	Connmaxlifetime int64  `json:"connmaxlifetime"`
}

type Tenant struct {
	Id            consts.TenantId `json:"id"`
	SpecCode      string          `json:"specCode"`
	SkuCode       string          `json:"skuCode"`
	ManagerId     uint64          `json:"managerId"`
	ManagerMobile string          `json:"managerMobile"`
	ManagerName   string          `json:"managerName"`
	ManagerMail   string          `json:"managerMail"`

	DbConfig DbConfig `json:"thirdpartyapiDB"`
}

type temp struct {
	Id            int64    `json:"id"`
	DbConfig      DbConfig `json:"thirdpartyapiDB"`
	SpecCode      string   `json:"specCode"`
	SkuCode       string   `json:"skuCode"`
	ManagerId     uint64   `json:"managerId"`
	ManagerMobile string   `json:"managerMobile"`
	ManagerName   string   `json:"managerName"`
	ManagerMail   string   `json:"managerMail"`
}

func (tenant *Tenant) MarshalJSON() (res []byte, err error) {
	x := temp{}
	x.Id, _ = strconv.ParseInt(string(tenant.Id), 10, 64)
	x.DbConfig = tenant.DbConfig
	x.SpecCode = tenant.SpecCode
	x.SkuCode = tenant.SkuCode
	x.ManagerMail = tenant.ManagerMail
	x.ManagerId = tenant.ManagerId
	x.ManagerName = tenant.ManagerName
	x.ManagerMobile = tenant.ManagerMobile
	return json.Marshal(x)
}

func (tenant *Tenant) UnmarshalJSON(data []byte) error {

	var x temp
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}

	tenant.Id = consts.TenantId(fmt.Sprintf("%d", x.Id))
	tenant.DbConfig = x.DbConfig
	tenant.SkuCode = x.SkuCode
	tenant.SpecCode = x.SpecCode
	tenant.ManagerMail = x.ManagerMail
	tenant.ManagerId = x.ManagerId
	tenant.ManagerName = x.ManagerName
	tenant.ManagerMobile = x.ManagerMobile

	return nil
}
