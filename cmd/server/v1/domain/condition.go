package serverDomain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
)

type ConditionMoveReq struct {
	Data               []int `json:"data"`
	IsForBenchmarkCase bool  `json:"isForBenchmarkCase"`

	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
}

type ExtractorConditionQuickCreateReq struct {
	Src          consts.ExtractorSrc `json:"src"`
	ConditionSrc consts.ConditionSrc `json:"conditionSrc"`

	Config domain.ExtractorBase `json:"conf"`
	Info   domain.DebugInfo     `json:"info"`
}
