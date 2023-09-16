package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/kataras/iris/v12"
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
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Method     consts.HttpMethod `json:"method"`
	Desc       string            `json:"desc"`
	EndpointId uint              `json:"endpointId"`
	ServeId    uint              `json:"serveId"`
	ProjectId  uint              `json:"projectId"`

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`

	DebugInterfaceId    int    `json:"debugInterfaceId"`
	EndpointInterfaceId int    `json:"endpointInterfaceId"`
	UsedBy              string `json:"usedBy"`

	DebugData domain.DebugData `json:"debugData"`
}

type EndpointCaseAlternativeLoadReq struct {
	EndpointId uint              `json:"endpointId"`
	Method     consts.HttpMethod `json:"method"`

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
}

type EndpointCaseAlternativeGenerateReq struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Method string `json:"method"`
	Desc   string `json:"desc"`

	EndpointId uint `json:"endpointId"`
	ServeId    uint `json:"serveId"`
	ProjectId  uint `json:"projectId"`

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`

	DebugInterfaceId    int    `json:"debugInterfaceId"`
	EndpointInterfaceId int    `json:"endpointInterfaceId"`
	UsedBy              string `json:"usedBy"`
}

type EndpointCaseTree struct {
	Key int64  `json:"key"`
	Id  string `json:"id"`

	Name   string                            `json:"name"`
	Method string                            `json:"method"`
	Desc   string                            `json:"desc"`
	Type   serverConsts.EndpointCaseTreeType `json:"type"`
	IsDir  bool                              `json:"isDir"`

	CategoryId       int64  `json:"categoryId"`
	EndpointId       uint   `json:"endpointId"`
	DebugInterfaceId uint   `json:"debugInterfaceId"`
	CaseInterfaceId  uint   `json:"caseInterfaceId"`
	ParentId         string `json:"parentId"`
	ProjectId        uint   `json:"projectId"`
	ServeId          uint   `json:"serveId"`
	UseID            uint   `json:"useId"`

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

type CategoryEndpointCase struct {
	CaseUniqueId         string `json:"caseUniqueId"`
	EndpointUniqueId     string `json:"endpoint_unique_id"`
	CaseId               uint   `json:"caseId"`
	CaseName             string `json:"caseName"`
	Method               string `json:"method"`
	CaseDesc             string `json:"caseDesc"`
	CaseDebugInterfaceId uint   `json:"caseDebugInterfaceId"`
	ServeId              uint   `json:"serveId"`
	ProjectId            uint   `json:"projectId"`
	EndpointId           uint   `json:"endpointId"`
	EndpointTitle        string `json:"endpointTitle"`
	EndpointDescription  string `json:"endpointDescription"`
	CategoryId           int64  `json:"categoryId"`
}
