package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type LlmToolService struct {
	LlmToolRepo *repo.LlmToolRepo `inject:""`
}

func (s *LlmToolService) List(tenantId consts.TenantId, keywords string, projectId int, ignoreDisabled bool) (ret []model.LlmTool, err error) {
	ret, err = s.LlmToolRepo.List(tenantId, keywords, projectId, ignoreDisabled)
	return
}

func (s *LlmToolService) Get(tenantId consts.TenantId, id uint) (model.LlmTool, error) {
	return s.LlmToolRepo.Get(tenantId, id)
}

func (s *LlmToolService) Save(tenantId consts.TenantId, po *model.LlmTool) (err error) {
	return s.LlmToolRepo.Save(tenantId, po)
}

func (s *LlmToolService) Delete(tenantId consts.TenantId, id uint) (err error) {
	return s.LlmToolRepo.Delete(tenantId, id)
}

func (s *LlmToolService) SetDefault(tenantId consts.TenantId, id uint) (err error) {
	return s.LlmToolRepo.SetDefault(tenantId, id)
}

func (s *LlmToolService) Disable(tenantId consts.TenantId, id uint) (err error) {
	return s.LlmToolRepo.Disable(tenantId, id)
}

func (s *LlmToolService) UpdateName(tenantId consts.TenantId, req v1.ToolLlmReq) (err error) {
	err = s.LlmToolRepo.UpdateName(tenantId, req)

	return
}
