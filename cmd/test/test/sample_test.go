package test

import (
	commonUtils "github.com/aaronchen2k/deeptest/pkg/lib/comm"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"log"
	"testing"
)

func TestSample(t *testing.T) {
	json := `{
  "openapi": "3.0.2",
  "info": {
    "title": "My other API",
    "version": "0.1.0"
  },
  "components": {
    "schemas": {
      "DefaultObject": {
        "type": "object",
        "properties": {
          "foo": {
            "type": "string"
          },
          "bar": {
            "type": "integer"
          }
        }
      }
    },
    "responses": {
      "DefaultResponse": {
        "description": "",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/DefaultObject"
            }
          }
        }
      }
    }
  }
}
`
	doc3 := openapi3.T{}
	commonUtils.JsonDecode(json, &doc3)

	ret, err := gorillamux.NewRouter(&doc3)
	if err != nil {
		return
	}

	log.Println(ret)
}
