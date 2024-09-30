package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/server/modules/model"
	"github.com/getkin/kin-openapi/openapi3"
	"path"
	"strings"
)

func ConvertPathsToInterfaces(doc openapi3.T) (interfaces []model.EndpointInterface, err error) {
	// TODO: convert to new openapi3 style models and save

	for pth, item := range doc.Paths {
		url := "${server}"
		url = path.Join(url, pth)

		interf, _ := convertOperations(url, item)
		interfaces = append(interfaces, interf...)
	}

	return
}

func convertOperations(url string, pth *openapi3.PathItem) (interfaces []model.EndpointInterface, err error) {
	if pth.Connect != nil {
		interf, err := convertOperation(url, pth.Connect)
		interf.Method = consts.CONNECT
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Delete != nil {
		interf, err := convertOperation(url, pth.Delete)
		interf.Method = consts.DELETE
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Get != nil {
		interf, err := convertOperation(url, pth.Get)
		interf.Method = consts.GET
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Head != nil {
		interf, err := convertOperation(url, pth.Head)
		interf.Method = consts.HEAD
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Options != nil {
		interf, err := convertOperation(url, pth.Options)
		interf.Method = consts.OPTIONS
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Patch != nil {
		interf, err := convertOperation(url, pth.Patch)
		interf.Method = consts.PATCH
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Post != nil {
		interf, err := convertOperation(url, pth.Post)
		interf.Method = consts.POST
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Put != nil {
		interf, err := convertOperation(url, pth.Put)
		interf.Method = consts.PUT
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	if pth.Trace != nil {
		interf, err := convertOperation(url, pth.Trace)
		interf.Method = consts.TRACE
		if err == nil {
			interfaces = append(interfaces, interf)
		}
	}

	return
}

func convertOperation(url string, operation *openapi3.Operation) (interf model.EndpointInterface, err error) {
	// common
	interf.Url = url
	interf.Name = operation.Summary
	if interf.Name == "" {
		interf.Name = "新接口"
	}

	// headers and params
	for _, item := range operation.Parameters {
		paramIn := item.Value.In

		if paramIn == "header" {
			header, err := genHeader(item)
			if err == nil {
				interf.Headers = append(interf.Headers, header)
			}
		} else if paramIn == "path" {
		} else if paramIn == "query" {
			param, err := genQueryParam(item)
			if err == nil {
				interf.Params = append(interf.Params, param)
			}
		} else if paramIn == "body" {
			interf.Body, err = genBody(item)
		}
	}

	return
}

func genBody(param *openapi3.ParameterRef) (ret string, err error) {
	val := param.Value.Schema.Value

	if val.Type == "object" {
		ret, err = genBodyFromSchema(val)
	}

	return
}
func genBodyFromSchema(schema *openapi3.Schema) (ret string, err error) {
	mp := map[string]interface{}{}

	for key, prop := range schema.Properties {
		var val interface{}
		example := prop.Value.Example

		fmt.Println(prop.Value.Type)

		if prop.Value.Type == "string" {
			if example != nil {
				val = example.(string)
			} else {
				val = ""
			}

		} else if prop.Value.Type == "integer" {
			if example != nil {
				val = example.(int)
			} else {
				val = 0
			}
		}

		mp[key] = val
	}

	bytes, _ := json.Marshal(mp)
	ret = string(bytes)

	return
}

func genHeader(param *openapi3.ParameterRef) (ret model.EndpointInterfaceHeader, err error) {
	ret.Name = param.Value.Name
	ret.Desc = param.Value.Description
	ret.Type, _ = genDataType(param.Value.ExtensionProps.Extensions)
	ret.Value = getExampleFromParam(param.Value)

	return
}

func genQueryParam(param *openapi3.ParameterRef) (ret model.EndpointInterfaceParam, err error) {
	ret.Name = param.Value.Name
	ret.Desc = param.Value.Description
	ret.Type, _ = genDataType(param.Value.ExtensionProps.Extensions)
	ret.Value = getExampleFromParam(param.Value)

	return
}

func getExampleFromParam(param *openapi3.Parameter) (ret string) {
	if param.Example != nil {
		ret = fmt.Sprintf("%v", param.Example)
		return
	} else if param.Examples != nil {
		for _, item := range param.Examples {
			ret = fmt.Sprintf("%v", item.Value.Value)
			ret = strings.TrimSpace(ret)
			if ret != "" {
				return
			}
		}
	} else if param.Schema != nil {
		if param.Schema.Value.Example != nil {
			ret = fmt.Sprintf("%v", param.Schema.Value.Example)
			return
		}
	}

	return
}

func genDataType(mp map[string]interface{}) (ret string, err error) {
	for key, _ := range mp {
		if key == "type" {
			val := mp[key].(json.RawMessage)

			json.Unmarshal(val, &ret)
			return
		}
	}

	return
}
