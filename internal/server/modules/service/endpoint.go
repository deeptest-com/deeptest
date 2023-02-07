package service

import (
	"fmt"
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_domain "github.com/aaronchen2k/deeptest/pkg/domain"
	"github.com/jinzhu/copier"
)

type EndpointService struct {
	EndpointRepo *repo.EndpointRepo `inject:""`
}

func NewEndpointService() *EndpointService {
	return &EndpointService{}
}

func (s *EndpointService) Paginate(req v1.EndpointReqPaginate) (ret _domain.PageData, err error) {
	ret, err = s.EndpointRepo.Paginate(req)
	return
}

func (s *EndpointService) Save(req v1.EndpointReq) (res model.Endpoint, err error) {
	var endpoint model.Endpoint
	copier.CopyWithOption(&endpoint, req, copier.Option{DeepCopy: true})
	fmt.Println(endpoint, "++++++++++++++++")
	err = s.EndpointRepo.SaveAll(endpoint)
	return endpoint, err
}
