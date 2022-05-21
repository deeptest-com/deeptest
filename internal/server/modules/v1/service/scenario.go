package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioService struct {
	ScenarioRepo *repo.ScenarioRepo `inject:""`
}

func NewScenarioService() *ScenarioService {
	return &ScenarioService{}
}

func (s *ScenarioService) Paginate(req serverDomain.ScenarioReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioRepo.Paginate(req)

	if err != nil {
		return
	}

	return
}

func (s *ScenarioService) FindById(id uint) (serverDomain.ScenarioResp, error) {
	return s.ScenarioRepo.FindById(id)
}

func (s *ScenarioService) Create(req serverDomain.ScenarioReq) (uint, error) {
	return s.ScenarioRepo.Create(req)
}

func (s *ScenarioService) Update(id uint, req serverDomain.ScenarioReq) error {
	return s.ScenarioRepo.Update(id, req)
}

func (s *ScenarioService) DeleteById(id uint) error {
	return s.ScenarioRepo.DeleteById(id)
}
