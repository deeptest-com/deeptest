package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
)

type CronConfigLecang struct {
	BaseModel
	SyncType         consts.DataSyncType               `json:"syncType"`
	CategoryId       int                               `json:"categoryId"`
	Url              string                            `json:"url"`
	Cron             string                            `json:"cron"`
	ProjectId        uint                              `json:"projectId"`
	ServeId          uint                              `json:"serveId"`
	CreateUserId     uint                              `json:"createUserId"`
	Engineering      string                            `json:"engineering"`    //所属工程
	ServiceCodes     string                            `json:"serviceCodes"`   //服务名，逗号分隔
	MessageType      consts.CronLecangMessageType      `json:"messageType"`    //消息类型
	ExtendOverride   consts.CronLecangIsExtendOverride `json:"extendOverride"` //继承父类
	Overridable      string                            `json:"overridable"`    //是否允许重写
	AddServicePrefix bool                              `json:"addServicePrefix"`
}

func (CronConfigLecang) TableName() string {
	return "biz_project_cron_config_lecang"
}
