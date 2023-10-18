package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type EndpointReqPaginate struct {
	_domain.PaginateReq
	ProjectId    int64    `json:"projectId"`
	CategoryId   int64    `json:"categoryId"`
	Status       []int64  `json:"status"`
	CreateUser   []string `json:"createUser"`
	Title        string   `json:"title"`
	ServeId      uint     `json:"serveId"`
	ServeVersion string   `json:"serveVersion"`
	TagNames     []string `json:"tagNames"`
}

type EndpointInterfaceReqPaginate struct {
	_domain.PaginateReq

	ServeId    uint   `json:"serveId"`
	Keywords   string `json:"Keywords"`
	CategoryId int64  `json:"categoryId"`
	ProjectId  uint   `json:"projectId"`
}

type OpenApiParam struct {
	domain.Param
	Format      string  `json:"format"`
	Example     string  `json:"example"`
	Pattern     string  `json:"pattern"`
	MinLength   int64   `json:"minLength"`
	MaxLength   int64   `json:"maxLength"`
	Default     string  `json:"default"`
	MultipleOf  int64   `json:"multipleOf"`
	MinItems    int64   `json:"minItems"`
	MaxItems    int64   `json:"maxItems"`
	UniqueItems bool    `json:"uniqueItems"`
	Ref         string  `json:"ref"`
	Required    bool    `json:"required"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Minimum     float64 `json:"minimum"`
	Maximum     float64 `json:"maximum"`
}

type EndpointReq struct {
	ID          int64           `json:"id"`
	ProjectId   uint            `json:"projectId" validate:"required"`
	ServeId     uint            `json:"serveId" validate:"required"`
	ServerId    uint            `json:"serverId"`
	Status      int64           `json:"status"`
	Title       string          `json:"title" validate:"required"`
	Path        string          `json:"path"`
	Version     string          `json:"version"`
	CreateUser  string          `json:"CreateUser"`
	UpdateUser  string          `json:"updateUser"`
	CategoryId  int64           `json:"categoryId"`
	PathParams  []OpenApiParam  `gorm:"-" json:"pathParams"`
	Interfaces  []InterfaceResp `gorm:"-" json:"interfaces"`
	Description string          `json:"description"`
	Curl        string          `json:"curl"`
}

type EndpointRes struct {
	ProjectId  int64           `json:"project_id"`
	Status     int64           `json:"status"`
	Title      string          `json:"title"`
	Version    string          `json:"version"`
	PathParams []OpenApiParam  `json:"pathParams"`
	Interfaces []InterfaceResp `json:"interfaces"`
}

type EndpointVersionReq struct {
	EndpointId int64  `json:"endpointId"`
	Version    string `json:"version"`
}

type ImportEndpointDataReq struct {
	ServeId       uint                `json:"serveId" validate:"required"`    //服务ID
	DriverType    convert.DriverType  `json:"driverType" validate:"required"` //接口数据来源
	CategoryId    int64               `json:"categoryId"`                     //所属分类
	DataSyncType  consts.DataSyncType `json:"dataSyncType"`                   //数据同步方式
	OpenUrlImport bool                `json:"openUrlImport"`                  //开启url导入
	FilePath      string              `json:"filePath" validate:"required"`
	ProjectId     uint                `json:"projectId"`
	UserId        uint                `json:"userId"`
	SourceType    consts.SourceType   `json:"sourceType"`
}

type BatchUpdateReq struct {
	FieldName   string      `json:"fieldName"`
	Value       interface{} `json:"value"`
	EndpointIds []uint      `json:"endpointIds"`
}

type EndpointTagReq struct {
	Id       uint     `json:"id"`
	TagNames []string `json:"tagNames"`
}

type GenerateFromResponseReq struct {
	Code        string `json:"code"`
	ContentType string `json:"contentType"`
	Description string `json:"description"`
	InterfaceId uint   `json:"interfaceId"`
	Data        string `json:"data"`
}

type GenerateFromRequestReq struct {
	ContentType string `json:"contentType"`
	Description string `json:"description"`
	InterfaceId uint   `json:"interfaceId"`
	Data        string `json:"data"`
}
