package service

import (
	"errors"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
)

type SwaggerCron struct {
	ServeRepo                *repo.ServeRepo             `inject:""`
	ProjectSettingsRepo      *repo.ProjectSettingsRepo   `inject:""`
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	Cron                     *cron.ServerCron            `inject:""`
	EndpointInterfaceService *EndpointInterfaceService   `inject:""`
	ProjectSettingsService   *ProjectSettingsService     `inject:""`
}

func (s *SwaggerCron) Run(option map[string]interface{}) (f func() error) {
	f = func() error {
		taskId, ok := option["taskId"].(uint)
		if !ok {
			return errors.New("taskId is not existed")
		}
		task, err := s.GetSwaggerSyncById(taskId)
		logUtils.Info("swagger定时任务开启：" + _commUtils.JsonEncode(task))
		if err != nil {
			logUtils.Errorf("swagger定时导入任务失败,任务ID：%v,错误原因：%v", task.ID, err.Error())
			return err
		}
		if task.Switch == consts.SwitchOFF {
			logUtils.Infof("swagger定时导入关闭,任务ID:%v", task.ID)
			return errors.New("task is off")
		}
		req := v1.ImportEndpointDataReq{ProjectId: uint(task.ProjectId), ServeId: uint(task.ServeId), CategoryId: int64(task.CategoryId), OpenUrlImport: true, DriverType: convert.SWAGGER, FilePath: task.Url, DataSyncType: task.SyncType, SourceType: 1}
		err = s.EndpointInterfaceService.ImportEndpointData(req)
		if err != nil {
			logUtils.Error("swagger定时导入任务失败，错误原因：" + err.Error())
		}

		//更新实现执行时间
		//s.UpdateSwaggerSyncExecTimeById(taskId)
		logUtils.Info("swagger定时任务结束：" + _commUtils.JsonEncode(task))

		return nil
	}
	return
}

func (s *SwaggerCron) SaveSwaggerSync(req v1.SwaggerSyncReq) (data model.SwaggerSync, err error) {
	var swaggerSync model.SwaggerSync
	copier.CopyWithOption(&swaggerSync, req, copier.Option{DeepCopy: true})
	serve, _ := s.ServeRepo.GetDefault(req.ProjectId)
	swaggerSync.ServeId = int(serve.ID)
	err = s.ProjectSettingsRepo.SaveSwaggerSync(&swaggerSync)
	//s.AddSwaggerCron(swaggerSync)
	//任务
	data, err = s.ProjectSettingsRepo.GetSwaggerSync(req.ProjectId)
	return
}

func (s *SwaggerCron) GetSwaggerSyncById(id uint) (data model.SwaggerSync, err error) {
	data, err = s.ProjectSettingsRepo.GetSwaggerSyncById(id)
	return
}

func (s *SwaggerCron) CallBack(option map[string]interface{}, err error) func() {
	f := func() {
		taskId, ok := option["taskId"].(uint)
		if !ok {
			return
		}
		s.ProjectSettingsService.UpdateCronExecTimeById(taskId, consts.CronSourceSwagger, err)
	}

	return f
}
