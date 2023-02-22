package model

type ServeVersion struct {
	BaseModel
	Description string `json:"description"`
	CreateUser  string `json:"createUser"`
	Value       string `json:"value"`
	ServeId     int64  `json:"serveId"`
}

func (ServeVersion) TableName() string {
	return "biz_project_serve_version"
}
