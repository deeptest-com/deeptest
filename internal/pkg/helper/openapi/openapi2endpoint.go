package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/server/modules/model"
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinzhu/copier"
	"reflect"
	"strconv"
	"strings"
)

type openapi2endpoint struct {
	doc              *openapi3.T
	endpoints        []*model.Endpoint
	dirs             *Dirs
	components       map[string]*model.ComponentSchema
	componentSchemas map[string]Schema
}

type Dirs struct {
	Id   int64
	Dirs map[string]*Dirs
}

func NewOpenapi2endpoint(doc *openapi3.T, dirId int64) *openapi2endpoint {
	return &openapi2endpoint{doc: doc, dirs: &Dirs{Id: dirId}}
}

func (o *openapi2endpoint) Convert() (endpoints []*model.Endpoint, dirs *Dirs, components map[string]*model.ComponentSchema) {
	o.convertComponents()
	o.convertEndpoints()
	return o.endpoints, o.dirs, o.components
}

func (o *openapi2endpoint) convertComponents() {
	o.components = make(map[string]*model.ComponentSchema)
	for key, schema := range o.doc.Components.Schemas {
		content, err := json.Marshal(schema.Value)
		if err != nil {
			panic(err)
		}

		ref := "#/components/schemas/" + key
		if schema.Value.Type == "" {
			if schema.Value.Properties != nil {
				schema.Value.Type = openapi3.TypeObject
			} else if schema.Value.Items != nil {
				schema.Value.Type = openapi3.TypeArray
			}
		}
		component := model.ComponentSchema{Name: key, Type: schema.Value.Type, Content: string(content), Ref: ref}
		o.components[ref] = &component
		//o.componentSchemas[ref] = Schema{Type: schema.Value.Type,Items: }
	}
}

func (o *openapi2endpoint) convertEndpoints() {

	for url, path := range o.doc.Paths {
		o.addMethod(url, path.ExtensionProps)
		interfaces := o.interfaces(url, o.doc.Paths[url])
		pathParams := o.pathParams(o.doc.Paths[url].Parameters)
		for _, interf := range interfaces {
			endpoint := new(model.Endpoint)
			endpoint.Path = url
			if len(pathParams) == 0 {
				pathParams = interf.PathParams
			}
			endpoint.PathParams = pathParams
			endpoint.Title = interf.Name
			endpoint.Interfaces = append(endpoint.Interfaces, interf)
			endpoint.Tags = interf.Tags
			endpoint.CreateUser = interf.Creator
			o.endpoints = append(o.endpoints, endpoint)
		}

	}

	return
}

func (o *openapi2endpoint) addMethod(url string, extensions openapi3.ExtensionProps) {
	for method, extension := range extensions.Extensions {
		var operation *openapi3.Operation
		json.Unmarshal(extension.(json.RawMessage), &operation)
		o.doc.AddOperation(url, method, operation)
	}

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
	var interf model.EndpointInterface
	if path.Get != nil {
		interf = o.interf("GET", url, path.Get)
		interfaces = append(interfaces, interf)
	}

	if path.Post != nil {
		interf = o.interf("POST", url, path.Post)
		interfaces = append(interfaces, interf)
	}

	if path.Put != nil {
		interf = o.interf("PUT", url, path.Put)
		interfaces = append(interfaces, interf)
	}

	if path.Delete != nil {
		interf = o.interf("DELETE", url, path.Delete)
		interfaces = append(interfaces, interf)
	}

	if path.Trace != nil {
		interf = o.interf("TRACE", url, path.Trace)
		interfaces = append(interfaces, interf)
	}

	if path.Head != nil {
		interf = o.interf("HEAD", url, path.Head)
		interfaces = append(interfaces, interf)
	}

	if path.Options != nil {
		interf = o.interf("OPTIONS", url, path.Options)
		interfaces = append(interfaces, interf)
	}

	if path.Patch != nil {
		interf = o.interf("OPTIONS", url, path.Patch)
		interfaces = append(interfaces, interf)
	}
	return
}

func (o *openapi2endpoint) interf(method consts.HttpMethod, url string, operation *openapi3.Operation) (interf model.EndpointInterface) {
	interf = model.EndpointInterface{}
	interf.Name = operation.Summary
	interf.Method = method
	interf.Url = url
	if interf.Name == "" {
		interf.Name = interf.Url
	}
	var requestBodyItem []model.EndpointInterfaceRequestBodyItem
	interf.Headers, interf.Cookies, interf.Params, interf.PathParams, requestBodyItem = o.parameters(operation)
	if operation.RequestBody != nil {
		interf.BodyType, interf.RequestBody = o.requestBody(operation.RequestBody.Value.Content)
		if len(requestBodyItem) > 0 {
			interf.RequestBody.SchemaItem = requestBodyItem[0]
		}
	}
	interf.ResponseBodies = o.responseBodies(operation.Responses)
	interf.Tags = o.makeDirs(operation.Tags)
	interf.Creator = o.creator(operation.Extensions)
	return
}

func (o *openapi2endpoint) creator(extensions map[string]interface{}) (res string) {
	if value, ok := extensions["x-creator"]; ok {
		json.Unmarshal(value.(json.RawMessage), &res)
	}
	return
}

