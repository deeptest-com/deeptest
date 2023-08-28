package service

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	mockGenerator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	mockResponder "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/responder"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/aaronchen2k/deeptest/internal/server/modules/repo"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	encoder "github.com/zwgblue/yaml-encoder"
	"log"
	"net/http"
	"net/url"
)

type MockService struct {
	IsInit            bool
	endpointRouterMap map[uint]routers.Router // maintain router for each endpoint in a map

	generator mockGenerator.ResponseGenerator
	responder mockResponder.Responder

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ProjectRepo           *repo.ProjectRepo           `inject:""`

	EndpointService *EndpointService `inject:""`
}

func (s *MockService) ByRequest(req *MockRequest, ctx iris.Context) (resp *MockResponse, err error) {
	// init mock generator
	if !s.IsInit {
		s.endpointRouterMap = map[uint]routers.Router{}

		options := mockData.Options{
			//UseExamples:     config.UseExamples,
			//NullProbability: config.NullProbability,
			//DefaultMinInt:   config.DefaultMinInt,
			//DefaultMaxInt:   config.DefaultMaxInt,
			//DefaultMinFloat: config.DefaultMinFloat,
			//DefaultMaxFloat: config.DefaultMaxFloat,
			//SuppressErrors:  config.SuppressErrors,
		}
		dataGenerator := mockData.New(options)
		s.generator = mockGenerator.New(dataGenerator)
		s.responder = mockResponder.New()

		s.IsInit = true
	}

	// load endpoint interface
	endpointInterface, err := s.GetEndpointInterface(req)
	if err != nil {
		return
	}

	// generate openapi spec
	endpoint, err := s.EndpointRepo.GetAll(endpointInterface.EndpointId, "v0.1.0")
	spec := s.EndpointService.Yaml(endpoint)
	doc3 := spec.(*openapi3.T)

	var result interface{}
	commonUtils.JsonDecode(commonUtils.JsonEncode(doc3), &result)
	respContent, _ := encoder.NewEncoder(result).Encode()

	log.Println(string(respContent))

	// fix spec issues
	doc3.Servers = nil                                                 // if not empty, will be used by s.router.FindRout() method
	doc3.Paths["/json"].Post = nil                                     // just ignore for testing
	doc3.Info.Version = "1.0.0"                                        // cannot be empty
	desc := "描述"                                                       // cannot be empty
	doc3.Paths["/json"].Get.Responses["200"].Value.Description = &desc // cannot be empty

	// load openapi spec from url or file
	//specificationLoader := mockLoader.New()
	//specification, err := specificationLoader.LoadFromURI(config.SpecificationURL)

	// init mock router
	s.endpointRouterMap[endpointInterface.EndpointId], err = legacy.NewRouter(doc3)
	if err != nil {
		return
	}

	// find matched route
	httpRequest := http.Request{
		Method: req.EndpointMethod.String(),
		URL: &url.URL{
			Path: "/" + req.EndpointPath,
		},
	}

	endpointRouter := s.endpointRouterMap[endpointInterface.EndpointId]
	route, pathParameters, err := endpointRouter.FindRoute(&httpRequest)
	if err != nil {
		return
	}

	// validate request
	routingValidation := &openapi3filter.RequestValidationInput{
		Request:    &httpRequest,
		PathParams: pathParameters,
		Route:      route,
		Options: &openapi3filter.Options{
			ExcludeRequestBody: true,
		},
	}
	err = openapi3filter.ValidateRequest(ctx, routingValidation)
	var requestError *openapi3filter.RequestError
	if errors.As(err, &requestError) {
		return
	}

	// generate response
	response, err := s.generator.GenerateResponse(&httpRequest, route)
	if err != nil {
		return
	}

	log.Println(response)

	return
}

func (s *MockService) GetEndpointInterface(req *MockRequest) (ret model.EndpointInterface, err error) {
	if req.EndpointInterfaceId <= 0 {
		project, _ := s.ProjectRepo.GetByCode(req.ProjectCode)
		serve, _ := s.ServeRepo.GetByCode(project.ID, req.ServeCode)
		endpoint, _ := s.EndpointRepo.GetByPath(serve.ID, req.EndpointPath)

		_, req.EndpointInterfaceId = s.EndpointInterfaceRepo.GetByMethod(endpoint.ID, req.EndpointMethod)
	}

	ret, err = s.EndpointInterfaceRepo.Get(req.EndpointInterfaceId)

	return
}

type MockRequest struct {
	ProjectCode string `json:"projectCode"`
	ServeCode   string `json:"serveCode"`

	EndpointPath   string            `json:"endpointPath"`
	EndpointMethod consts.HttpMethod `json:"endpointMethod"`

	EndpointInterfaceId uint `json:"endpointInterfaceId"`
}

type MockResponse struct {
	StatusCode  int         `json:"statusCode"`
	ContentType string      `json:"contentType"`
	Data        interface{} `json:"data"`
}
