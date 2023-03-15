package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	_commUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
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
	components.Schemas = openapi3.Schemas{}
	for _, component := range s.serve.Components {
		schema := new(openapi3.Schema)
		if component.Type == openapi3.TypeObject {
			var schemas openapi3.Schemas
			_commUtils.JsonDecode(component.Content, &schemas)
			schema.Properties = schemas
		} else {
			var items *openapi3.SchemaRef
			_commUtils.JsonDecode(component.Content, &items)
			schema.Items = items
		}
		components.Schemas[component.Name] = openapi3.NewSchemaRef("", schema)
	}

	return
}

func (s *serve2conv) servers() (servers openapi3.Servers) {
	servers = openapi3.Servers{}
	for _, server := range s.serve.Servers {
		servers = append(servers, &openapi3.Server{URL: server.Url, Description: server.Description})
	}
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
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)

			case "POST":
				paths[endpoint.Path].Post = new(openapi3.Operation)
				//paths[endpoint.Path].Get.OperationID = item.OperationId
				//paths[endpoint.Path].Get.Description = item.Description
				//paths[endpoint.Path].Get.Summary = item.Description
				paths[endpoint.Path].Post.RequestBody = s.requestBody(item.RequestBody)
				paths[endpoint.Path].Post.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Post.Security = nil
				paths[endpoint.Path].Post.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			case "PUT":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			case "PATCH":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			case "DELETE":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			case "HEAD":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			case "OPTIONS":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			case "TRACE":
				paths[endpoint.Path].Get = new(openapi3.Operation)
				/*				paths[endpoint.Path].Get.OperationID = item.OperationId
								paths[endpoint.Path].Get.Description = item.Description
								paths[endpoint.Path].Get.Summary = item.Description*/
				paths[endpoint.Path].Get.Responses = s.responsesBody(item.ResponseBodies)
				paths[endpoint.Path].Get.Security = nil
				paths[endpoint.Path].Get.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
			}

		}
	}
	return
}

func (s *serve2conv) pathParameters(params []model.EndpointPathParam) (parameters openapi3.Parameters) {
	parameters = openapi3.Parameters{}
	for _, param := range params {
		parameterRef := new(openapi3.ParameterRef)
		parameterRef.Value = new(openapi3.Parameter)
		parameterRef.Value.In = "path"
		parameterRef.Value.Name = param.Name
		parameterRef.Value.Required = true
		parameterRef.Value.Schema = new(openapi3.SchemaRef)
		parameterRef.Value.Schema.Value = new(openapi3.Schema)
		parameterRef.Value.Schema.Value.Type = param.Type
		parameters = append(parameters, parameterRef)
	}
	return
}

func (s *serve2conv) parameters(cookies []model.InterfaceCookie, headers []model.InterfaceHeader, params []model.InterfaceParam) (parameters openapi3.Parameters) {
	parameters = openapi3.Parameters{}
	for _, param := range params {
		parameterRef := new(openapi3.ParameterRef)
		parameterRef.Value = new(openapi3.Parameter)
		parameterRef.Value.In = "query"
		parameterRef.Value.Name = param.Name
		parameterRef.Value.Required = true
		parameterRef.Value.Schema = new(openapi3.SchemaRef)
		parameterRef.Value.Schema.Value = new(openapi3.Schema)
		parameterRef.Value.Schema.Value.Type = param.Type
		parameters = append(parameters, parameterRef)
	}
	for _, header := range headers {
		parameterRef := new(openapi3.ParameterRef)
		parameterRef.Value = new(openapi3.Parameter)
		parameterRef.Value.In = "header"
		parameterRef.Value.Name = header.Name
		parameterRef.Value.Required = true
		parameterRef.Value.Schema = new(openapi3.SchemaRef)
		parameterRef.Value.Schema.Value = new(openapi3.Schema)
		parameterRef.Value.Schema.Value.Type = header.Type
		parameters = append(parameters, parameterRef)
	}
	for _, cookie := range cookies {
		parameterRef := new(openapi3.ParameterRef)
		parameterRef.Value = new(openapi3.Parameter)
		parameterRef.Value.In = "cookie"
		parameterRef.Value.Name = cookie.Name
		parameterRef.Value.Required = true
		parameterRef.Value.Schema = new(openapi3.SchemaRef)
		parameterRef.Value.Schema.Value = new(openapi3.Schema)
		parameterRef.Value.Schema.Value.Type = cookie.Type
		parameters = append(parameters, parameterRef)
	}
	return
}

