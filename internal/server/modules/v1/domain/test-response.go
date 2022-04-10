package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
)

type TestResponse struct {
	StatusCode    consts.HttpRespCode `json:"statusCode"`
	StatusContent string              `json:"statusContent"`

	Headers     []domain.Header        `gorm:"-" json:"headers"`
	Content     string                 `gorm:"default:''" json:"content"`
	ContentType consts.HttpContentType `json:"contentType"`

	ContentLang    consts.HttpRespLangType `json:"contentLang"`
	ContentCharset consts.HttpRespCharset  `json:"contentCharset"`
	ContentLength  int                     `json:"contentLength"`

	Time int64 `json:"time"`
}
