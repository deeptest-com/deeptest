package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type SwaggerSync struct {
	BaseModel
	Switch     consts.SwitchStatus `json:"switch"`
	SyncType   consts.SyncType     `json:"syncType"`
	CategoryId int                 `json:"categoryId"`
	Url        string              `json:"url"`
	Cron       string              `json:"cron"`
	ProjectId  int                 `json:"projectId" gorm:"unique"`
	ServeId    int                 `json:"ServeId"`
}

func (SwaggerSync) TableName() string {
	return "biz_project_serve_swagger_sync"
}
