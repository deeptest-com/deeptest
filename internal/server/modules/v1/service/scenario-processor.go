package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverDomain "github.com/aaronchen2k/deeptest/internal/server/modules/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/v1/repo"
)

type ScenarioProcessorService struct {
	ScenarioProcessorRepo *repo.ScenarioProcessorRepo `inject:""`
}

func (s *ScenarioProcessorService) Get(id int) (ret interface{}, err error) {
	processor, _ := s.ScenarioProcessorRepo.Get(uint(id))

	if processor.EntityCategory == consts.ProcessorInterface {
		ret, _ = s.ScenarioProcessorRepo.GetInterface(uint(id), processor)

	} else if processor.EntityCategory == consts.ProcessorGroup {
		ret, _ = s.ScenarioProcessorRepo.GetGroup(uint(id), processor)

	} else if processor.EntityCategory == consts.ProcessorLogic {
		ret, _ = s.ScenarioProcessorRepo.GetLogic(uint(id), processor)

	}

	return
}

func (s *ScenarioProcessorService) UpdateName(req serverDomain.ScenarioNodeReq) (err error) {
	err = s.ScenarioProcessorRepo.UpdateName(req.Id, req.Name)
	return
}

func (s *ScenarioProcessorService) SaveLogic(req model.ProcessorLogic) (err error) {
	err = s.ScenarioProcessorRepo.Save(req)
	return
}
