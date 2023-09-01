package mockContent

import (
	"context"
	"fmt"
	mockErrors "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/errors"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
)

type delegatingGenerator struct {
	matchers []contentMatcher
}

type contentMatcher struct {
	pattern   *regexp.Regexp
	generator Generator
}

func (processor *delegatingGenerator) GenerateContent(ctx context.Context, response *openapi3.Response, contentType string) (interface{}, error) {
	if contentType == "" {
		return "", nil
	}

	for _, matcher := range processor.matchers {
		if matcher.pattern.MatchString(contentType) {
			return matcher.generator.GenerateContent(ctx, response, contentType)
		}
	}

	return nil, mockErrors.NewNotSupported(fmt.Sprintf("generating response for content type '%s' is not supported", contentType))
}
