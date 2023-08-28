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
	responseGenerator mockGenerator.ResponseGenerator
	router            routers.Router
	responder         mockResponder.Responder

	EndpointInterfaceRepo *repo.EndpointInterfaceRepo `inject:""`
	EndpointRepo          *repo.EndpointRepo          `inject:""`
	ServeRepo             *repo.ServeRepo             `inject:""`
	ProjectRepo           *repo.ProjectRepo           `inject:""`

	EndpointService *EndpointService `inject:""`
}

func (s *MockService) ByRequest(req *MockRequest, ctx iris.Context) (resp *MockResponse, err error) {
	endpointInterface, err := s.GetEndpointInterface(req, req.EndpointInterfaceId)
	if err != nil {
		return
	}

	resp, err = s.ByEndpointInterface(req, endpointInterface, ctx)

	return
}

func (s *MockService) ByEndpointInterface(request *MockRequest, endpointInterface model.EndpointInterface, ctx iris.Context) (resp *MockResponse, err error) {
	generatorOptions := mockData.Options{
		//UseExamples:     factory.configuration.UseExamples,
		//NullProbability: factory.configuration.NullProbability,
		//DefaultMinInt:   factory.configuration.DefaultMinInt,
		//DefaultMaxInt:   factory.configuration.DefaultMaxInt,
		//DefaultMinFloat: factory.configuration.DefaultMinFloat,
		//DefaultMaxFloat: factory.configuration.DefaultMaxFloat,
		//SuppressErrors:  factory.configuration.SuppressErrors,
	}

	dataGeneratorInstance := mockData.New(generatorOptions)
	s.responseGenerator = mockGenerator.New(dataGeneratorInstance)
	s.responder = mockResponder.New()

	// generate openapi spec
	endpoint, err := s.EndpointRepo.GetAll(endpointInterface.EndpointId, "v0.1.0")
	spec := s.EndpointService.Yaml(endpoint)
	doc3 := spec.(*openapi3.T)

	var result interface{}
	commonUtils.JsonDecode(commonUtils.JsonEncode(doc3), &result)
	respContent, _ := encoder.NewEncoder(result).Encode()

	log.Println(string(respContent))

	// TODO: check
	doc3.Servers = nil
	doc3.Paths["/get"].Post = nil
	doc3.Info.Version = "1.0.0"
	desc := "描述"
	doc3.Paths["/get"].Get.Responses["200"].Value.Description = &desc

	// load openapi spec
	//specificationLoader := mockLoader.New()
	//specification, err := specificationLoader.LoadFromURI(factory.configuration.SpecificationURL)

	// init mock router
	s.router, err = legacy.NewRouter(doc3)
	if err != nil {
		//http.NotFound(writer, request)
		return
	}

	// generate mock response
	u := url.URL{
		Path: "/get",
	}
	httpRequest := http.Request{
		Method: http.MethodGet,
		URL:    &u,
	}
	route, pathParameters, err := s.router.FindRoute(&httpRequest)
	if err != nil {
		//http.NotFound(writer, request)
		return
	}

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
		//http.NotFound(writer, request)
		return
	}

	response, err := s.responseGenerator.GenerateResponse(&httpRequest, route)
	if err != nil {
		//handler.responder.WriteError(ctx, writer, err)
		return
	}

	log.Println(response)
	//handler.responder.WriteResponse(ctx, writer, response)

	return
}

func (s *MockService) GetEndpointInterface(req *MockRequest, endpointInterfaceId uint) (ret model.EndpointInterface, err error) {
	if endpointInterfaceId <= 0 {
		project, _ := s.ProjectRepo.GetByCode(req.ProjectCode)
		serve, _ := s.ServeRepo.GetByCode(project.ID, req.ServeCode)
		endpoint, _ := s.EndpointRepo.GetByPath(serve.ID, req.EndpointPath)

		_, endpointInterfaceId = s.EndpointInterfaceRepo.GetByMethod(endpoint.ID, req.EndpointMethod)
	}

	ret, err = s.EndpointInterfaceRepo.Get(endpointInterfaceId)

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
