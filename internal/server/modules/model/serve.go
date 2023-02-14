package model

type Serve struct {
	BaseModel
	Name        string `json:"name"`
	ProjectId   uint   `json:"projectId"`
	UserId      int64  `json:"userId"`
	Status      int64  `json:"status"`
	Description string `json:"description"`
}

func (Serve) TableName() string {
	return "biz_project_serve"
}
