package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi/convert"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/jinzhu/copier"
	"strconv"
)

type ProjectSettingsService struct {
	ServeRepo                *repo.ServeRepo             `inject:""`
	ProjectSettingsRepo      *repo.ProjectSettingsRepo   `inject:""`
	EndpointRepo             *repo.EndpointRepo          `inject:""`
	EndpointInterfaceRepo    *repo.EndpointInterfaceRepo `inject:""`
	Cron                     *cron.ServerCron            `inject:""`
	EndpointInterfaceService *EndpointInterfaceService   `inject:""`
}

func (s *ProjectSettingsService) SaveSwaggerSync(tenantId consts.TenantId, req v1.SwaggerSyncReq) (data model.SwaggerSync, err error) {
	var swaggerSync model.SwaggerSync
	copier.CopyWithOption(&swaggerSync, req, copier.Option{DeepCopy: true})
	serve, _ := s.ServeRepo.GetDefault(tenantId, req.ProjectId)
	swaggerSync.ServeId = int(serve.ID)
	err = s.ProjectSettingsRepo.SaveSwaggerSync(tenantId, &swaggerSync)
	s.AddSwaggerCron(tenantId, swaggerSync)
	data, err = s.ProjectSettingsRepo.GetSwaggerSync(tenantId, req.ProjectId)
	return
}

func (s *ProjectSettingsService) SwaggerSyncDetail(tenantId consts.TenantId, projectId uint) (data model.SwaggerSync, err error) {
	return s.ProjectSettingsRepo.GetSwaggerSync(tenantId, projectId)
}

func (s *ProjectSettingsService) SwaggerSyncList(tenantId consts.TenantId) (data []model.SwaggerSync, err error) {
	return s.ProjectSettingsRepo.GetSwaggerSyncList(tenantId)
}

func (s *ProjectSettingsService) GetSwaggerSyncById(tenantId consts.TenantId, id uint) (data model.SwaggerSync, err error) {
	data, err = s.ProjectSettingsRepo.GetSwaggerSyncById(tenantId, id)
	return
}

func (s *ProjectSettingsService) AddSwaggerCron(tenantId consts.TenantId, item model.SwaggerSync) {
	name := "swaggerSync" + "_" + strconv.Itoa(int(item.ID))
	s.Cron.RemoveTask(name)

	if item.Switch == consts.SwitchOFF {
		return
	}

	taskId := item.ID
	s.Cron.AddCommonTask(name, item.Cron, func() {

		task, err := s.GetSwaggerSyncById(tenantId, taskId)
		logUtils.Info("swagger定时任务开启：" + _commUtils.JsonEncode(item))
		if err != nil {
			logUtils.Errorf("swagger定时导入任务失败,任务ID：%v,错误原因：%v", name, err.Error())
			return
		}
		if task.Switch == consts.SwitchOFF {
			logUtils.Infof("swagger定时导入关闭,任务ID:%v", name)
			return
		}
		req := v1.ImportEndpointDataReq{ProjectId: uint(task.ProjectId), ServeId: uint(task.ServeId), CategoryId: int64(task.CategoryId), OpenUrlImport: true, DriverType: convert.SWAGGER, FilePath: task.Url, DataSyncType: task.SyncType, SourceType: 1}
		err = s.EndpointInterfaceService.ImportEndpointData(tenantId, req)
		if err != nil {
			logUtils.Error("swagger定时导入任务失败，错误原因：" + err.Error())
		}

		//更新实现执行时间
		s.UpdateSwaggerSyncExecTimeById(tenantId, taskId)
		logUtils.Info("swagger定时任务结束：" + _commUtils.JsonEncode(item))
	})

}

func (s *ProjectSettingsService) UpdateSwaggerSyncExecTimeById(tenantId consts.TenantId, id uint) (err error) {
	return s.ProjectSettingsRepo.UpdateSwaggerSyncExecTimeById(tenantId, id)
}

func (s *ProjectSettingsService) GetMock(tenantId consts.TenantId, projectId uint) (data model.ProjectMockSetting, err error) {
	data, err = s.ProjectSettingsRepo.GetMock(tenantId, projectId)
	return
}
func (s *ProjectSettingsService) SaveMock(tenantId consts.TenantId, req v1.MockReq) (ret model.ProjectMockSetting, err error) {
	ret = model.ProjectMockSetting{
		BaseModel: model.BaseModel{ID: req.ID},
		Priority:  req.Priority,
		ProjectId: req.ProjectId,
	}

	err = s.ProjectSettingsRepo.SaveMock(tenantId, &ret)

	return
}
