package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
)

type TestInterfaceReq struct {
	TargetId int    `json:"targetId"`
	Name     string `json:"name"`
	Mode     string `json:"mode"`

	Id        int `json:"id"`
	ProjectId int `json:"projectId"`
}

type TestInterfaceResp struct {
	model.TestInterface
}