func (s *serve2conv) requestBody(body model.InterfaceRequestBody) (requestBody *openapi3.RequestBodyRef) {
	requestBody = new(openapi3.RequestBodyRef)
	requestBody.Value = new(openapi3.RequestBody)
	requestBody.Value.Description = ""
	requestBody.Value.Content = openapi3.Content{}
	requestBody.Value.Content[body.MediaType] = new(openapi3.MediaType)
	requestBody.Value.Content[body.MediaType].Schema = new(openapi3.SchemaRef)
	//if body.SchemaItem.RequestBodyId != 0 {
	requestBody.Value.Content[body.MediaType].Schema.Value = s.requestBodySchema(body.SchemaItem)
	requestBody.Value.Content[body.MediaType].Examples = s.requestBodyExamples(body.Examples)
	//}
	return
}

func (s *serve2conv) requestBodySchema(item model.InterfaceRequestBodyItem) (schema *openapi3.Schema) {
	schema = new(openapi3.Schema)
	schema.Type = item.Type
	if item.Type == openapi3.TypeObject {
		var schemas openapi3.Schemas
		_commUtils.JsonDecode(item.Content, &schemas)
		schema.Properties = schemas
	} else {
		var items *openapi3.SchemaRef
		_commUtils.JsonDecode(item.Content, &items)
		schema.Items = items
	}
	return
}

func (s *serve2conv) requestBodyExamples(examplesStr string) (examples openapi3.Examples) {
	//examplesStr = "{\"user\":{\"value\":{\"id\":1,\"name\":\"王大锤\"}},\"product\":{\"value\":{\"id\":1,\"name\":\"服装\"}}}"
	_commUtils.JsonDecode(examplesStr, &examples)
	return
}

func (s *serve2conv) responsesBody(bodies []model.InterfaceResponseBody) (responsesBody openapi3.Responses) {
	responsesBody = openapi3.Responses{}
	for _, body := range bodies {
		responsesBody[body.Code] = new(openapi3.ResponseRef)
		responsesBody[body.Code].Value = new(openapi3.Response)
		responsesBody[body.Code].Value.Description = nil
		responsesBody[body.Code].Value.Content = openapi3.Content{}
		responsesBody[body.Code].Value.Content[body.MediaType] = new(openapi3.MediaType)
		responsesBody[body.Code].Value.Content[body.MediaType].Schema = new(openapi3.SchemaRef)
		responsesBody[body.Code].Value.Content[body.MediaType].Schema.Value = s.responsesBodySchema(body.SchemaItem)
	}
	return
}

func (s *serve2conv) responsesBodySchema(item model.InterfaceResponseBodyItem) (schema *openapi3.Schema) {
	schema = new(openapi3.Schema)
	schema.Type = item.Type
	//fmt.Println(item, "++++++++++++++++")
	if item.Type == "object" {
		//schema.Properties = openapi3.Schemas{}
		var schemas openapi3.Schemas
		item.Content = strings.ReplaceAll(item.Content, "\n", "")
		_commUtils.JsonDecode(item.Content, &schemas)
		//fmt.Println(item.Content, &schemas)
		//for _,val := range content{
		//v := val.(map[string]interface{})
		//schema.Properties[v["name"].(string)] =
		//}
		schema.Properties = schemas
	} else {

	}

	return
}
