package service

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type MockAdvanceService struct {
	EndpointService     *EndpointService          `inject:""`
	ProjectSettingsRepo *repo.ProjectSettingsRepo `inject:""`
	EndpointRepo        *repo.EndpointRepo        `inject:""`
}

func (s *MockAdvanceService) IsAdvanceMockDisabled(endpointId uint) (ret bool) {
	endpoint, _ := s.EndpointRepo.Get(endpointId)

	ret = endpoint.AdvancedMockDisabled
	return
}
