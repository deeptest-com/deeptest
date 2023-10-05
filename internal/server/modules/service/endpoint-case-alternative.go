package service

import (
	"context"
	serverDomain "github.com/aaronchen2k/deeptest/cmd/server/v1/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	"github.com/aaronchen2k/deeptest/internal/pkg/helper/cases"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	_stringUtils "github.com/aaronchen2k/deeptest/pkg/lib/string"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kataras/iris/v12"
	"log"
)

type EndpointCaseAlternativeService struct {
	EndpointCaseAlternativeRepo *repo.EndpointCaseAlternativeRepo `inject:""`
	EndpointCaseRepo            *repo.EndpointCaseRepo            `inject:""`

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	ServeServerRepo       *repo.ServeServerRepo       `inject:""`
	DebugInterfaceRepo    *repo.DebugInterfaceRepo    `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	PreConditionRepo      *repo.PreConditionRepo      `inject:""`
	PostConditionRepo     *repo.PostConditionRepo     `inject:""`
	CategoryRepo          *repo.CategoryRepo          `inject:""`

	EndpointService          *EndpointService          `inject:""`
	DebugInterfaceService    *DebugInterfaceService    `inject:""`
	EndpointMockParamService *EndpointMockParamService `inject:""`
}

func (s *EndpointCaseAlternativeService) LoadAlternative(baseId uint) (
	root casesHelper.AlternativeCase, err error) {

	root.Title = "备选用例"
	root.Category = consts.AlternativeCaseRoot
	root.Key = _stringUtils.Uuid()
	root.Slots = iris.Map{"icon": "icon"}
	root.IsDir = true

	casePo, _ := s.EndpointCaseRepo.Get(baseId)

	//_, endpointInterfaceId := s.EndpointInterfaceRepo.GetByMethod(req.EndpointId, req.Method)
	//if endpointInterfaceId == 0 {
	//	return
	//}
	//
	//endpointInterface, _ := s.EndpointInterfaceRepo.Get(endpointInterfaceId)
	//endpoint, err := s.EndpointRepo.GetWithInterface(endpointInterface.EndpointId, "v0.1.0")
	//
	//// get spec
	//doc3 := s.EndpointService.Yaml(endpoint)

	pth := "/Users/aaron/rd/project/gudi/deeptest/xdoc/openapi/openapi3/test2.yaml"
	loader := &openapi3.Loader{Context: context.Background(), IsExternalRefsAllowed: true}
	doc3, err := loader.LoadFromFile(pth)
	if err != nil || doc3 == nil {
		return
	}
	apiPathItem, _ := casesHelper.GetApiPathItem(doc3)

	apiOperation, err := casesHelper.GetApiOperation(casePo.Method, apiPathItem)
	if err != nil || apiOperation == nil {
		return
	}

	root.Children = append(root.Children, casesHelper.LoadForQueryParams(apiOperation.Parameters))
	root.Children = append(root.Children, casesHelper.LoadForPathParams(apiOperation.Parameters))
	root.Children = append(root.Children, casesHelper.LoadForHeaders(apiOperation.Parameters))
	root.Children = append(root.Children, casesHelper.LoadForBody(apiOperation.RequestBody))

	return
}

func (s *EndpointCaseAlternativeService) LoadAlternativeSaved(caseId uint) (
	ret map[string]model.EndpointCaseAlternative, err error) {

	pos, err := s.EndpointCaseAlternativeRepo.List(caseId)

	for _, po := range pos {
		ret[po.Path] = po
	}

	return
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
	apiPathItem, _ := casesHelper.GetApiPathItem(doc3)

	for _, interf := range endpoint.Interfaces {
		basicDebugData, err1 := s.getBaseRequest(interf)
		if err1 != nil {
			continue
		}
		log.Println(basicDebugData)

		apiOperation, err1 := casesHelper.GetApiOperation(interf.Method, apiPathItem)
		if err1 != nil {
			continue
		}
		log.Println(apiOperation)

		//alternativeCases, err1 := casesHelper.GenerateAlternativeCase(basicDebugData, apiOperation)
		//if err1 != nil {
		//	continue
		//}
		//log.Println(alternativeCases)
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
