package request

import (
	"bytes"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var fieldTags = []string{"proto", "status_code", "headers", "cookies", "body", textExtractorSubRegexp}

type httpRespObjMeta struct {
	Proto      string            `json:"proto"`
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Cookies    map[string]string `json:"cookies"`
	Body       interface{}       `json:"body"`
}

func newHttpResponseObject(t *testing.T, parser *Parser, resp *http.Response) (*responseObject, error) {
	// prepare response headers
	headers := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	// prepare response cookies
	cookies := make(map[string]string)
	for _, cookie := range resp.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	// read response body
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse response body
	var body interface{}
	if err := json.Unmarshal(respBodyBytes, &body); err != nil {
		// response body is not json, use raw body
		body = string(respBodyBytes)
	}

	respObjMeta := httpRespObjMeta{
		Proto:      resp.Proto,
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Cookies:    cookies,
		Body:       body,
	}

	return convertToResponseObject(t, parser, respObjMeta)
}

type wsCloseRespObject struct {
	StatusCode int    `json:"status_code"`
	Text       string `json:"body"`
}

func newWsCloseResponseObject(t *testing.T, parser *Parser, resp *wsCloseRespObject) (*responseObject, error) {
	return convertToResponseObject(t, parser, resp)
}

type wsReadRespObject struct {
	Message     interface{} `json:"body"`
	messageType int
}

func newWsReadResponseObject(t *testing.T, parser *Parser, resp *wsReadRespObject) (*responseObject, error) {
	byteMessage, ok := resp.Message.([]byte)
	if !ok {
		return nil, errors.New("websocket message type should be []byte")
	}
	var msg interface{}
	if err := json.Unmarshal(byteMessage, &msg); err != nil {
		// response body is not json, use raw body
		msg = string(byteMessage)
	}
	resp.Message = msg
	return convertToResponseObject(t, parser, resp)
}

func convertToResponseObject(t *testing.T, parser *Parser, respObjMeta interface{}) (*responseObject, error) {
	respObjMetaBytes, _ := json.Marshal(respObjMeta)
	var data interface{}
	decoder := json.NewDecoder(bytes.NewReader(respObjMetaBytes))
	decoder.UseNumber()
	if err := decoder.Decode(&data); err != nil {
		log.Error().
			Str("respObjectMeta", string(respObjMetaBytes)).
			Err(err).
			Msg("[convertToResponseObject] convert respObjectMeta to interface{} failed")
		return nil, err
	}
	return &responseObject{
		t:           t,
		parser:      parser,
		respObjMeta: data,
	}, nil
}

type responseObject struct {
	t                 *testing.T
	parser            *Parser
	respObjMeta       interface{}
	validationResults []*ValidationResult
}

const textExtractorSubRegexp string = `(.*)`
