package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioService struct {
	ScenarioRepo          *repo.ScenarioRepo          `inject:""`
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
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

func (s *ScenarioService) FindById(id uint) (model.TestScenario, error) {
	return s.ScenarioRepo.Get(id)
}

func (s *ScenarioService) Create(req model.TestScenario) (po model.TestScenario, bizErr *_domain.BizErr) {
	po, bizErr = s.ScenarioRepo.Create(req)

	s.ScenarioProcessorRepo.CreateDefault(po.ID)

	return
}

func (s *ScenarioService) Update(req model.TestScenario) error {
	return s.ScenarioRepo.Update(req)
}

func (s *ScenarioService) DeleteById(id uint) error {
	return s.ScenarioRepo.DeleteById(id)
}
