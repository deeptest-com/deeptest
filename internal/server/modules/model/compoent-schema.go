package model

type ComponentSchema struct {
	BaseModel
	Name        string `json:"name"`
	Type        string `json:"type"`
	Content     string `json:"content" gorm:"type:text"`
	ServeId     int64  `json:"serveId"`
	Examples    string `json:"examples" gorm:"type:text"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
	Ref         string `json:"ref"`
}

func (ComponentSchema) TableName() string {
	return "biz_project_serve_component_schema"
}
