package model

type LlmTool struct {
	BaseModel

	Name string `json:"name"`
	Desc string `gorm:"type:text" json:"desc"`

	Model   string `json:"model"`
	ApiBase string `json:"apiBase"`
	ApiKey  string `json:"apiKey"`
	Version string `json:"version"`

	IsDefault bool `json:"isDefault"`
	ProjectId uint `json:"projectId"`

	CreateUser string `json:"createUser"`
	UpdateUser string `json:"updateUser"`
}

func (LlmTool) TableName() string {
	return "biz_llm_tool"
}
