package mockResponder

import (
	"context"
	"errors"
	"fmt"
	mockErrors "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/errors"
	mockLogcontext "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/logcontext"
	mockGenerator "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/generator"
	mockSerializer "github.com/deeptest-com/deeptest/internal/pkg/helper/openapi-mock/openapi/responder/serializer"
	"net/http"
	"regexp"
	"strings"
)

type coordinatingResponder struct {
	serializer     mockSerializer.Serializer
	formatGuessers []formatGuess
}

type formatGuess struct {
	format  string
	pattern *regexp.Regexp
}

func (responder *coordinatingResponder) WriteResponse(ctx context.Context, writer http.ResponseWriter,
	response *mockGenerator.Response) {
	format := responder.guessSerializationFormat(response.ContentType.String())

	data, err := responder.serializer.Serialize(response.Data, format)
	if err != nil {
		responder.WriteError(ctx, writer, err)
		return
	}

	if response.ContentType != "" {
		writer.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", response.ContentType))
	}

	writer.WriteHeader(response.StatusCode.Int())
	_, _ = writer.Write(data)
}

func (responder *coordinatingResponder) WriteError(ctx context.Context, writer http.ResponseWriter, err error) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusInternalServerError)

	html := ""

	var unsupported *mockErrors.NotSupported
	if errors.As(err, &unsupported) {
		html = responder.generateUnsupportedErrorHTML(unsupported)
	} else {
		html = responder.generateUnexpectedErrorHTML(ctx, err)
	}

	_, _ = writer.Write([]byte(html))
}

func (responder *coordinatingResponder) generateUnexpectedErrorHTML(ctx context.Context, err error) string {
	html := strings.ReplaceAll(errorTemplate, "{{title}}", "Unexpected error")
	message := "An unexpected error occurred:<br>" + strings.ReplaceAll(err.Error(), ":", ":<br>")
	html = strings.ReplaceAll(html, "{{message}}", message)
	html = strings.ReplaceAll(html, "{{hint}}", errorHint)

	logger := mockLogcontext.LoggerFromContext(ctx)
	logger.Errorf("an unexpected error occurred: %+v", err)

	return html
}

func (responder *coordinatingResponder) generateUnsupportedErrorHTML(err *mockErrors.NotSupported) string {
	html := strings.ReplaceAll(errorTemplate, "{{title}}", "Feature is not supported")
	message := fmt.Sprintf("An error occurred: %s.", err.Error())
	html = strings.ReplaceAll(html, "{{message}}", message)
	html = strings.ReplaceAll(html, "{{hint}}", unsupportedHint)

	return html
}

func (responder *coordinatingResponder) guessSerializationFormat(contentType string) string {
	format := "raw"

	for _, guesser := range responder.formatGuessers {
		if guesser.pattern.MatchString(contentType) {
			format = guesser.format
			break
		}
	}

	return format
}
