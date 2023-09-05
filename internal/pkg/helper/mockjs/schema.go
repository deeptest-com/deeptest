package mockjsHelper

import (
	"encoding/json"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
)

func IsMockJsSchema(schema *openapi3.Schema) bool {
	extensionProps := schema.Extensions

	for key, val := range extensionProps {
		if key == consts.KEY_MOCKJS {
			value, _ := val.(json.RawMessage).MarshalJSON()
			if len(value) > 0 {
				return true
			}
		}
	}

	return false
}

func GetMockJsSchemaExpression(schema *openapi3.Schema) string {
	extensionProps := schema.Extensions

	for key, val := range extensionProps {
		if key == consts.KEY_MOCKJS {
			value, _ := val.(json.RawMessage).MarshalJSON()
			if len(value) > 0 {
				return string(value)
			}
		}
	}

	return ""
}
