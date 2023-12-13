package model

type ComponentSchemaSecurity struct {
	BaseModel
	Name             string `json:"name"`
	Type             string `json:"type"`
	ProjectId        int64  `json:"project_id"`
	ServeId          int64  `json:"serve_id"`
	Description      string `gorm:"type:text" json:"description"`
	In               string `json:"in"`
	Scheme           string `json:"scheme"`
	BearerFormat     string `json:"bearerFormat"`
	Flows            string `json:"flows,omitempty" gorm:"type:text"`
	OpenIdConnectUrl string `json:"openIdConnectUrl"`
	Key              string `json:"key"`
	Value            string `json:"value"`
	Token            string `json:"token"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Default          bool   `json:"default"`
}

func (ComponentSchemaSecurity) TableName() string {
	return "biz_project_serve_component_security"
}
