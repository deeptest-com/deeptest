package mockGenerator

import (
	mockContent "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/content"
	mockNegotiator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/negotiator"
	"net/http"

	"github.com/getkin/kin-openapi/routers"
	"github.com/pkg/errors"
)

type coordinatingGenerator struct {
	statusCodeNegotiator  mockNegotiator.StatusCodeNegotiator
	contentTypeNegotiator mockNegotiator.ContentTypeNegotiator
	contentGenerator      mockContent.Generator
}

func (generator *coordinatingGenerator) GenerateResponse(request *http.Request, route *routers.Route, code string) (*Response, error) {
	responseKey, statusCode, err := generator.statusCodeNegotiator.NegotiateStatusCode(request, route.Operation.Responses, code)
	if err != nil {
		return nil, errors.WithMessage(err, "[coordinatingGenerator] failed to negotiate response")
	}

	bestResponse := route.Operation.Responses[responseKey].Value
	contentType := generator.contentTypeNegotiator.NegotiateContentType(request, bestResponse)

	contentData, err := generator.contentGenerator.GenerateContent(request.Context(), bestResponse, contentType)
	if err != nil {
		return nil, errors.WithMessage(err, "[coordinatingGenerator] failed to generate response data")
	}

	response := &Response{
		StatusCode:  statusCode,
		ContentType: contentType,
		Data:        contentData,
	}

	return response, nil
}
