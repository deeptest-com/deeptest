package mockContent

import (
	"context"
	mockData "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"

	"github.com/getkin/kin-openapi/openapi3"
)

type mediaGenerator struct {
	contentGenerator mockData.MediaGenerator
}

func (generator *mediaGenerator) GenerateContent(ctx context.Context, response *openapi3.Response, contentType string) (interface{}, error) {
	mediaType := response.Content[contentType]
	if mediaType == nil {
		return "", nil
	}

	return generator.contentGenerator.GenerateData(ctx, mediaType)
}
