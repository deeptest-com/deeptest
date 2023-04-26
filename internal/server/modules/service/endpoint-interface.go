package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
)

type EndpointInterfaceService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
}

func NewEndpointInterfaceService() *EndpointInterfaceService {
	return &EndpointInterfaceService{}
}

func (s *EndpointInterfaceService) Paginate(req v1.EndpointInterfaceReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointInterfaceRepo.Paginate(req)
	return
}
