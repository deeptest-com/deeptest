package mockData

import (
	"context"
	"fmt"
	mockjsHelper "github.com/aaronchen2k/deeptest/internal/pkg/helper/mockjs"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
)

type coordinatingSchemaGenerator struct {
	generatorsByType map[string]schemaGenerator
}

func (generator *coordinatingSchemaGenerator) GenerateDataBySchema(ctx context.Context, schema *openapi3.Schema) (Data, error) {
	schemaType := generator.detectSchemaType(schema)

	specificGenerator, exists := generator.generatorsByType[schemaType]
	logUtils.Infof("*** use generator %#v", specificGenerator)
	if !exists {
		return nil, errors.WithStack(&ErrGenerationFailed{
			GeneratorID: "coordinatingSchemaGenerator",
			Message:     fmt.Sprintf("data generation for objects of type '%s' is not supported", schemaType),
		})
	}

	return specificGenerator.GenerateDataBySchema(ctx, schema)
}

func (generator *coordinatingSchemaGenerator) detectSchemaType(schema *openapi3.Schema) string {
	schemaType := schema.Type
	if mockjsHelper.IsMockJsSchema(schema) {
		return "mockjs"
	}

	switch {
	case schema.OneOf != nil:
		schemaType = "oneOf"
	case schema.AnyOf != nil:
		schemaType = "anyOf"
	case schema.AllOf != nil:
		schemaType = "allOf"
	case schemaType == "" && len(schema.Properties) > 0:
		schemaType = "object"
	case schemaType == "":
		schemaType = "object"
	}

	return schemaType
}
