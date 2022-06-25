package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ScenarioProcessorService) UpdateName(req serverDomain.ScenarioProcessorReq) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *ScenarioProcessorService) Save(req serverDomain.ScenarioProcessorReq) (err error) {
	err = s.ScenarioProcessorRepo.Save(req)
	return
}
