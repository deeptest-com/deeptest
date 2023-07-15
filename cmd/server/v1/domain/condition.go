package serverDomain

import serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"

type ConditionMoveReq struct {
	DragId uint `json:"dragId"`
	DropId uint `json:"dropId"`

	Position serverConsts.DropPos `json:"position"`

	DebugInterfaceId    uint `json:"debugInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
}
