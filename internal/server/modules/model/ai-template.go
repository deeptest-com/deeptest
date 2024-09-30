package model

import "github.com/deeptest-com/deeptest/internal/pkg/domain"

type AiTemplate struct {
	BaseModel

	domain.AiTemplateBase
}

func (AiTemplate) TableName() string {
	return "ai_template"
}
