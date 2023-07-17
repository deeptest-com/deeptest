package serverDomain

type ConditionMoveReq struct {
	Data []int `json:"data"`

	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
}
