package mockContent

import (
	"context"
	mockLogcontext "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/logcontext"
	mockData "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	"github.com/getkin/kin-openapi/openapi3"
)

type htmlGenerator struct {
	contentGenerator mockData.MediaGenerator
}

func (generator *htmlGenerator) GenerateContent(ctx context.Context, response *openapi3.Response, contentType string) (interface{}, error) {
	originMediaType := response.Content.Get(contentType)

	schema := originMediaType.Schema

	if schema == nil || schema.Value.Type != "string" || schema.Value.Format != "html" {
		logger := mockLogcontext.LoggerFromContext(ctx)
		logger.Warnf("only string schema with html format is supported for '%s' content type", contentType)

		schema = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:   "string",
				Format: "html",
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
