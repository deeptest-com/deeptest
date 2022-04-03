package model

import (
	"github.com/aaronchen2k/deeptest/internal/comm/consts"
	"github.com/aaronchen2k/deeptest/internal/comm/domain"
)

type TestResponse struct {
	BaseModel

	Code     consts.HttpRespCode     `json:"code"`
	Headers  []domain.Header         `gorm:"-" json:"headers"`
	Body     string                  `gorm:"default:''" json:"body"`
	BodyType consts.HttpRespBodyType `gorm:"default:json" json:"bodyType"`
}

func (TestResponse) TableName() string {
	return "biz_test_response"
}
