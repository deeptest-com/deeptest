package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type SaasService struct {
	SassRepo *repo.SaasRepo `inject:""`
}

func (s *SaasService) GetUserList(tenantId consts.TenantId) (data interface{}, err error) {
	data, err = s.SassRepo.GetUserList(tenantId)
	return
}
