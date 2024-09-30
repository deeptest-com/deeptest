package serverDomain

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
	_domain "github.com/deeptest-com/deeptest/pkg/domain"
)

type ServeReqPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Name      string `json:"name"`
}

type ServeReq struct {
	ProjectId   uint   `json:"projectId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Status      int64  `json:"status"`
	ID          int64  `json:"id"`
	CreateUser  string `json:"createUser"`
	Description string `json:"description"`
}

type ServeVersionReq struct {
	ServeId     int64  `json:"serveId"`
	Value       string `json:"value"`
	ID          uint   `json:"id"`
	CreateUser  string `json:"createUser"`
	Description string `json:"description"`
}

type ServeVersionPaginate struct {
	_domain.PaginateReq
	ServeId    int64  `json:"serveId"`
	Version    string `json:"version"`
	CreateUser string `json:"createUser"`
}

type ServeSchemaReq struct {
	ID          uint   `json:"id"`
	ServeId     int64  `json:"serveId"`
	Name        string `json:"name" validate:"required"`
	Tag         string `json:"tag"`
	Content     string `json:"content"`
	Examples    string `json:"examples"`
	Type        string `json:"type"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
	Ref         string `json:"ref"`
	ProjectId   uint   `json:"projectId"`
	TargetId    int    `json:"targetId"`
}

type ServeSchemaRefReq struct {
	ServeId int64  `json:"serveId" validate:"required"`
	Ref     string `json:"ref" validate:"required"`
}

type EnvironmentReq struct {
	ID           uint
	ProjectId    uint                  `json:"projectId" validate:"required"`
	Name         string                `json:"name" validate:"required"`
	ServeServers []ServeServer         `json:"serveServers"`
	Vars         []EnvironmentVariable `json:"vars"`
}

type ServeServer struct {
	ID              uint   `json:"id"`
	ServeId         uint   `json:"serveId"`
	ServerId        uint   `json:"serverId"` // load by server id in scenario design page
	Url             string `json:"url"`
	EnvironmentName string `json:"environmentName"`
}

type EnvironmentVariable struct {
	Name        string `json:"name" validate:"required"`
	LocalValue  string `json:"localValue" validate:"required"`
	RemoteValue string `json:"remoteValue" validate:"required"`
	Description string `json:"description"`
}

type ServeSchemaPaginate struct {
	_domain.PaginateReq
	ProjectId int64  `json:"projectId"`
	Tag       string `json:"tag"`
	Type      string `json:"type"`
	Name      string `json:"name"`
}

type JsonContent struct {
	ProjectId uint   `json:"projectId"`
	Data      string `json:"data"`
}

type SchemaContent struct {
	Data string
}

type EndpointVersions struct {
	EndpointId int64  `json:"endpointId"`
	Version    string `json:"version"`
}

type ServeVersionBindEndpointReq struct {
	ServeId          int64              `json:"serveId"`
	ServeVersion     string             `json:"serveVersion"`
	EndpointVersions []EndpointVersions `json:"endpointVersions"`
}

type EnvironmentParam struct {
	Name         string `json:"name" validate:"required"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
	DefaultValue string `json:"defaultValue"`
	Description  string `json:"description"`
}

type EnvironmentParamsReq struct {
	ProjectId uint               `json:"projectId"`
	Header    []EnvironmentParam `json:"header"`
	Cookie    []EnvironmentParam `json:"cookie"`
	Query     []EnvironmentParam `json:"query"`
	Body      []EnvironmentParam `json:"body"`
	Path      []EnvironmentParam `json:"path"`
}

type EnvironmentIdsReq []uint

type ServeSecurityPaginate struct {
	_domain.PaginateReq
	ServeId int64  `json:"serveId"`
	Name    string `json:"name"`
}

type ServeSecurityReq struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name" validate:"required"`
	Type        serverConsts.AuthType `json:"type" validate:"required"`
	ProjectId   int64                 `json:"projectId" validate:"required"`
	ServeId     int64                 `json:"serveId" validate:"required"`
	Description string                `json:"description"`
	In          string                `json:"in"`
	Key         string                `json:"key"`
	Value       string                `json:"value"`
	Token       string                `json:"token"`
	Username    string                `json:"username"`
	Password    string                `json:"password"`
	Default     bool                  `json:"default"`
}

type ChangeServeReq struct {
	Id uint
}

type SwaggerSyncReq struct {
	ID         int                 `json:"id"`
	Switch     consts.SwitchStatus `json:"switch"`
	SyncType   consts.DataSyncType `json:"syncType"`
	CategoryId int                 `json:"categoryId"`
	Url        string              `json:"url"`
	Cron       string              `json:"cron"`
	ProjectId  uint                `json:"projectId"`
}

type MockReq struct {
	ID        uint                `json:"id"`
	Priority  consts.MockPriority `json:"priority"`
	ProjectId uint                `json:"projectId"`
}

type HistoryServeAddServesReq struct {
	ServerName string `json:"serverName"`
	Url        string `json:"url"`
}

type SaveSchemaRes struct {
	EntityId   uint `json:"entityId"`
	CategoryId uint `json:"categoryId"`
}
