package service

import (
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	"log"
)

type EndpointCaseAlternativeService struct {
	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	PreConditionRepo      *repo.PreConditionRepo      `inject:""`
	PostConditionRepo     *repo.PostConditionRepo     `inject:""`
	CategoryRepo          *repo.CategoryRepo          `inject:""`

	EndpointService       *EndpointService       `inject:""`
	DebugInterfaceService *DebugInterfaceService `inject:""`
}

func (s *EndpointCaseAlternativeService) GenerateFromSpec(req serverDomain.EndpointCaseAlternativeGenerateReq) (err error) {
	endpointInterfaceId := req.EndpointInterfaceId
	if endpointInterfaceId == 0 {
		return
	}

	endpointInterface, _ := s.EndpointInterfaceRepo.Get(uint(endpointInterfaceId))
	endpoint, err := s.EndpointRepo.GetWithInterface(endpointInterface.EndpointId, "v0.1.0")

	// get spec
	spec := s.EndpointService.Yaml(endpoint)
	doc3 := spec
	apiPathItem, _ := cases.GetApiPathItem(doc3)

	for _, interf := range endpoint.Interfaces {
		basicDebugData, err1 := s.getBaseRequest(interf)
		if err1 != nil {
			continue
		}
		log.Println(basicDebugData)

		apiOperation, err1 := cases.GetApiOperation(interf.Method, apiPathItem)
		if err1 != nil {
			continue
		}
		log.Println(apiOperation)

		alternativeCases, err1 := cases.GenerateAlternativeCase(basicDebugData, apiOperation)
		if err1 != nil {
			continue
		}
		log.Println(alternativeCases)
	}

	return
}

func (s *EndpointCaseAlternativeService) getBaseRequest(endpointInterface model.EndpointInterface) (debugData domain.DebugData, err error) {
	info := domain.DebugInfo{
		DebugInterfaceId:    endpointInterface.DebugInterfaceId,
		EndpointInterfaceId: endpointInterface.ID,
	}
	debugData, err = s.DebugInterfaceService.Load(info)

	return
}
