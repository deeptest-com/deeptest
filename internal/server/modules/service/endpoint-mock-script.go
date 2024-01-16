package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type EndpointMockScriptService struct {
	EndpointMockScriptRepo *repo.EndpointMockScriptRepo `inject:""`
}

func (s *EndpointMockScriptService) Get(id uint) (model.EndpointMockScript, error) {
	return s.EndpointMockScriptRepo.Get(id)
}

func (s *EndpointMockScriptService) Update(req model.EndpointMockScript) error {
	return s.EndpointMockScriptRepo.Update(req)
}

func (s *EndpointMockScriptService) Disable(endpointId uint) error {
	return s.EndpointMockScriptRepo.Disable(endpointId)
}

func (s *EndpointMockScriptService) Copy(endpointId, newEndpointId uint) (err error) {
	if newEndpointId == 0 {
		return
	}

	mockScript, err := s.Get(endpointId)
	if err != nil {
		return
	}

	mockScript.EndpointId = newEndpointId
	newMockScript, err := s.Get(newEndpointId)
	if err != nil {
		return
	}
	mockScript.ID = newMockScript.ID

	return s.Update(mockScript)
}
