package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type EndpointDiffReq struct {
	EndpointId uint `json:"endpointId"`
	IsChanged  bool `json:"isChanged"`
}

type EndpointDiffRes struct {
	Current       string               `json:"current"`
	Latest        string               `json:"latest"`
	CurrentDesc   string               `json:"currentDesc"`
	LatestDesc    string               `json:"latestDesc"`
	ChangedStatus consts.ChangedStatus `json:"changedStatus"`
}
