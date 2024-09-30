package model

import "github.com/aaronchen2k/deeptest/internal/pkg/domain"

type AiTemplate struct {
	BaseModel

	domain.AiTemplateBase
}

func (AiTemplate) TableName() string {
	return "ai_template"
}
