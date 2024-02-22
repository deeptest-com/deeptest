package service

import (
	"context"
	"errors"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	mockResponder "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/responder"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/kataras/iris/v12"
	"net/http"
	"net/url"
	"sync"
)

type MockService struct {
	responder mockResponder.Responder

	endpointRouterMap   sync.Map // maintain router for each endpoint in a map
	projectGeneratorMap sync.Map // maintain generate for each project in a map
	projectOptionsMap   sync.Map // maintain options for each project in a map

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ProjectRepo           *repo.ProjectRepo           `inject:""`

	MockAdvanceService       *MockAdvanceService       `inject:""`
	EndpointService          *EndpointService          `inject:""`
	EndpointMockParamService *EndpointMockParamService `inject:""`
	ProjectSettingsRepo      *repo.ProjectSettingsRepo `inject:""`
}

func (s *MockService) ByRequest(tenantId consts.TenantId, req *MockRequest, ctx iris.Context) (resp mockGenerator.Response, err error) {
	// load endpoint interface
	endpointInterface, _, err := s.FindEndpointInterface(tenantId, req)
	if err != nil {
		return
	}

	resp, ok := s.MockAdvanceService.ByAdvanceMock(tenantId, endpointInterface, ctx)
	if ok {
		return
	}

	// init mock generator if needed
	settings, _ := s.ProjectSettingsRepo.GetMock(tenantId, endpointInterface.ProjectId)
	req.UseExamples = settings.UseExamples

	// init and cache endpoint router if needed
	err = s.generateEndpointRouter(tenantId, endpointInterface.EndpointId)
	if err != nil {
		return
	}

	// simulate an API request
	apiRequest := http.Request{
		Method: req.EndpointMethod.String(),
		URL: &url.URL{
			Path: req.EndpointPath,
		},
	}

	// find request route
	requestRoute, _, err := s.findRequestRoute(apiRequest, endpointInterface.EndpointId)
	if err != nil {
		return
	}

	generator, _ := s.getGeneratorFromMap(endpointInterface.ProjectId, req)
	response, err := (*generator).GenerateResponse(&apiRequest, requestRoute, req.Code)

	if err != nil {
		return
	}

	resp = *response

	return
}

func (s *MockService) initMockGenerator(config mockData.Options) (ret mockGenerator.ResponseGenerator, err error) {
	options := mockData.Options{
		UseExamples: config.UseExamples,
		//NullProbability: config.NullProbability,
		//DefaultMinInt:   config.DefaultMinInt,
		//DefaultMaxInt:   config.DefaultMaxInt,
		//DefaultMinFloat: config.DefaultMinFloat,
		//DefaultMaxFloat: config.DefaultMaxFloat,
		//SuppressErrors:  config.SuppressErrors,
	}
	dataGenerator := mockData.New(options)
	ret = mockGenerator.New(dataGenerator)
	s.responder = mockResponder.New()

	return
}

func (s *MockService) generateEndpointRouter(tenantId consts.TenantId, endpointId uint) (err error) {
	// cache if need
	//endpointRouter, ok := s.getRouterFromMap(endpointId)
	//if ok && endpointRouter != nil {
	//	return
	//}

	// generate openapi spec

	endpoint := s.EndpointService.GetById(tenantId, endpointId, "v0.1.0")
	if endpoint.ID == 0 {

		return
	}

	doc3 := s.EndpointService.Yaml(tenantId, endpoint)

	/* chenqi test
	pth := "/Users/aaron/rd/project/gudi/deeptest/xdoc/openapi/openapi3/test1.json"
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc3, err = loader.LoadFromFile(pth)
	*/

	/*
		pth := "C:/Users/Lenovo/go/src/deeptest/xdoc/openapi/openapi3/test1.json"

		ctx := context.Background()
		loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}

		doc3, err := loader.LoadFromFile(pth)

		var result interface{}
		commonUtils.JsonDecode(commonUtils.JsonEncode(doc3), &result)
		respContent, _ := encoder.NewEncoder(result).Encode()

		log.Println(string(respContent))
	*/
	/*
		// fix spec issues
		doc3.Servers = nil                                                 // if not empty, will be used by s.router.FindRout() method
		doc3.Paths["/json"].Post = nil                                     // ignore post method for testing
		doc3.Info.Version = "1.0.0"                                        // cannot be empty
		desc := "描述"                                                       // cannot be empty
		doc3.Paths["/json"].Get.Responses["200"].Sample.Description = &desc // cannot be empty

		// load openapi spec from url or file
		//specificationLoader := mockLoader.New()
		//specification, err := specificationLoader.LoadFromURI(config.SpecificationURL)

		// init mock router
	*/
	loader := &openapi3.Loader{Context: context.Background(), IsExternalRefsAllowed: true}
	x := commonUtils.JsonEncode(doc3)
	//logUtils.Info(x)
	doc3, err = loader.LoadFromData([]byte(x))
	if err != nil {
		return
	}
	doc3.Servers = nil
	//pth := "/Users/aaron/rd/project/gudi/deeptest/xdoc/openapi/openapi3/test1.json"
	//ctx := context.Background()
	//loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	//doc3, err = loader.LoadFromFile(pth)
	ret, err := gorillamux.NewRouter(doc3)

	if err != nil {
		return
	}

	s.endpointRouterMap.Store(endpointId, ret)

	return
}

