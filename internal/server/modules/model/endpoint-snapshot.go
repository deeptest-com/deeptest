package model

type EndpointSnapshot struct {
	BaseModel
	EndpointId uint   `json:"endpointId"`
	DocumentId uint   `gorm:"index:document_id_index;not null" json:"documentId"`
	Content    string `gorm:"type:longtext" json:"content"`
}

func (EndpointSnapshot) TableName() string {
	return "biz_endpoint_snapshot"
}
