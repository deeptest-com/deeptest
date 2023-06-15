package openapi

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
)

type openapi2endpoint struct {
	doc       *openapi3.T
	endpoints []*model.Endpoint
}

func NewOpenapi2endpoint(doc *openapi3.T) *openapi2endpoint {
	return &openapi2endpoint{doc: doc}
}

func (o *openapi2endpoint) Convert() (endpoints []*model.Endpoint) {
	o.convertEndpoints()
	return o.endpoints
}

func (o *openapi2endpoint) convertEndpoints() {

	for url, path := range o.doc.Paths {
		endpoint := new(model.Endpoint)
		endpoint.Path = url
		endpoint.Interfaces = o.interfaces(url, path)
		endpoint.PathParams = o.pathParams(path.Parameters)
		if len(endpoint.Interfaces) > 0 {
			endpoint.Title = endpoint.Interfaces[0].Name
		}
		o.endpoints = append(o.endpoints, endpoint)
	}

	return
}

func (o *openapi2endpoint) pathParams(parameters openapi3.Parameters) (pathParams []model.EndpointPathParam) {
	for _, parameter := range parameters {
		var pathParam model.EndpointPathParam
		copier.CopyWithOption(&pathParam, o.parameter(parameter), copier.Option{IgnoreEmpty: true, DeepCopy: true})
		pathParams = append(pathParams, pathParam)
	}
	return
}

func (o *openapi2endpoint) interfaces(url string, path *openapi3.PathItem) (interfaces []model.EndpointInterface) {
	if path.Get != nil {
		interf := o.interf("GET", url, path.Get)
		interfaces = append(interfaces, interf)
	}

	if path.Post != nil {
		interf := o.interf("POST", url, path.Post)
		interfaces = append(interfaces, interf)
	}

	if path.Put != nil {
		interf := o.interf("PUT", url, path.Put)
		interfaces = append(interfaces, interf)
	}

	if path.Delete != nil {
		interf := o.interf("DELETE", url, path.Delete)
		interfaces = append(interfaces, interf)
	}

	if path.Trace != nil {
		interf := o.interf("TRACE", url, path.Trace)
		interfaces = append(interfaces, interf)
	}

	if path.Head != nil {
		interf := o.interf("HEAD", url, path.Head)
		interfaces = append(interfaces, interf)
	}

	if path.Options != nil {
		interf := o.interf("OPTIONS", url, path.Options)
		interfaces = append(interfaces, interf)
	}

	if path.Patch != nil {
		interf := o.interf("OPTIONS", url, path.Patch)
		interfaces = append(interfaces, interf)
	}
	return
}

func (o *openapi2endpoint) interf(method consts.HttpMethod, url string, operation *openapi3.Operation) (interf model.EndpointInterface) {
	interf = model.EndpointInterface{}
	interf.Name = operation.Summary
	interf.Method = method
	interf.Url = url
	interf.Headers, interf.Cookies, interf.Params = o.parameters(operation)
	if operation.RequestBody != nil {
		interf.BodyType, interf.RequestBody = o.requestBody(operation.RequestBody.Value.Content)
	}
	interf.ResponseBodies = o.responseBodies(operation.Responses)
	return
}

func (o *openapi2endpoint) BodyType() {

}

func (o *openapi2endpoint) requestBody(content openapi3.Content) (mediaType consts.HttpContentType, body model.EndpointInterfaceRequestBody) {
	for key, item := range content {
		mediaType = consts.HttpContentType(key)
		body = model.EndpointInterfaceRequestBody{}
		body.MediaType = key

		if item.Examples == nil {
			item.Examples = openapi3.Examples{}
			item.Examples["example"] = new(openapi3.ExampleRef)
			item.Examples["example"].Value = new(openapi3.Example)
			item.Examples["example"].Value.Value = item.Example
		}
		body.Examples = commonUtils.JsonEncode(item.Examples)
		body.SchemaItem = o.requestBodyItem(item)
		//body.Examples = item.Example
		//content.
		return
	}

	return
}

func (o *openapi2endpoint) requestBodyItem(item *openapi3.MediaType) (requestBodyItem model.EndpointInterfaceRequestBodyItem) {
	requestBodyItem = model.EndpointInterfaceRequestBodyItem{}
	requestBodyItem.Content = commonUtils.JsonEncode(item.Schema)
	requestBodyItem.Type = item.Schema.Value.Type

	return
}

func (o *openapi2endpoint) responseBodies(responses openapi3.Responses) (bodies []model.EndpointInterfaceResponseBody) {
	bodies = []model.EndpointInterfaceResponseBody{}
	for key, item := range responses {
		body := o.responseBody(item.Value)
		body.Code = key
		bodies = append(bodies, body)
		return
	}
	return
}

func (o *openapi2endpoint) responseBody(response *openapi3.Response) (body model.EndpointInterfaceResponseBody) {
	body = model.EndpointInterfaceResponseBody{}
	body.Headers = o.responseHeader(response.Headers)
	for key, item := range response.Content {
		body.MediaType = key
		body.Examples = commonUtils.JsonEncode(item.Examples)
		body.Description = *response.Description
		return
	}

	return
}

func (o *openapi2endpoint) responseHeader(h openapi3.Headers) (headers []model.EndpointInterfaceResponseBodyHeader) {
	for key, item := range h {
		header := model.EndpointInterfaceResponseBodyHeader{}
		header.Name = key
		header.Value = item.Value.Schema.Value.Default.(string)
		header.Type = item.Value.Schema.Value.Type
		headers = append(headers, header)
	}

	return
}

func (o *openapi2endpoint) parameters(operation *openapi3.Operation) (headers []model.EndpointInterfaceHeader, cookies []model.EndpointInterfaceCookie, params []model.EndpointInterfaceParam) {
	for _, parameter := range operation.Parameters {
		if parameter.Value.In == "header" {
			header := o.parameter(parameter)
			headers = append(headers, model.EndpointInterfaceHeader(header))
		} else if parameter.Value.In == "cookie" {
			cookie := o.parameter(parameter)
			cookies = append(cookies, model.EndpointInterfaceCookie(cookie))
		} else if parameter.Value.In == "query" {
			param := o.parameter(parameter)
			params = append(params, param)
		}
	}

	return
}

func (o *openapi2endpoint) parameter(parameter *openapi3.ParameterRef) (param model.EndpointInterfaceParam) {
	param = model.EndpointInterfaceParam{}
	param.Name = parameter.Value.Name
	param.Ref = parameter.Ref
	param.Required = parameter.Value.Required
	o.parameterValue(parameter.Value.Schema.Value, &param)
	return
}

func (*openapi2endpoint) parameterValue(schema *openapi3.Schema, param *model.EndpointInterfaceParam) {
	if schema.Example != nil {
		param.Example = schema.Example.(string)
	}

	if schema.MaxLength != nil {
		param.MaxLength = *schema.MaxLength
	}
	if schema.Default != nil {
		param.Default = schema.Default.(string)
	}

	if schema.MultipleOf != nil {
		param.MultipleOf = *schema.MultipleOf
	}

	if schema.MaxItems != nil {
		param.MaxItems = *schema.MaxItems
	}

	param.Pattern = schema.Pattern
	param.MinItems = schema.MinItems
	param.MinLength = schema.MinLength
	param.MinItems = schema.MinItems
	param.UniqueItems = schema.UniqueItems
	param.Type = schema.Type
}
