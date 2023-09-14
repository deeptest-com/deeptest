package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type MockReqJson struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type MockJsExpression struct {
	ID uint `json:"id"`

	Name       string `json:"name"`
	Expression string `json:"expression"`
	Format     string `json:"format"`
	Desc       string `json:"desc"`
	Ordr       int    `json:"ordr"`
	Type       string `json:"type"`

	Result interface{} `gorm:"-" json:"result"`
}

type MockExpectIdsReq []uint

type MockExpectRequestOption struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type MockExpectRequestOptions map[consts.ParamIn][]MockExpectRequestOption

type CreateExampleReq struct {
	EndpointId uint              `json:"endpointId"`
	Method     consts.HttpMethod `json:"method"`
	Code       string            `json:"code"`
}
