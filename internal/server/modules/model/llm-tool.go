package model

import "github.com/deeptest-com/deeptest/internal/pkg/domain"

type LlmTool struct {
	BaseModel

	domain.ToolModelBase

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (LlmTool) TableName() string {
	return "biz_llm_tool"
}
