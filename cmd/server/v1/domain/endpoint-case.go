package serverDomain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
	casesHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/cases"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
	"github.com/kataras/iris/v12"
)

type EndpointCaseReqPaginate struct {
	_domain.PaginateReq

	EndpointId int `json:"endpointId"`

	Keywords string `json:"keywords"`
	Enabled  string `json:"enabled"`
}

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
	CaseType  consts.CaseType  `json:"caseType"`
	BaseCase  uint             `json:"baseCase"`
}

type EndpointCaseBenchmarkCreateReq struct {
	Type string `json:"type"`
	Name string `json:"name"`

	EndpointInterfaceId uint `json:"endpointInterfaceId"` // from a endpointInterface
	BaseCaseId          int  `json:"baseCaseId"`          // from a exist case

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
}

type EndpointCaseFactorSaveReq struct {
	Path   string `json:"path"`
	Value  string `json:"value"`
	CaseId int    `json:"caseId"`
}

type EndpointCaseAlternativeSaveReq struct {
	Type   string                      `json:"type"`
	BaseId int                         `json:"baseId"`
	Values casesHelper.AlternativeCase `json:"values"`

	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
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
