package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
	uuid "github.com/satori/go.uuid"
)

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
	EndpointId uint   `json:"endpointId"`
	ServeId    uint   `json:"serveId"`
	ProjectId  uint   `json:"projectId"`

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`

	DebugInterfaceId    int    `json:"debugInterfaceId"`
	EndpointInterfaceId int    `json:"endpointInterfaceId"`
	UsedBy              string `json:"usedBy"`

	DebugData domain.DebugData `json:"debugData"`
}

type EndpointCaseTree struct {
	Key int64     `json:"key"`
	Id  uuid.UUID `json:"id"`

	Name   string                            `json:"name"`
	Method string                            `json:"method"`
	Desc   string                            `json:"desc"`
	Type   serverConsts.EndpointCaseTreeType `json:"type"`
	IsDir  bool                              `json:"isDir"`

	CategoryId       int64 `json:"categoryId"`
	EndpointId       uint  `json:"endpointId"`
	DebugInterfaceId uint  `json:"debugInterfaceId"`
	CaseInterfaceId  uint  `json:"caseInterfaceId"`
	ParentId         int64 `json:"parentId"`
	ProjectId        uint  `json:"projectId"`
	ServeId          uint  `json:"serveId"`
	UseID            uint  `json:"useId"`

	Children []*EndpointCaseTree `json:"children"`
	Slots    iris.Map            `json:"slots"`
	Count    int64               `json:"count"`
	Ordr     int                 `json:"ordr"`
}

type EndpointCount struct {
	Count      int64
	EndpointId int64
}

type InterfaceCase struct {
	ID               uint   `json:"ID"`
	Name             string `json:"name"`
	Method           string `json:"method"`
	Desc             string `json:"desc"`
	EndpointId       uint   `json:"endpointId"`
	ServeId          uint   `json:"serveId"`
	ProjectId        uint   `json:"projectId"`
	DebugInterfaceId uint   `json:"debugInterfaceId"`
}
