package serverDomain

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type SaveProjectCronReq struct {
	Id         uint                `json:"id"`
	Name       string              `json:"name"`
	Cron       string              `json:"cron"`
	ProjectId  uint                `json:"projectId"`
	Source     consts.CronSource   `json:"source"`
	Switch     consts.SwitchStatus `json:"switch"`
	SwaggerReq SwaggerSyncReq      `json:"swaggerReq"`
	LecangReq  LecangCronReq       `json:"lecangReq"`
}

type SwaggerCronReq struct {
}

type LecangCronReq struct {
}
