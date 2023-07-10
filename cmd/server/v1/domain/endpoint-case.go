package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type EndpointCase struct {
	Id int64 `json:"id"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	DebugInterfaceId uint  `json:"debugInterfaceId"`
	ParentId         int64 `json:"parentId"`
	ProjectId        uint  `json:"projectId"`
	ServeId          uint  `json:"serveId"`
	UseID            uint  `json:"useId"`
}

type EndpointCaseDetail struct {
	EndpointCase

	DebugData domain.DebugData `json:"debugData"`
}

type EndpointCaseSaveReq struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	ParentId   uint   `json:"parentId"`
	EndpointId uint   `json:"endpointId"`
	ServeId    uint   `json:"serveId"`
	ProjectId  uint   `json:"projectId"`

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
}
