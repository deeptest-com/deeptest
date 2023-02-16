package model

type ServeServer struct {
	BaseModel
	Name        string `json:"name"`
	ServerId    uint   `json:"serverId"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func (ServeServer) TableName() string {
	return "biz_project_serve_server"
}
