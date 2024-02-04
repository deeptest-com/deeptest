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
	SwaggerReq   SwaggerSyncReq      `json:"swaggerReq"`
	LecangReq    LecangCronReq       `json:"lecangReq"`
}

type SwaggerCronReq struct {
}

type LecangCronReq struct {
	SyncType         consts.DataSyncType               `json:"syncType"`
	CategoryId       int                               `json:"categoryId"`
	Url              string                            `json:"url"`
	ServeId          uint                              `json:"serveId"`
	EngineeringCode  string                            `json:"engineeringCode"` //所属工程
	ServiceCodes     string                            `json:"serviceCodes"`    //服务名，逗号分隔
	MessageType      consts.CronLecangMessageType      `json:"messageType"`     //消息类型
	ExtendOverride   consts.CronLecangIsExtendOverride `json:"extendOverride"`  //继承父类
	Overridable      string                            `json:"overridable"`     //是否允许重写
	AddServicePrefix bool                              `json:"addServicePrefix"`
}

type ProjectCronReqPaginate struct {
	_domain.PaginateReq

	ProjectId uint                `json:"projectId"`
	Name      string              `json:"name"`
	Source    consts.CronSource   `json:"source"`
	Switch    consts.SwitchStatus `json:"switch"`
}
