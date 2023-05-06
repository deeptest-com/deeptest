package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo2 "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"github.com/aaronchen2k/deeptest/pkg/domain"
)

type ScenarioService struct {
	ScenarioRepo     *repo2.ScenarioRepo     `inject:""`
	ScenarioNodeRepo *repo2.ScenarioNodeRepo `inject:""`
}

func NewScenarioService() *ScenarioService {
	return &ScenarioService{}
}

func (s *ScenarioService) ListByProject(serveId int) (pos []model.Scenario, err error) {
	pos, err = s.ScenarioRepo.ListByProject(serveId)
	return
}

func (s *ScenarioService) Paginate(req v1.ScenarioReqPaginate, projectId int) (ret _domain.PageData, err error) {
	ret, err = s.ScenarioRepo.Paginate(req, projectId)

	if err != nil {
		return
	}

	return
}

func (s *ScenarioService) GetById(id uint) (model.Scenario, error) {
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
