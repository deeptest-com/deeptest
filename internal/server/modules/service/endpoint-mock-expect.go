package service

import (
	v1 "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type EndpointMockExpectService struct {
	EndpointMockExpectRepo *repo.EndpointMockExpectRepo `inject:""`
}

func (s *EndpointMockExpectService) List(endpointId uint) (res []model.EndpointMockExpect, err error) {
	res, err = s.EndpointMockExpectRepo.ListByEndpointId(endpointId)
	return
}

func (s *EndpointMockExpectService) GetDetail(expectId uint) (res model.EndpointMockExpect, err error) {
	res, err = s.EndpointMockExpectRepo.GetExpectDetail(expectId)
	return
}

func (s *EndpointMockExpectService) Save(req model.EndpointMockExpect) (expectId uint, err error) {
	expectId, err = s.EndpointMockExpectRepo.Save(req)
	return
}

func (s *EndpointMockExpectService) Copy(expectId uint) (id uint, err error) {
	expectDetail, err := s.GetDetail(expectId)
	if err != nil {
		return
	}

	s.InitExpectId(&expectDetail)

	id, err = s.Save(expectDetail)

	return
}

func (s *EndpointMockExpectService) InitExpectId(expect *model.EndpointMockExpect) {
	expect.ID = 0
	expect.ResponseBody.ID = 0

	for k, v := range expect.RequestHeaders {
		v.ID = 0
		expect.RequestHeaders[k] = v
	}

	for k, v := range expect.RequestBodies {
		v.ID = 0
		expect.RequestBodies[k] = v
	}

	for k, v := range expect.RequestQueryParams {
		v.ID = 0
		expect.RequestQueryParams[k] = v
	}

	for k, v := range expect.RequestPathParams {
		v.ID = 0
		expect.RequestPathParams[k] = v
	}

	for k, v := range expect.ResponseHeaders {
		v.ID = 0
		expect.ResponseHeaders[k] = v
	}
}

func (s *EndpointMockExpectService) DeleteById(expectId uint) (err error) {
	err = s.EndpointMockExpectRepo.DeleteById(expectId)
	return
}

func (s *EndpointMockExpectService) SaveOrder(req v1.MockExpectIdsReq) (err error) {
	err = s.EndpointMockExpectRepo.SaveOrder(req)
	return
}

func (s *EndpointMockExpectService) UpdateExpectDisabled(expectId uint, disabled bool) (err error) {
	err = s.EndpointMockExpectRepo.UpdateDisabledStatus(expectId, disabled)
	return
}
