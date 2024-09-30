package service

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type EndpointMockScriptService struct {
	EndpointMockScriptRepo *repo.EndpointMockScriptRepo `inject:""`
}

func (s *EndpointMockScriptService) Get(tenantId consts.TenantId, id uint) (model.EndpointMockScript, error) {
	return s.EndpointMockScriptRepo.Get(tenantId, id)
}

func (s *EndpointMockScriptService) Update(tenantId consts.TenantId, req model.EndpointMockScript) error {
	return s.EndpointMockScriptRepo.Update(tenantId, req)
}

func (s *EndpointMockScriptService) Disable(tenantId consts.TenantId, endpointId uint) error {
	return s.EndpointMockScriptRepo.Disable(tenantId, endpointId)
}

func (s *EndpointMockScriptService) Copy(tenantId consts.TenantId, endpointId, newEndpointId uint) (err error) {
	if newEndpointId == 0 {
		return
	}

	mockScript, err := s.Get(tenantId, endpointId)
	if err != nil {
		return
	}

	mockScript.EndpointId = newEndpointId
	newMockScript, err := s.Get(tenantId, newEndpointId)
	if err != nil {
		return
	}
	mockScript.ID = newMockScript.ID

	return s.Update(tenantId, mockScript)
}
