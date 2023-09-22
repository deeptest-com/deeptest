package model

type ProjectUserServer struct {
	BaseModel
	ProjectId uint `json:"projectId"`
	UserId    uint `json:"userId"`
	ServerId  uint `json:"serverId"`
}

func (ProjectUserServer) TableName() string {
	return "biz_project_user_server"
}
