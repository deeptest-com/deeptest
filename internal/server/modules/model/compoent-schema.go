package model

type ComponentSchema struct {
	BaseModel
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	ServeId  int64  `json:"serveId"`
	Examples string `json:"examples"`
	Tags     string `json:"tags"`
}

func (ComponentSchema) TableName() string {
	return "biz_project_serve_component_schema"
}
