package service

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type SaasService struct {
	SassRepo *repo.SaasRepo `inject:""`
}

func (s *SaasService) GetUserList(tenantId consts.TenantId) (data interface{}, err error) {
	data, err = s.SassRepo.GetUserList(tenantId)
	return
}
