package model

type ServeEndpointVersion struct {
	BaseModel
	EndpointId      int64  `json:"EndpointId"`
	EndpointVersion string `json:"endpointVersion"`
	ServeVersion    string `json:"ServeVersion"`
	ServeId         int64  `json:"serveId"`
}

func (ServeEndpointVersion) TableName() string {
	return "biz_project_serve_endpoint_version"
}
