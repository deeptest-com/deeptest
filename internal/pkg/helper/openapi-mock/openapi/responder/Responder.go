package mockResponder

import (
	"context"
	mockGenerator "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	mockSerializer "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/responder/serializer"
	"net/http"
	"regexp"
)

type Responder interface {
	WriteResponse(ctx context.Context, writer http.ResponseWriter, response *mockGenerator.Response)
	WriteError(ctx context.Context, writer http.ResponseWriter, err error)
}

func New() Responder {
	return &coordinatingResponder{
		serializer: mockSerializer.New(),
		formatGuessers: []formatGuess{
			{
				format:  "json",
				pattern: regexp.MustCompile("^application/.*json$"),
			},
			{
				format:  "xml",
				pattern: regexp.MustCompile("^application/.*xml$"),
			},
		},
	}
}
