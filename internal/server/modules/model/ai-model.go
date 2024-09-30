package model

import "github.com/deeptest-com/deeptest/internal/pkg/domain"

type AiModel struct {
	BaseModel

	domain.AiModelBase
}

func (AiModel) TableName() string {
	return "ai_model"
}
