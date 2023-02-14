package model

type ServeVersion struct {
	BaseModel
	Description string `json:"description"`
	UserId      int64  `json:"userId"`
	Value       string `json:"value"`
	ServeId     int64  `json:"serve_id"`
}

func (ServeVersion) TableName() string {
	return "biz_project_serve_version"
}
