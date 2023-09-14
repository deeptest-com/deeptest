package mockGenerator

import (
	mockContent "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/content"
	mockData "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/data"
	mockNegotiator "github.com/aaronchen2k/deeptest/internal/pkg/helper/openapi-mock/openapi/generator/negotiator"
	"net/http"

	"github.com/getkin/kin-openapi/routers"
)

type ResponseGenerator interface {
	GenerateResponse(request *http.Request, route *routers.Route, code string) (*Response, error)
}

func New(dataGenerator mockData.MediaGenerator) ResponseGenerator {
	return &coordinatingGenerator{
		contentTypeNegotiator: mockNegotiator.NewContentTypeNegotiator(),
		statusCodeNegotiator:  mockNegotiator.NewStatusCodeNegotiator(),
		contentGenerator:      mockContent.NewGenerator(dataGenerator),
	}
}
