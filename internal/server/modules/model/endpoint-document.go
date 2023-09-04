package model

type EndpointDocument struct {
	BaseModel
	Name      string `json:"name"`
	Version   string `gorm:"index:version_project_index,unique;not null;type:varchar(200)" json:"version"`
	ProjectId uint   `gorm:"index:version_project_index,unique;not null;type:varchar(200)" json:"projectId"`
}

func (EndpointDocument) TableName() string {
	return "biz_endpoint_document"
}
