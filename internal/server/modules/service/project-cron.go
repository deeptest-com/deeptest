package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	third_party "github.com/aaronchen2k/deeptest/internal/server/modules/service/third-party"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type ProjectCronService struct {
	ProjectCronRepo      *repo.ProjectCronRepo          `inject:""`
	CronConfigLecangRepo *repo.CronConfigLecangRepo     `inject:""`
	LecangCronService    *third_party.LecangCronService `inject:""`
}

func (s *ProjectCronService) Paginate(req v1.ProjectCronReqPaginate) (ret _domain.PageData, err error) {
	return s.ProjectCronRepo.Paginate(req)
}

func (s *ProjectCronService) Get(id uint) (ret v1.ProjectCronReq, err error) {
	projectCron, err := s.ProjectCronRepo.GetById(id)
	if err != nil {
		return
	}

	copier.CopyWithOption(&ret, &projectCron, copier.Option{DeepCopy: true})

	if ret.Source == consts.CronSourceLecang {

		cronLecang, err := s.LecangCronService.Get(ret.ConfigId)
		if err != nil {
			return ret, err
		}

		copier.CopyWithOption(&ret.LecangReq, &cronLecang, copier.Option{DeepCopy: true})
	} else if ret.Source == consts.CronSourceSwagger {
		// TODO
	}

	return
}

func (s *ProjectCronService) Create(req v1.ProjectCronReq) (id uint, err error) {
	var configId uint

	if req.Source == consts.CronSourceLecang {
		configId, err = s.LecangCronService.Save(0, req.LecangReq)
	} else if req.Source == consts.CronSourceSwagger {
		// TODO
	}

	if err != nil {
		return
	}

	projectCron := model.ProjectCron{}
	copier.CopyWithOption(&projectCron, &req, copier.Option{DeepCopy: true})

	projectCron.ConfigId = configId
	id, err = s.ProjectCronRepo.Save(projectCron)

	return
}

func (s *ProjectCronService) Update(req v1.ProjectCronReq) (err error) {
	if req.Source == consts.CronSourceLecang {
		_, err = s.LecangCronService.Save(req.ConfigId, req.LecangReq)
	} else if req.Source == consts.CronSourceSwagger {
		// TODO
	}

	if err != nil {
		return
	}

	projectCron := model.ProjectCron{}
	copier.CopyWithOption(&projectCron, &req, copier.Option{DeepCopy: true})
	projectCron.ID = req.Id
	_, err = s.ProjectCronRepo.Save(projectCron)

	return
}

func (s *ProjectCronService) Save(req v1.ProjectCronReq) (id uint, err error) {
	var configId uint
	if req.Source == consts.CronSourceLecang {
		configId, err = s.LecangCronService.Save(req.ConfigId, req.LecangReq)
	} else if req.Source == consts.CronSourceSwagger {
		// TODO
	}

	if err != nil {
		return
	}

	projectCron := model.ProjectCron{}
	copier.CopyWithOption(&projectCron, &req, copier.Option{DeepCopy: true})
	projectCron.ID = req.Id
	projectCron.ConfigId = configId

	id, err = s.ProjectCronRepo.Save(projectCron)

	return
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
		// TODO
	}

	return
}

func (s *ProjectCronService) Clone(id, userId uint) (ret uint, err error) {
	oldCron, err := s.Get(id)
	if err != nil {
		return
	}

	oldCron.CreateUserId = userId
	ret, err = s.Create(oldCron)

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
