package model

type Endpoint struct {
	BaseModel
	Title        string              `json:"title"`
	ProjectId    uint                `json:"projectId"`
	ServeId      uint                `json:"serveId"`
	ServerId     uint                `json:"serverId"`
	Path         string              `json:"path"`
	Version      string              `json:"version"`
	CreateUser   string              `json:"createUser"`
	Status       int64               `json:"status"`
	CategoryId   uint                `json:"categoryId"`
	PathParams   []EndpointPathParam `gorm:"-" json:"pathParams"`
	Interfaces   []EndpointInterface `gorm:"-" json:"interfaces"`
	Versions     []EndpointVersion   `gorm:"-" json:"versions"`
	ServeName    string              `gorm:"-" json:"serveName"`
	Description  string              `json:"description"`
	SerialNumber string              `json:"serialNumber"`
}

func (Endpoint) TableName() string {
	return "biz_endpoint"
}

type EndpointPathParam struct {
	EndpointInterfaceParam
	EndpointId uint `json:"endpointId"`
}

func (EndpointPathParam) TableName() string {
	return "biz_endpoint_path_param"
}

type EndpointVersion struct {
	BaseModel
	Version    string `json:"version"`
	EndpointId uint   `json:"endpointId"`
}

func (EndpointVersion) TableName() string {
	return "biz_endpoint_version"
}
