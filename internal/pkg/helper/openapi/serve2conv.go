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
	info.Version = "1.0.0"
	info.Title = s.serve.Name
	info.Description = s.serve.Description
	return
}

func (s *serve2conv) components() (components openapi3.Components) {
	components = openapi3.NewComponents()

	components.Schemas = openapi3.Schemas{}
	for _, component := range s.serve.Components {
		schema := new(openapi3.Schema)
		if component.Type == openapi3.TypeObject {
			//var schemas openapi3.Schema
			_commUtils.JsonDecode(component.Content, &schema)
			//schema = schemas
		} else {
			var items *openapi3.SchemaRef
			_commUtils.JsonDecode(component.Content, &items)
			schema.Items = items
		}
		components.Schemas[component.Name] = openapi3.NewSchemaRef("", schema)
	}

	components.SecuritySchemes = s.security()
	return
}

func (s *serve2conv) security() (securitySchemes openapi3.SecuritySchemes) {
	securitySchemes = openapi3.SecuritySchemes{}
	for _, security := range s.serve.Securities {
		securityScheme := openapi3.NewSecurityScheme()
		securityScheme.Type = security.Type
		switch security.Type {
		case "apiKey":
			securityScheme.In = security.In
			securityScheme.Name = security.Key
		case "bearerToken":
			securityScheme.Type = "http"
			securityScheme.Scheme = "bearer"
			securityScheme.BearerFormat = "jwt"
		case "basicAuth":
			securityScheme.Type = "http"
			securityScheme.Scheme = "basic"
			securityScheme.BearerFormat = ""
		}
		if security.Default {
			securityRequirement := openapi3.NewSecurityRequirement()
			securityRequirement[securityScheme.Name] = nil
			s.doc3.Security = openapi3.SecurityRequirements{securityRequirement}
		}
		securitySchemes[security.Name] = &openapi3.SecuritySchemeRef{Value: securityScheme}
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
		paths[endpoint.Path].Description = endpoint.Description
		for _, item := range endpoint.Interfaces {
			switch item.Method {
			case "GET":
				paths[endpoint.Path].Get = s.operation(item)
			case "POST":
				paths[endpoint.Path].Post = s.operation(item)
			case "PUT":
				paths[endpoint.Path].Put = s.operation(item)
			case "PATCH":
				paths[endpoint.Path].Patch = s.operation(item)
			case "DELETE":
				paths[endpoint.Path].Delete = s.operation(item)
			case "HEAD":
				paths[endpoint.Path].Head = s.operation(item)
			case "OPTIONS":
				paths[endpoint.Path].Options = s.operation(item)
			case "TRACE":
				paths[endpoint.Path].Trace = s.operation(item)
			}

		}
	}
	return
}

