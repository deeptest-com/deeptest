package serverDomain

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
)

type TestRequest struct {
	Id uint `gorm:"-" json:"id"`

	Url               string                 `json:"url"`
	Method            consts.HttpMethod      `gorm:"default:GET" json:"method"`
	Params            []domain.Param         `gorm:"-" json:"params"`
	Headers           []domain.Header        `gorm:"-" json:"headers"`
	Body              string                 `gorm:"default:{}" json:"body"`
	BodyType          consts.HttpContentType `gorm:"default:json" json:"bodyType"`
	AuthorizationType string                 `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string                 `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string                 `gorm:"default:''" json:"validationScript"`

	BasicAuth   domain.BasicAuth   `gorm:"-" json:"basicAuth"`
	BearerToken domain.BearerToken `gorm:"-" json:"bearerToken"`
	OAuth20     domain.OAuth20     `gorm:"-" json:"oAuth20"`
	ApiKey      domain.ApiKey      `gorm:"-" json:"apiKey"`
}
