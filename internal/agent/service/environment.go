package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/getkin/kin-openapi/openapi3"
)

type EnvironmentService struct {
	EnvironmentRepo *repo.EnvironmentRepo `inject:""`
}

func (s *EnvironmentService) Generate(doc *openapi3.T, projectId uint) (err error) {
	env, err := s.EnvironmentRepo.GetByProject(projectId)
	if err != nil {
		return
	}

	envVars, err := openapi.ConvertServersToEnvironments(doc.Servers)
	if err != nil {
		return
	}

	for _, vari := range envVars {
		po, _ := s.EnvironmentRepo.GetSameVar(vari, env.ID)

		if po.ID == 0 {
			vari.EnvironmentId = env.ID
			s.EnvironmentRepo.SaveVar(&vari)
		}
	}

	return
}