func (o *openapi2endpoint) makeDirs(tags []string) []string {
	d := o.dirs
	if len(tags) > 0 {
		tags = strings.Split(tags[0], "/")
	}
	for _, tag := range tags {
		d = o.makeDir(tag, d)
	}

	return tags
}

func (o *openapi2endpoint) makeDir(tag string, d *Dirs) *Dirs {
	tag = strings.TrimSpace(tag)
	if d.Dirs == nil {
		d.Dirs = map[string]*Dirs{}
	}

	if _, ok := d.Dirs[tag]; !ok {
		d.Dirs[tag] = new(Dirs)
	}

	return d.Dirs[tag]
}

func (o *openapi2endpoint) BodyType() {

}

func (o *openapi2endpoint) requestBody(content openapi3.Content) (mediaType consts.HttpContentType, body model.EndpointInterfaceRequestBody) {
	for key, item := range content {
		mediaType = consts.HttpContentType(key)
		body = model.EndpointInterfaceRequestBody{}
		body.MediaType = key
		body.Examples = o.requestBodyExamples(item)
		body.SchemaItem = o.requestBodyItem(item.Schema)
		return
	}

	return
}

func (o *openapi2endpoint) requestBodyExamples(item *openapi3.MediaType) string {
	if item.Examples == nil {
		item.Examples = openapi3.Examples{}
		item.Examples["example"] = new(openapi3.ExampleRef)
		item.Examples["example"].Value = new(openapi3.Example)
		if item.Example == nil {
			return ""
		}
		item.Examples["example"].Value.Value = item.Example
	}
	var examples []map[string]string
	for name, example := range item.Examples {
		value := map[string]string{"name": name, "content": commonUtils.JsonEncode(example.Value.Value)}
		examples = append(examples, value)
	}

	return commonUtils.JsonEncode(examples)
}

func (o *openapi2endpoint) requestBodyItem(schema *openapi3.SchemaRef) (requestBodyItem model.EndpointInterfaceRequestBodyItem) {
	requestBodyItem = model.EndpointInterfaceRequestBodyItem{}
	requestBodyItem.Content = commonUtils.JsonEncode(schema)
	if schema.Value != nil {
		requestBodyItem.Type = schema.Value.Type
	}

	return
}

func (o *openapi2endpoint) responseBodies(responses openapi3.Responses) (bodies []model.EndpointInterfaceResponseBody) {
	bodies = []model.EndpointInterfaceResponseBody{}
	for key, item := range responses {
		body := o.responseBody(item.Value)
		if _, err := strconv.Atoi(key); err != nil {
			key = "200"
		}
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
		body.Examples = commonUtils.JsonEncode(o.responseBodyExamples(item.Examples))
		body.Description = *response.Description
		body.SchemaItem = o.responseBodyItem(item.Schema)
		return
	}

	return
}

func (o *openapi2endpoint) responseBodyExamples(examples openapi3.Examples) (ret []interface{}) {
	for key, example := range examples {
		ret = append(ret, map[string]interface{}{"name": "example_" + key, "content": example})
	}
	return
}

func (o *openapi2endpoint) responseBodyItem(schema *openapi3.SchemaRef) (item model.EndpointInterfaceResponseBodyItem) {
	item = model.EndpointInterfaceResponseBodyItem{}
	item.Content = commonUtils.JsonEncode(schema)
	if schema.Value != nil {
		item.Type = schema.Value.Type
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

func (o *openapi2endpoint) parameters(operation *openapi3.Operation) (headers []model.EndpointInterfaceHeader, cookies []model.EndpointInterfaceCookie, params []model.EndpointInterfaceParam, pathParams []model.EndpointPathParam, bodyItem []model.EndpointInterfaceRequestBodyItem) {
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
		} else if parameter.Value.In == "path" {
			param := o.parameter(parameter)
			pathParams = append(pathParams, model.EndpointPathParam{EndpointInterfaceParam: param})
		} else if parameter.Value.In == "body" {
			if parameter.Value.Schema != nil {
				item := o.requestBodyItem(parameter.Value.Schema)
				bodyItem = append(bodyItem, item)
			}

		}

	}

	return
}

func (o *openapi2endpoint) parameter(parameter *openapi3.ParameterRef) (param model.EndpointInterfaceParam) {
	param = model.EndpointInterfaceParam{}
	param.Name = parameter.Value.Name
	param.Ref = parameter.Ref
	param.Required = parameter.Value.Required
	if parameter.Value.Schema != nil {
		o.parameterValue(parameter.Value.Schema.Value, &param)
	}

	return
}

func (o *openapi2endpoint) parameterValue(schema *openapi3.Schema, param *model.EndpointInterfaceParam) {
	if schema.Example != nil {
		param.Example = fmt.Sprintf("%v", schema.Example)
	}

	if schema.MaxLength != nil {
		param.MaxLength = *schema.MaxLength
	}
	if schema.Default != nil {
		param.Default = fmt.Sprintf("%v", schema.Default)

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

func (o *openapi2endpoint) schemaRefAddType(schema *openapi3.SchemaRef) {

	if schema.Ref != "" {
		refSchema := o.components[schema.Ref]
		fieldValue := reflect.ValueOf(refSchema.Type)
		fieldName := reflect.ValueOf("Type")
		schema.Value = openapi3.NewSchema()
		indirectValue := reflect.Indirect(reflect.ValueOf(schema.Value))
		indirectValue.FieldByName(fieldName.String()).Set(fieldValue)
	}

}
