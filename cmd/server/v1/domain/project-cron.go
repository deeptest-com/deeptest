package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ProjectCronReq struct {
	Id           uint                `json:"id"`
	Name         string              `json:"name"`
	Cron         string              `json:"cron"`
	ProjectId    uint                `json:"projectId"`
	Source       consts.CronSource   `json:"source"`
	Switch       consts.SwitchStatus `json:"switch"`
	ConfigId     uint                `json:"configId"`
	CreateUserId uint                `json:"createUserId"`
	UpdateUserId uint                `json:"updateUserId"`
}

type LecangCronReq struct {
	Token            string              `json:"token"`
	SyncType         consts.DataSyncType `json:"syncType"`
	ProjectId        uint                `json:"projectId"`
	CategoryId       int                 `json:"categoryId"`
	Url              string              `json:"url"`
	ServeId          uint                `json:"serveId"`
	CreateUserId     uint                `json:"createUserId"`
	ServiceCode      string              `json:"serviceCode"` //乐仓服务名
	AddServicePrefix bool                `json:"addServicePrefix"`
	LecangFuncLimit
}

type LecangFuncLimit struct {
	MessageType    consts.CronLecangMessageType      `json:"messageType"`    //消息类型
	ExtendOverride consts.CronLecangIsExtendOverride `json:"extendOverride"` //继承父类
	Overridable    consts.IntegrationFuncOverridable `json:"overridable"`    //是否允许重写
}

type ProjectCronReqPaginate struct {
	_domain.PaginateReq

	ProjectId uint                `json:"projectId"`
	Name      string              `json:"name"`
	Source    consts.CronSource   `json:"source"`
	Switch    consts.SwitchStatus `json:"switch"`
}

type SaveLcEndpointReq struct {
	Title         string              `json:"title"`
	ProjectId     uint                `json:"projectId"`
	ServeId       uint                `json:"serveId"`
	UserId        uint                `json:"userId"`
	OldEndpointId uint                `json:"oldEndpointId"`
	Path          string              `json:"path"`
	Snapshot      string              `json:"snapshot"`
	DataSyncType  consts.DataSyncType `json:"dataSyncType"`
	CategoryId    int64               `json:"categoryId"`
}
