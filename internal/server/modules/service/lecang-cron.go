package service

import (
	"errors"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
)

type LecangCronService struct {
	CronConfigLecangRepo  *repo.CronConfigLecangRepo `inject:""`
	BaseRepo              *repo.BaseRepo             `inject:""`
	ProjectCronService    *ProjectCronService        `inject:""`
	ThirdPartySyncService *ThirdPartySyncService     `inject:""`
}

func (s *LecangCronService) Run(options map[string]interface{}) (f func() error) {
	f = func() error {
		taskId, ok := options["taskId"].(uint)
		if !ok {
			return errors.New("taskId is not existed")
		}
		task, err := s.Get(taskId)
		logUtils.Info("lecang定时任务开启：" + _commUtils.JsonEncode(task))
		if err != nil {
			logUtils.Errorf("lecang定时导入任务失败,任务ID：%v,错误原因：%v", task.ID, err.Error())
			return err
		}

		projectId, ok := options["projectId"].(uint)
		if !ok {
			return errors.New("projectId is not existed")
		}

		err = s.ThirdPartySyncService.ImportEndpoint(projectId, task)

		return err
	}

	return
}

func (s *LecangCronService) CallBack(options map[string]interface{}, err error) func() {
	f := func() {
		taskId, ok := options["taskId"].(uint)
		if !ok {
			return
		}
		s.ProjectCronService.UpdateCronExecTimeById(taskId, consts.CronSourceSwagger, err)
	}

	return f
}

func (s *LecangCronService) Get(id uint) (ret model.CronConfigLecang, err error) {
	return s.CronConfigLecangRepo.GetById(id)
}

func (s *LecangCronService) Save(req model.CronConfigLecang) (ret uint, err error) {
	ret, err = s.CronConfigLecangRepo.Save(req)

	return
}
