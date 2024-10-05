package service

import (
	v1 "github.com/deeptest-com/deeptest/cmd/server/v1/domain"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	jslibHelper "github.com/deeptest-com/deeptest/internal/pkg/helper/jslib"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/deeptest-com/deeptest/internal/server/modules/repo"
)

type JslibService struct {
	JslibRepo *repo.JslibRepo `inject:""`
}

func (s *JslibService) List(tenantId consts.TenantId, keywords string, projectId int) (ret []model.Jslib, err error) {
	ret, err = s.JslibRepo.List(tenantId, keywords, projectId, false)
	return
}

func (s *JslibService) Get(tenantId consts.TenantId, id uint) (model.Jslib, error) {
	return s.JslibRepo.Get(tenantId, id)
}

func (s *JslibService) Save(tenantId consts.TenantId, req *model.Jslib) (err error) {
	err = s.JslibRepo.Save(tenantId, req)
	if err != nil {
		return
	}

	jslibHelper.InitJslibCache(tenantId)

	return
}

func (s *JslibService) UpdateName(tenantId consts.TenantId, req v1.JslibReq) (err error) {
	err = s.JslibRepo.UpdateName(tenantId, req)

	jslibHelper.InitJslibCache(tenantId)

	return
}

func (s *JslibService) Delete(tenantId consts.TenantId, id uint) (err error) {
	return s.JslibRepo.Delete(tenantId, id)
}

func (s *JslibService) Disable(tenantId consts.TenantId, id uint) (err error) {
	return s.JslibRepo.Disable(tenantId, id)
}
