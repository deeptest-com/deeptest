package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/core/cron"
	"github.com/deeptest-com/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
	_commUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	logUtils "github.com/deeptest-com/deeptest/pkg/lib/log"
)

type SwaggerCron struct {
	ServeRepo                *repo.ServeRepo             `inject:""`
	ProjectSettingsRepo      *repo.ProjectSettingsRepo   `inject:""`
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	ProjectCronRepo          *repo.ProjectCronRepo       `inject:""`
	Cron                     *cron.ServerCron            `inject:""`
	EndpointInterfaceService *EndpointInterfaceService   `inject:""`
	ProjectCronService       *ProjectCronService         `inject:""`
}

func (s *SwaggerCron) Run(options map[string]interface{}) (f func() error) {
	f = func() error {
		tenantId := options["tenantId"].(consts.TenantId)

		taskId := options["taskId"].(uint)
		task, err := s.GetSwaggerSyncById(tenantId, taskId)
		logUtils.Info("swagger定时任务开启：" + _commUtils.JsonEncode(task))
		if err != nil {
			logUtils.Errorf("swagger定时导入任务失败,任务ID：%v,错误原因：%v", task.ID, err.Error())
			panic(err)
		}

		projectId := options["projectId"].(uint)

		req := v1.ImportEndpointDataReq{ProjectId: projectId, ServeId: uint(task.ServeId), CategoryId: int64(task.CategoryId), OpenUrlImport: true, DriverType: convert.SWAGGER, FilePath: task.Url, DataSyncType: task.SyncType, SourceType: 1}
		err = s.EndpointInterfaceService.ImportEndpointData(tenantId, req)
		if err != nil {
			logUtils.Error("swagger定时导入任务失败，错误原因：" + err.Error())
			panic(err)
		}

		//更新实现执行时间
		logUtils.Info("swagger定时任务结束：" + _commUtils.JsonEncode(task))

		return nil
	}
	return
}

func (s *SwaggerCron) SaveSwaggerSync(tenantId consts.TenantId, req model.SwaggerSync) (id uint, err error) {
	err = s.ProjectSettingsRepo.SaveSwaggerSync(tenantId, &req)
	id = req.ID

	return
}

func (s *SwaggerCron) GetSwaggerSyncById(tenantId consts.TenantId, id uint) (data model.SwaggerSync, err error) {
	data, err = s.ProjectSettingsRepo.GetSwaggerSyncById(tenantId, id)
	return
}

func (s *SwaggerCron) CallBack(options map[string]interface{}, err error) func() {
	f := func() {
	}

	return f
}
