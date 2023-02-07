package model

type Endpoint struct {
	BaseModel
	Id         int64               `json:"id"`
	Title      string              `json:"title"`
	ProjectId  int64               `json:"project_id"`
	ParentId   int64               `json:"parent_id"`
	PathParams []EndpointPathParam `gorm:"-" json:"path_params"`
	Interfaces []Interface         `gorm:"-" json:"interfaces"`
}

func (Endpoint) TableName() string {
	return "biz_endpoint"
}

type EndpointPathParam struct {
	BaseModel
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Value      string `json:"type"`
	Type       string `json:"type"`
	EndpointId int64  `json:"endpoint_id"`
}

func (EndpointPathParam) TableName() string {
	return "biz_endpoint_path_param"
}
