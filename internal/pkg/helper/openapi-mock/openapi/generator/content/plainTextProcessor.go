package mockContent

import (
	"context"
	mockLogcontext "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/logcontext"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	"github.com/getkin/kin-openapi/openapi3"
)

type plainTextGenerator struct {
	contentGenerator mockData.MediaGenerator
}

func (generator *plainTextGenerator) GenerateContent(ctx context.Context, response *openapi3.Response, contentType string) (interface{}, error) {
	originMediaType := response.Content.Get(contentType)

	schema := originMediaType.Schema

	if schema == nil || schema.Value.Type != "string" {
		logger := mockLogcontext.LoggerFromContext(ctx)
		logger.Warnf("only string schema is supported for '%s' content type", contentType)

		schema = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: "string",
			},
		}
	}

	mediaType := &openapi3.MediaType{
		Schema:   schema,
		Example:  originMediaType.Example,
		Examples: originMediaType.Examples,
	}

	return generator.contentGenerator.GenerateData(ctx, mediaType)
}
