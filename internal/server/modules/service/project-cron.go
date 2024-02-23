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

func (s *ProjectCronService) Paginate(req v1.ProjectCronReqPaginate) (ret _domain.PageData, err error) {
	return s.ProjectCronRepo.Paginate(req)
}

func (s *ProjectCronService) Get(id uint) (ret model.ProjectCron, err error) {
	ret, err = s.ProjectCronRepo.GetById(id)
	if err != nil {
		return
	}

	if ret.Source == consts.CronSourceLecang {
		lecangConfig, err := s.LecangCronService.Get(ret.ConfigId)
		if err != nil {
			return ret, err
		}

		ret.LecangConfig = lecangConfig
	} else if ret.Source == consts.CronSourceSwagger {
		swaggerConfig, err := s.SwaggerCron.GetSwaggerSyncById(ret.ConfigId)
		if err != nil {
			return ret, err
		}

		ret.SwaggerConfig = swaggerConfig
	}

	return
}

func (s *ProjectCronService) Save(req model.ProjectCron) (ret model.ProjectCron, err error) {
	s.initCron(&req)

	var configId uint
	if req.Source == consts.CronSourceLecang {
		configId, err = s.LecangCronService.Save(req.LecangConfig)
	} else if req.Source == consts.CronSourceSwagger {
		configId, err = s.SwaggerCron.SaveSwaggerSync(req.SwaggerConfig)
	}

	if err != nil {
		return
	}

	req.ConfigId = configId
	ret, err = s.ProjectCronRepo.Save(req)

	return
}

func (s *ProjectCronService) initCron(req *model.ProjectCron) {
	if req.ID == 0 {
		req.LecangConfig.ID = 0
		req.SwaggerConfig.ID = 0
		req.Switch = consts.SwitchON
	}
}

func (s *ProjectCronService) Delete(id uint) (err error) {
	projectCron, err := s.ProjectCronRepo.GetById(id)
	if err != nil {
		return
	}

	err = s.ProjectCronRepo.DeleteById(id)
	if err != nil {
		return
	}

	if projectCron.Source == consts.CronSourceLecang {
		err = s.CronConfigLecangRepo.DeleteById(projectCron.ConfigId)
	} else if projectCron.Source == consts.CronSourceSwagger {
		err = s.ProjectSettingsRepo.DeleteSwaggerSyncById(projectCron.ConfigId)
	}
	if err != nil {
		return
	}

	taskName := projectCron.Source.String() + "_" + strconv.Itoa(int(projectCron.ConfigId))
	s.ServerCron.RemoveTask(taskName)
	return
}

func (s *ProjectCronService) Clone(id, userId uint) (ret model.ProjectCron, err error) {
	oldCron, err := s.Get(id)
	if err != nil {
		return
	}

	oldCron.ID = 0
	oldCron.CreateUserId = userId
	ret, err = s.Save(oldCron)

	return
}

func (s *ProjectCronService) UpdateSwitchStatus(id uint, switchStatus consts.SwitchStatus) (err error) {
	return s.ProjectCronRepo.UpdateSwitchById(id, switchStatus)
}

func (s *ProjectCronService) UpdateCronExecTimeById(configId uint, source consts.CronSource, err error) error {
	execStatus := consts.CronExecSuccess
	execErr := ""
	if err != nil {
		execStatus = consts.CronExecFail
		execErr = err.Error()
	}

	return s.ProjectCronRepo.UpdateExecResult(configId, source, execStatus, execErr)
}

//func (s *ProjectCronService) AddCronItem(cronConfig model.ProjectCron) (err error) {
//	options := make(map[string]interface{})
//	options["projectId"] = cronConfig.ProjectId
//	options["taskId"] = cronConfig.ConfigId
//
//	proxy := task.NewProxy(string(cronConfig.Source), cronConfig.Cron)
//	err = proxy.Add(options)
//
//	return
//}
//
//func (s *ProjectCronService) BatchAddCron() (err error) {
//	cronList, err := s.ProjectCronRepo.ListAllCron()
//	if err != nil {
//		return
//	}
//
//	for _, cron := range cronList {
//		err = s.AddCronItem(cron)
//		if err != nil {
//			logUtils.Errorf("AddCronItem fail, cronItem:%+v, err:%+v", cron, err)
//		}
//	}
//
//	return
//}
