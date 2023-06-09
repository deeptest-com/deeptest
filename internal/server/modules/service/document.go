package service

import (
	domain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
)

type DocumentService struct {
}

func (s *DocumentService) Content(req domain.DocumentReq) (res domain.DocumentReq, err error) {
	return
}

func (s *DocumentService) GetInterfaces(projectId uint, serves, endpointIds []uint) (uint, []uint, []model.Endpoint) {
	var endpoints []model.Endpoint
	if projectId != 0 {

	} else if len(serves) != 0 {

	} else if len(endpointIds) != 0 {

	}

	return projectId, serves, endpoints
}
