package mockContent

import (
	"context"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
)

type Generator interface {
	GenerateContent(ctx context.Context, response *openapi3.Response, contentType string) (interface{}, error)
}

func NewGenerator(generator mockData.MediaGenerator) Generator {
	mediaGenerator := &mediaGenerator{contentGenerator: generator}

	return &delegatingGenerator{
		matchers: []contentMatcher{
			{
				pattern:   regexp.MustCompile("^application/.*json$"),
				generator: mediaGenerator,
			},
			{
				pattern:   regexp.MustCompile("^application/.*xml$"),
				generator: mediaGenerator,
			},
			{
				pattern:   regexp.MustCompile("^text/html$"),
				generator: &htmlGenerator{contentGenerator: generator},
			},
			{
				pattern:   regexp.MustCompile("^text/plain$"),
				generator: &plainTextGenerator{contentGenerator: generator},
			},
		},
	}
}
