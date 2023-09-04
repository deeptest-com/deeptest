package model

type EndpointTag struct {
	BaseModel
	Name      string `gorm:"index:name_project_index,unique;not null;type:varchar(200)" json:"name"`
	ProjectId uint   `gorm:"index:name_project_index,unique;not null;type:varchar(200)" json:"projectId"`
}

func (EndpointTag) TableName() string {
	return "biz_endpoint_tag"
}
