package domain

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
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
	Id       consts.TenantId `json:"id"`
	SpecCode string          `json:"specCode"`
	SkuCode  string          `json:"skuCode"`
	DbConfig DbConfig        `json:"leyanapiDB"`
}

type temp struct {
	Id       int64    `json:"id"`
	DbConfig DbConfig `json:"leyanapiDB"`
	SpecCode string   `json:"specCode"`
	SkuCode  string   `json:"skuCode"`
}

func (tenant *Tenant) MarshalJSON() (res []byte, err error) {
	x := temp{}
	x.Id, _ = strconv.ParseInt(string(tenant.Id), 10, 64)
	x.DbConfig = tenant.DbConfig
	x.SpecCode = tenant.SpecCode
	x.SkuCode = tenant.SkuCode
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

	return nil
}
