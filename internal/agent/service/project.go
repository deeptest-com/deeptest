package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/internal/server/modules/service"
)

type ProjectService struct {
	ProjectRepo        *repo.ProjectRepo           `inject:""`
	EnvironmentService *service.EnvironmentService `inject:""`
}

func (s *ProjectService) CreateOrGetBySpec(pth, url string) (project model.Project, err error) {
	spec := ""
	if pth != "" {
		spec = pth
	} else {
		spec = url
	}

	project, _ = s.ProjectRepo.GetBySpec(spec)

	if project.ID == 0 {
		project.Spec = spec
		s.ProjectRepo.Save(&project)

		env := model.Environment{
			Name: "默认环境",
		}
		s.EnvironmentService.Create(&env, project.ID)
	}

	return
}
