package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type LecangCronService struct {
	CronConfigLecangRepo  *repo.CronConfigLecangRepo `inject:""`
	BaseRepo              *repo.BaseRepo             `inject:""`
	ProjectCronRepo       *repo.ProjectCronRepo      `inject:""`
	ProjectCronService    *ProjectCronService        `inject:""`
	ThirdPartySyncService *ThirdPartySyncService     `inject:""`
}

func (s *LecangCronService) Run(options map[string]interface{}) (f func() error) {
	f = func() error {
		tenantId := options["tenantId"].(consts.TenantId)

		taskId := options["taskId"].(uint)
		task, err := s.Get(tenantId, taskId)
		logUtils.Info("lecang定时任务开启：" + _commUtils.JsonEncode(task))
		if err != nil {
			logUtils.Errorf("lecang定时导入任务失败,任务ID：%+v,错误原因：%+v", task.ID, err.Error())
			panic(err)
		}

		projectId := options["projectId"].(uint)

		err = s.ThirdPartySyncService.ImportEndpoint(tenantId, projectId, task)
		if err != nil {
			logUtils.Errorf("lecang定时导入任务失败，任务ID:%+v, 错误原因：%+v", task.ID, err.Error())
			panic(err)
		}

		return err
	}

	return
}

func (s *LecangCronService) CallBack(options map[string]interface{}, err error) func() {
	f := func() {
		taskId, ok := options["taskId"].(string)
		if !ok {
			return
		}

		tenantId, ok := options["tenantId"].(consts.TenantId)
		if !ok {
			return
		}
		s.ProjectCronService.UpdateCronExecTimeById(tenantId, taskId, consts.CronSourceLecang, err)
	}

	return f
}

func (s *LecangCronService) Get(tenantId consts.TenantId, id uint) (ret model.CronConfigLecang, err error) {
	return s.CronConfigLecangRepo.GetById(tenantId, id)
}

func (s *LecangCronService) Save(tenantId consts.TenantId, req model.CronConfigLecang) (ret uint, err error) {
	ret, err = s.CronConfigLecangRepo.Save(tenantId, req)

	return
}
