package model

type EndpointTagRel struct {
	BaseModel
	EndpointId uint   `gorm:"index:endpoint_tag_project_index,unique;not null;type:varchar(200)" json:"endpointId"`
	TagId      uint   `json:"tagId"`
	TagName    string `gorm:"index:endpoint_tag_project_index,unique;not null;type:varchar(200)" json:"tagName"`
	ProjectId  uint   `gorm:"index:endpoint_tag_project_index,unique;not null;type:varchar(200)" json:"projectId"`
}

func (EndpointTagRel) TableName() string {
	return "biz_endpoint_tag_rel"
}
