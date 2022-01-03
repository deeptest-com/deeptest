package model

type ProjectMember struct {
	BaseModel

	ProjectId     uint `json:"projectId"`
	ProjectRoleId uint `json:"projectRoleId"`
	UserId        uint `json:"userId"`
}

func (ProjectMember) TableName() string {
	return "biz_project_member"
}
