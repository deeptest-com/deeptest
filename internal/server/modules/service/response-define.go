package service

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type ResponseDefineService struct {
	ResponseDefineRepo *repo.ResponseDefineRepo `inject:""`
}

func (s *ResponseDefineService) Update(tenantId consts.TenantId, id uint, disabled bool, code string) (err error) {
	data := map[string]interface{}{
		"disabled": disabled,
		"code":     code,
	}

	err = s.ResponseDefineRepo.Update(tenantId, id, data)

	return
}
