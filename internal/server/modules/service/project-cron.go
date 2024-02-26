package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/core/cron"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"strconv"
)

type ProjectCronService struct {
	ProjectCronRepo      *repo.ProjectCronRepo      `inject:""`
	CronConfigLecangRepo *repo.CronConfigLecangRepo `inject:""`
	ProjectSettingsRepo  *repo.ProjectSettingsRepo  `inject:""`
	BaseRepo             *repo.BaseRepo             `inject:""`
	LecangCronService    *LecangCronService         `inject:""`
	SwaggerCron          *SwaggerCron               `inject:""`
	ServerCron           *cron.ServerCron           `inject:""`
}

func (s *ProjectCronService) Paginate(tenantId consts.TenantId, req v1.ProjectCronReqPaginate) (ret _domain.PageData, err error) {
	return s.ProjectCronRepo.Paginate(tenantId, req)
}

func (s *ProjectCronService) Get(tenantId consts.TenantId, id uint) (ret model.ProjectCron, err error) {
	ret, err = s.ProjectCronRepo.GetById(tenantId, id)
	if err != nil {
		return
	}

	if ret.Source == consts.CronSourceLecang {
		lecangConfig, err := s.LecangCronService.Get(tenantId, ret.ConfigId)
		if err != nil {
			return ret, err
		}

		ret.LecangConfig = lecangConfig
	} else if ret.Source == consts.CronSourceSwagger {
		swaggerConfig, err := s.SwaggerCron.GetSwaggerSyncById(tenantId, ret.ConfigId)
		if err != nil {
			return ret, err
		}

		ret.SwaggerConfig = swaggerConfig
	}

	return
}

func (s *ProjectCronService) Save(tenantId consts.TenantId, req model.ProjectCron) (ret model.ProjectCron, err error) {
	s.initCron(&req)

	var configId uint
	if req.Source == consts.CronSourceLecang {
		configId, err = s.LecangCronService.Save(tenantId, req.LecangConfig)
	} else if req.Source == consts.CronSourceSwagger {
		configId, err = s.SwaggerCron.SaveSwaggerSync(tenantId, req.SwaggerConfig)
	}

	if err != nil {
		return
	}

	req.ConfigId = configId
	ret, err = s.ProjectCronRepo.Save(tenantId, req)

	return
}

func (s *ProjectCronService) initCron(req *model.ProjectCron) {
	if req.ID == 0 {
		req.LecangConfig.ID = 0
		req.SwaggerConfig.ID = 0
		req.Switch = consts.SwitchON
	}

	if req.SwaggerConfig.CategoryId == 0 {
		req.SwaggerConfig.CategoryId = -1
	}
	if req.LecangConfig.CategoryId == 0 {
		req.LecangConfig.CategoryId = -1
	}
}

func (s *ProjectCronService) Delete(tenantId consts.TenantId, id uint) (err error) {
	projectCron, err := s.ProjectCronRepo.GetById(tenantId, id)
	if err != nil {
		return
	}

	err = s.ProjectCronRepo.DeleteById(tenantId, id)
	if err != nil {
		return
	}

	if projectCron.Source == consts.CronSourceLecang {
		err = s.CronConfigLecangRepo.DeleteById(tenantId, projectCron.ConfigId)
	} else if projectCron.Source == consts.CronSourceSwagger {
		err = s.ProjectSettingsRepo.DeleteSwaggerSyncById(tenantId, projectCron.ConfigId)
	}
	if err != nil {
		return
	}

	taskName := projectCron.Source.String() + "_" + strconv.Itoa(int(projectCron.ConfigId))
	s.ServerCron.RemoveTask(taskName)
	return
}

func (s *ProjectCronService) Clone(tenantId consts.TenantId, id, userId uint) (ret model.ProjectCron, err error) {
	oldCron, err := s.Get(tenantId, id)
	if err != nil {
		return
	}

	oldCron.ID = 0
	oldCron.CreateUserId = userId
	oldCron.UpdatedAt = nil
	ret, err = s.Save(tenantId, oldCron)

	return
}

func (s *ProjectCronService) UpdateSwitchStatus(tenantId consts.TenantId, id uint, switchStatus consts.SwitchStatus) (err error) {
	return s.ProjectCronRepo.UpdateSwitchById(tenantId, id, switchStatus)
}

func (s *ProjectCronService) UpdateCronExecTimeById(tenantId consts.TenantId, configId uint, source consts.CronSource, err error) error {
	execStatus := consts.CronExecSuccess
	execErr := ""
	if err != nil {
		execStatus = consts.CronExecFail
		execErr = err.Error()
	}

	return s.ProjectCronRepo.UpdateExecResult(tenantId, configId, source, execStatus, execErr)
}
