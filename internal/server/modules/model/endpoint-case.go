package model

type EndpointCase struct {
	BaseModel
	CreatedBy uint `json:"createdBy"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	EndpointId uint `json:"endpointId"`
	ServeId    uint `json:"serveId"`
	ProjectId  uint `json:"projectId"`

	DebugInterfaceId uint            `gorm:"default:0" json:"debugInterfaceId"`
	DebugData        *DebugInterface `gorm:"-" json:"debugData"`

	SerialNumber   string `json:"serialNumber"`
	CreateUserId   uint   `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
}

func (EndpointCase) TableName() string {
	return "biz_endpoint_case"
}
