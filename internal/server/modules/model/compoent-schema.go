package model

type ComponentSchema struct {
	BaseModel
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	ProjectId int64  `json:"project_id"`
	ServeId   int64  `json:"serve_id"`
}

func (ComponentSchema) TableName() string {
	return "biz_component_schema"
}
