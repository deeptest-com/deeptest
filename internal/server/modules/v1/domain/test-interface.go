package serverDomain

import (
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

// maintain
type TestInterfaceMaintReq struct {
	Mode      serverConsts.NodeCreateMode `json:"mode"`
	Type      serverConsts.NodeCreateType `json:"type"`
	Target    int                         `json:"target"`
	Name      string                      `json:"name"`
	Id        int                         `json:"id"`
	ProjectId int                         `json:"projectId"`
}

type TestInterfaceMoveReq struct {
	DragKey int                  `json:"dragKey"`
	DropKey int                  `json:"dropKey"`
	DropPos serverConsts.DropPos `json:"dropPos"`
}

type TestInterfaceMaintResp struct {
	model.TestInterface
}
