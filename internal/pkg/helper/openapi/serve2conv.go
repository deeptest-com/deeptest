package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	"github.com/getkin/kin-openapi/openapi3"
)

type serve2conv struct {
	serve     model.Serve
	endpoints []model.Endpoint
	doc3      *openapi3.T
}

func NewServe2conv(serve model.Serve, endpoints []model.Endpoint) *serve2conv {
	doc3 := &openapi3.T{
		OpenAPI: "3.0.3",
	}
	return &serve2conv{serve: serve, endpoints: endpoints, doc3: doc3}
}

func (s *serve2conv) ToV3() *openapi3.T {
	s.doc3.Info = s.info()
	s.doc3.Components = s.components()
	s.doc3.Servers = s.servers()
	s.doc3.Paths = s.paths()
	return s.doc3
}

func (s *serve2conv) info() (info *openapi3.Info) {
	info = new(openapi3.Info)
	info.Version = ""
	info.Title = s.serve.Name
	info.Description = s.serve.Description
	return
}

func (s *serve2conv) components() (components openapi3.Components) {
	components = openapi3.Components{}
	return
}

func (s *serve2conv) servers() (servers openapi3.Servers) {
	servers = openapi3.Servers{}
	return
}

func (s *serve2conv) paths() (paths openapi3.Paths) {
	paths = openapi3.Paths{}
	for _, endpoint := range s.endpoints {
		paths[endpoint.Path] = new(openapi3.PathItem)
		paths[endpoint.Path].Parameters = s.pathParameters(endpoint.PathParams) //
		for _, item := range endpoint.Interfaces {
			switch item.Method {
			case "GET":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				paths[endpoint.Path].Get.Description = item.Desc
				paths[endpoint.Path].Get.RequestBody = s.requestBody(item.RequestBody)
				paths[endpoint.Path].Get.Responses = nil
				paths[endpoint.Path].Get.Security = nil
			}

		}
	}
	return
}

func (s *serve2conv) pathParameters(params []model.EndpointPathParam) (parameters openapi3.Parameters) {
	parameters = openapi3.Parameters{}
	return
}

func (s *serve2conv) requestBody(body model.InterfaceRequestBody) (requestBody *openapi3.RequestBodyRef) {
	requestBody = new(openapi3.RequestBodyRef)
	return
}
