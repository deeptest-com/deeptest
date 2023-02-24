package model

type Serve struct {
	BaseModel
	Name        string `json:"name"`
	ProjectId   uint   `json:"projectId"`
	CreateUser  string `json:"createUser"`
	Status      int64  `json:"status"`
	Description string `json:"description"`
}

func (Serve) TableName() string {
	return "biz_project_serve"
}
