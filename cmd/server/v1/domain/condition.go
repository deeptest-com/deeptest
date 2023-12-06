package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type ConditionMoveReq struct {
	Data               []int `json:"data"`
	IsForBenchmarkCase bool  `json:"isForBenchmarkCase"`

	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
}

type ExtractorConditionQuickCreateReq struct {
	Config domain.ExtractorBase `json:"conf"`
	Info   domain.DebugInfo     `json:"info"`
}
