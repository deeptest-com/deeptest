package service

import (
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

func (s *EndpointService) Save(req v1.EndpointReq) (res uint, err error) {
	var endpoint model.Endpoint
	copier.CopyWithOption(&endpoint, req, copier.Option{DeepCopy: true})
	//fmt.Println(_commUtils.JsonEncode(endpoint), "++++++", _commUtils.JsonEncode(req))
	err = s.EndpointRepo.SaveAll(&endpoint)
	return endpoint.ID, err
}

func (s *EndpointService) GetById(id uint) (res model.Endpoint) {
	res, _ = s.EndpointRepo.GetAll(id)
	return
}

func (s *EndpointService) DeleteById(id uint) (err error) {
	err = s.EndpointRepo.DeleteById(id)
	return
}

func (s *EndpointService) DisableById(id uint) (err error) {
	err = s.EndpointRepo.DisableById(id)
	return
}
