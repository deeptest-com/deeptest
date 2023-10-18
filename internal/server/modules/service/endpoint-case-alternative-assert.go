package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	model "github.com/aaronchen2k/deeptest/internal/server/modules/model"
	repo "github.com/aaronchen2k/deeptest/internal/server/modules/repo"
)

type EndpointCaseAlternativeAssertService struct {
	EndpointCaseAlternativeAssertRepo *repo.EndpointCaseAlternativeAssertRepo `inject:""`
}

func (s *EndpointCaseAlternativeAssertService) List(alternativeCaseId uint) (asserts []model.EndpointCaseAlternativeAssert, err error) {
	asserts, err = s.EndpointCaseAlternativeAssertRepo.List(alternativeCaseId)

	return
}

func (s *EndpointCaseAlternativeAssertService) Get(id uint) (condition model.EndpointCaseAlternativeAssert, err error) {
	condition, err = s.EndpointCaseAlternativeAssertRepo.Get(id)

	return
}

func (s *EndpointCaseAlternativeAssertService) Save(condition *model.EndpointCaseAlternativeAssert) (err error) {
	err = s.EndpointCaseAlternativeAssertRepo.Save(condition)

	return
}

func (s *EndpointCaseAlternativeAssertService) Delete(reqId uint) (err error) {
	err = s.EndpointCaseAlternativeAssertRepo.Delete(reqId)

	return
}

func (s *EndpointCaseAlternativeAssertService) Disable(reqId uint) (err error) {
	err = s.EndpointCaseAlternativeAssertRepo.Disable(reqId)

	return
}

func (s *EndpointCaseAlternativeAssertService) Move(req serverDomain.ConditionMoveReq) (err error) {
	err = s.EndpointCaseAlternativeAssertRepo.UpdateOrders(req)

	return
}