func (s *serve2conv) operation(item model.EndpointInterface) (operation *openapi3.Operation) {
	operation = new(openapi3.Operation)
	operation.OperationID = item.OperationId
	operation.Description = item.Description
	operation.Summary = item.Description
	operation.RequestBody = s.requestBody(item.RequestBody)
	operation.Responses = s.responsesBody(item.ResponseBodies)
	operation.Parameters = s.parameters(item.Cookies, item.Headers, item.Params)
	if item.Security != "" {
		securityRequirement := openapi3.NewSecurityRequirement()
		securityRequirement[item.Security] = nil
		operation.Security = &openapi3.SecurityRequirements{securityRequirement}
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

func (s *serve2conv) parameters(cookies []model.EndpointInterfaceCookie, headers []model.EndpointInterfaceHeader, params []model.EndpointInterfaceParam) (parameters openapi3.Parameters) {
	parameters = openapi3.Parameters{}
	for _, param := range params {
		parameterRef := s.parameterRef("query", param)
		parameters = append(parameters, parameterRef)
	}
	for _, header := range headers {
		parameterRef := s.parameterRef("header", model.EndpointInterfaceParam(header))
		parameters = append(parameters, parameterRef)
	}
	for _, cookie := range cookies {
		parameterRef := s.parameterRef("cookie", model.EndpointInterfaceParam(cookie))
		parameters = append(parameters, parameterRef)
	}
	return
}

func (s *serve2conv) parameterRef(in string, param model.EndpointInterfaceParam) (parameterRef *openapi3.ParameterRef) {
	parameterRef = new(openapi3.ParameterRef)
	parameterRef.Value = new(openapi3.Parameter)
	parameterRef.Value.In = in
	parameterRef.Value.Name = param.Name
	parameterRef.Value.Schema = new(openapi3.SchemaRef)
	parameterRef.Value.Schema.Ref = param.Ref
	parameterRef.Value.Required = param.Required
	parameterRef.Value.Schema.Value = s.schemaValue(param)
	return
}

func (s *serve2conv) schemaValue(param model.EndpointInterfaceParam) (schema *openapi3.Schema) {
	schema = new(openapi3.Schema)
	schema.Example = param.Example
	schema.Pattern = param.Pattern
	schema.MinLength = param.MinLength
	schema.MaxLength = &param.MaxLength
	schema.Default = param.Default
	schema.MultipleOf = &param.MultipleOf
	schema.MinItems = param.MinItems
	schema.MaxItems = &param.MaxItems
	schema.UniqueItems = param.UniqueItems
	schema.Type = param.Type
	return
}

func (s *serve2conv) requestBody(body model.EndpointInterfaceRequestBody) (requestBody *openapi3.RequestBodyRef) {
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

func (s *serve2conv) requestBodySchema(item model.EndpointInterfaceRequestBodyItem) (schema *openapi3.Schema) {
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
	examples = make(openapi3.Examples)
	//examplesStr = "{\"user\":{\"value\":{\"id\":1,\"name\":\"王大锤\"}},\"product\":{\"value\":{\"id\":1,\"name\":\"服装\"}}}"
	var examplesArr []map[string]string
	_commUtils.JsonDecode(examplesStr, &examplesArr)
	for _, item := range examplesArr {
		example := new(openapi3.ExampleRef)
		content := item["content"]
		content = strings.ReplaceAll(content, "\r\n", "")
		content = strings.ReplaceAll(content, "\n", "")
		var res interface{}
		_commUtils.JsonDecode(content, &res)
		if res != nil {
			example.Value = new(openapi3.Example)
			example.Value.Value = res
			examples[item["name"]] = example
		}

	}

	return
}

func (s *serve2conv) responsesBody(bodies []model.EndpointInterfaceResponseBody) (responsesBody openapi3.Responses) {
	responsesBody = openapi3.Responses{}
	for _, body := range bodies {
		responsesBody[body.Code] = new(openapi3.ResponseRef)
		responsesBody[body.Code].Value = new(openapi3.Response)
		responsesBody[body.Code].Value.Description = nil
		responsesBody[body.Code].Value.Content = openapi3.Content{}
		responsesBody[body.Code].Value.Content[body.MediaType] = new(openapi3.MediaType)
		responsesBody[body.Code].Value.Content[body.MediaType].Schema = new(openapi3.SchemaRef)
		responsesBody[body.Code].Value.Content[body.MediaType].Schema = s.responsesBodySchema(body)
		responsesBody[body.Code].Value.Content[body.MediaType].Examples = s.responsesBodyExamples(body)
		responsesBody[body.Code].Value.Headers = s.responseBodyHeaders(body.Headers)
		responsesBody[body.Code].Value.Description = &body.Description
	}
	return
}

func (s *serve2conv) responseBodyHeaders(headers []model.EndpointInterfaceResponseBodyHeader) (res openapi3.Headers) {
	res = openapi3.Headers{}
	for _, item := range headers {
		res[item.Name] = new(openapi3.HeaderRef)
		res[item.Name].Value = new(openapi3.Header)
		res[item.Name].Value.Schema = new(openapi3.SchemaRef)
		res[item.Name].Value.Schema.Value = new(openapi3.Schema)
		res[item.Name].Value.Schema.Value.Type = item.Type
		if item.Default == "" {
			item.Default = item.Example
		}
		res[item.Name].Value.Schema.Value.Default = item.Default
		//TODO DETAIL
	}
	return
}

func (s *serve2conv) responsesBodySchema(responseBody model.EndpointInterfaceResponseBody) (schema *openapi3.SchemaRef) {
	//schema = new(openapi3.Schema)

	/*
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
	*/

	schema = new(openapi3.SchemaRef)
	responseBody.SchemaItem.Content = strings.ReplaceAll(responseBody.SchemaItem.Content, "\n", "")
	_commUtils.JsonDecode(responseBody.SchemaItem.Content, schema)

	return
}

func (s *serve2conv) responsesBodyExamples(responseBody model.EndpointInterfaceResponseBody) (examples openapi3.Examples) {
	//schema = new(openapi3.Schema)

	/*
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
	*/

	return s.requestBodyExamples(responseBody.Examples)

}
