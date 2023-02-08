package model

type Endpoint struct {
	BaseModel
	Title      string              `json:"title"`
	ProjectId  int64               `json:"projectId"`
	ParentId   int64               `json:"parentId"`
	Path       string              `json:"path"`
	Version    string              `json:"version"`
	PathParams []EndpointPathParam `gorm:"-" json:"pathParams"`
	Interfaces []Interface         `gorm:"-" json:"interfaces"`
}

func (Endpoint) TableName() string {
	return "biz_endpoint"
}

type EndpointPathParam struct {
	BaseModel
	Name       string `json:"name"`
	Value      string `json:"type"`
	Type       string `json:"type"`
	EndpointId uint   `json:"endpointId"`
}

func (EndpointPathParam) TableName() string {
	return "biz_endpoint_path_param"
}
