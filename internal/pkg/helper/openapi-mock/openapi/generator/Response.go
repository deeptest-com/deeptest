package mockGenerator

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	"github.com/deeptest-com/deeptest/internal/pkg/domain"
)

type Request struct {
	Method consts.HttpMethod `json:"method"`

	Url         string              `json:"url"`
	QueryParams []domain.Param      `json:"queryParams,omitempty"`
	PathParams  []domain.Param      `json:"pathParams,omitempty"`
	Headers     []domain.Param      `json:"headers,omitempty"`
	Cookies     []domain.ExecCookie `json:"cookies,omitempty"`

	Body     string                    `json:"body,omitempty"`
	FormData []domain.BodyFormDataItem `json:"formData,omitempty"`
}

type Response struct {
	StatusCode consts.HttpRespCode `json:"statusCode"`

	ContentType consts.HttpContentType `json:"contentType"`
	Headers     []domain.Param         `json:"headers,omitempty"`

	Data interface{} `json:"data,omitempty"`

	// used by adv mock
	UseAdvMock bool   `json:"useAdvMock,omitempty"`
	Content    string `json:"content"`
	DelayTime  uint   `json:"delayTime,omitempty"`
}
