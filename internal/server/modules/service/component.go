package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type ComponentService struct {
	ComponentSchemaRepo         *repo.ComponentSchemaRepo         `inject:""`
	ComponentSchemaSecurityRepo *repo.ComponentSchemaSecurityRepo `inject:""`
}

func NewComponentService() *ComponentService {
	return &ComponentService{}
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
