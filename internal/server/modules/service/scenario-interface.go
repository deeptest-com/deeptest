package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type ScenarioInterfaceService struct {
	ScenarioInterfaceRepo *repo.ScenarioInterfaceRepo `inject:""`
}

func NewScenarioInterfaceService() *ScenarioInterfaceService {
	return &ScenarioInterfaceService{}
}

func (s *ScenarioInterfaceService) GetById(id uint) (model.ProcessorInterface, error) {
	return s.ScenarioInterfaceRepo.GetDetail(id)
}

func (s *ScenarioInterfaceService) ListInvocation(id uint) (invocations []model.ProcessorInvocation, err error) {
	return s.ScenarioInterfaceRepo.ListInvocation(id)
}
