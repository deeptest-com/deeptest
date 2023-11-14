package mockjs

import (
	"encoding/json"
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
	"strconv"
	"strings"
)

func IsMockJsSchema(schema *openapi3.Schema) bool {
	extensionProps := schema.Extensions

	for key, val := range extensionProps {
		if key == consts.KEY_MOCKJS {
			value, _ := json.Marshal(val.(json.RawMessage))
			valueTrim := strings.Trim(string(value), "\"")

			if len(valueTrim) > 0 {
				return true
			}
		}
	}

	return false
}

func GetMockJsSchemaExpression(schema *openapi3.Schema) (ret string) {
	extensionProps := schema.Extensions

	for key, val := range extensionProps {
		if key == consts.KEY_MOCKJS {
			value, _ := json.Marshal(val.(json.RawMessage))
			ret = strings.Trim(string(value), "\"")
			if len(ret) > 0 {
				if ret[:1] != "@" {
					ret = "@" + ret
				}

				return
			}
		}
	}

	return ""
}

func ConvertData(data interface{}, schemaType string) (ret interface{}) {
	str := fmt.Sprintf("%v", data)

	switch schemaType {

	case openapi3.TypeBoolean:
		ret, _ = strconv.ParseBool(str)

	case openapi3.TypeInteger:
		ret, _ = strconv.ParseInt(str, 10, 64)

	case openapi3.TypeNumber:
		ret, _ = strconv.ParseFloat(str, 64)

	default:
		ret = fmt.Sprintf("%v", data)
	}

	return
}
