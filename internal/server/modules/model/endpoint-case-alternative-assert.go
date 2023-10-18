package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type EndpointCaseAlternativeAssert struct {
	BaseModel

	AlternativeCaseId uint `json:"alternativeCaseId"`

	domain.CheckpointBase
	Disabled bool `json:"disabled,omitempty"` // used by gorm fetch

	Desc string `json:"desc" gorm:"type:text"`
	Ordr int    `json:"ordr"`
}

func (EndpointCaseAlternativeAssert) TableName() string {
	return "biz_endpoint_case_alternative_assert"
}
