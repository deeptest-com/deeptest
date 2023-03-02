package model

type ComponentSchemaSecurity struct {
	BaseModel
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content" gorm:"type:text"`
	ProjectId int64  `json:"project_id"`
	ServeId   int64  `json:"serve_id"`
}

func (ComponentSchemaSecurity) TableName() string {
	return "biz_project_serve_component_security"
}
