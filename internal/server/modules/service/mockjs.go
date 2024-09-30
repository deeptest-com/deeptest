package service

import (
	serverDomain "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	repo "github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type MockJsService struct {
	MockJsRepo *repo.MockJsRepo `inject:""`
}

func (s *MockJsService) ListExpressions(tenantId consts.TenantId) (pos []serverDomain.MockJsExpression, err error) {
	pos, err = s.MockJsRepo.ListExpressions(tenantId)

	return
}