func (s *MockService) FindEndpointInterface(tenantId consts.TenantId, req *MockRequest) (
	endpointInterface model.EndpointInterface, paramsMap map[string]string, err error) {

	if req.EndpointInterfaceId <= 0 {
		var serve model.Serve
		serve, err = s.ServeRepo.Get(tenantId, uint(req.ServeId))
		if err != nil {
			return
		}

		endpoint, paramsMap1, err1 := s.findEndpointByPath(tenantId, serve.ID, req.EndpointPath, req.EndpointMethod)
		if err1 != nil {
			err = errors.New("not found")
			return
		}

		paramsMap = paramsMap1
		_, req.EndpointInterfaceId = s.EndpointInterfaceRepo.GetByMethod(tenantId, endpoint.ID, req.EndpointMethod)

	}

	endpointInterface, err = s.EndpointInterfaceRepo.Get(tenantId, req.EndpointInterfaceId)
	if err != nil {
		return
	}

	return
}

func (s *MockService) findEndpointByPath(tenantId consts.TenantId, serveId uint, mockPath string, method consts.HttpMethod) (
	ret model.Endpoint, paramsMap map[string]string, err error) {

	endpoints, _ := s.EndpointRepo.GetByPath(tenantId, serveId, mockPath, method)
	if len(endpoints) == 0 {
		endpoints, _ = s.EndpointRepo.ListByProjectIdAndServeId(tenantId, serveId, method)
	}

	for _, endpoint := range endpoints {
		paramsMap1, matched := s.EndpointMockParamService.MatchEndpointByMockPath(tenantId, mockPath, *endpoint)

		if matched {
			ret = *endpoint
			paramsMap = paramsMap1
			return
		}
	}

	err = errors.New("not found")
	return
}

func (s *MockService) findRequestRoute(httpRequest http.Request, endpointId uint) (requestRoute *routers.Route, pathParameters map[string]string, err error) {
	// find matched requestRoute
	endpointRouter, ok := s.getRouterFromMap(endpointId)
	if !ok {
		return
	}

	requestRoute, pathParameters, err = endpointRouter.FindRoute(&httpRequest)

	return
}

func (s *MockService) getGeneratorFromMap(projectId uint, req *MockRequest) (ret *mockGenerator.ResponseGenerator, ok bool) {
	obj, ok := s.projectGeneratorMap.Load(projectId)
	if ok {
		temp := obj.(*mockGenerator.ResponseGenerator)
		ret = temp
	}

	if obj == nil || s.isOptionsChanged(projectId, req) {
		newConfig := mockData.Options{
			UseExamples: req.UseExamples,
		}
		temp, _ := s.initMockGenerator(newConfig)
		ret = &temp

		s.projectOptionsMap.Store(projectId, newConfig)
		s.projectGeneratorMap.Store(projectId, ret)
	}

	return
}
func (s *MockService) getOptionsFromMap(projectId uint) (ret *mockData.Options, ok bool) {
	obj, ok := s.projectOptionsMap.Load(projectId)
	if ok {
		val := obj.(mockData.Options)
		ret = &val
	}

	if ret == nil {
		ret = &mockData.Options{
			UseExamples: mockData.No,
		}
		s.endpointRouterMap.Store(projectId, ret)
	}

	return
}
func (s *MockService) isOptionsChanged(projectId uint, req *MockRequest) (ret bool) {
	old, _ := s.getOptionsFromMap(projectId)

	if old.UseExamples != req.UseExamples {
		return true
	} else {
		return false
	}
}

func (s *MockService) getRouterFromMap(key uint) (ret routers.Router, ok bool) {
	obj, ok := s.endpointRouterMap.Load(key)

	if ok {
		ret = obj.(routers.Router)
	}

	return
}

type MockRequest struct {
	ProjectId int `json:"projectCode"`
	ServeId   int `json:"serveCode"`

	EndpointPath        string            `json:"endpointPath"`
	EndpointMethod      consts.HttpMethod `json:"endpointMethod"`
	EndpointInterfaceId uint              `json:"endpointInterfaceId"`

	Code        string                   `json:"code"`
	UseExamples mockData.UseExamplesEnum `json:"endpointInterfaceId"`
}

type MockResponse struct {
	StatusCode  int                    `json:"statusCode"`
	ContentType string                 `json:"contentType"`
	Data        mockGenerator.Response `json:"data"`
}
