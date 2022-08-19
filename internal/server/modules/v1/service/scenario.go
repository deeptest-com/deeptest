package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioService struct {
	ScenarioRepo     *repo.ScenarioRepo     `inject:""`
	ScenarioNodeRepo *repo.ScenarioNodeRepo `inject:""`
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

func (s *ScenarioService) FindById(id uint) (model.Scenario, error) {
	return s.ScenarioRepo.Get(id)
}

func (s *ScenarioService) Create(req model.Scenario) (po model.Scenario, bizErr *_domain.BizErr) {
	po, bizErr = s.ScenarioRepo.Create(req)

	s.ScenarioNodeRepo.CreateDefault(po.ID)

	return
}

func (s *ScenarioService) Update(req model.Scenario) error {
	return s.ScenarioRepo.Update(req)
}

func (s *ScenarioService) DeleteById(id uint) error {
	return s.ScenarioRepo.DeleteById(id)
}
