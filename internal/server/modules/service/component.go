package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	serverConsts "github.com/aaronchen2k/deeptest/internal/server/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ComponentService struct {
	ComponentSchemaRepo         *repo.ComponentSchemaRepo         `inject:""`
	ComponentSchemaSecurityRepo *repo.ComponentSchemaSecurityRepo `inject:""`
	CategoryRepo                *repo.CategoryRepo                `inject:""`
}

func (s *ComponentService) GetSchemaByServiceId(serveId int64) (ret _domain.PageData, err error) {
	//ret, err = s.ComponentRepo.Paginate(req)
	return
}

func (s *ComponentService) SaveSchema(req v1.SchemaReq) (interface{}, error) {
	return nil, nil
}

func (s *ComponentService) SaveSchemaSecurity(req v1.SchemaReq) (interface{}, error) {
	return nil, nil
}

func (s *ComponentService) UpdateRefById(tenantId consts.TenantId, id uint) (err error) {
	category, err := s.CategoryRepo.GetByEntityId(tenantId, id, serverConsts.SchemaCategory)
	if err != nil {
		return
	}

	err = s.ComponentSchemaRepo.ChangeRef(tenantId, id, category.ID)

	return
}
