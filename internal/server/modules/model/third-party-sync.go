package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"time"
)

type ThirdPartySync struct {
	BaseModel
	Switch       consts.SwitchStatus `json:"switch"`
	SyncType     consts.DataSyncType `json:"syncType"`
	CategoryId   int                 `json:"categoryId"`
	Url          string              `json:"url"`
	Cron         string              `json:"cron"`
	ProjectId    uint                `json:"projectId"`
	ServeId      uint                `json:"serveId"`
	ServiceCode  string              `json:"serviceCode"`
	ExecTime     *time.Time          `json:"execTime"`
	CreateUserId uint                `json:"createUserId"`
}

func (ThirdPartySync) TableName() string {
	return "biz_project_serve_third_party_sync"
}
