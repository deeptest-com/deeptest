package model

type WebsocketInterface struct {
	BaseModel

	Name    string `json:"name"`
	Message string `json:"message"`

	Address      string `json:"address"`
	ExtMode      *bool  `json:"extMode"`
	Namespace    string `json:"Namespace"`
	Room         string `json:"room"`
	Event        string `json:"event"`
	ListenEvents string `json:"listenEvents"`

	CreatedBy uint `json:"createdBy"`
	UpdatedBy uint `json:"updatedBy"`

	DiagnoseInterfaceId uint `json:"diagnoseInterfaceId"`
	ProjectId           uint `json:"projectId"`
	UseID               uint `json:"useId"`
}

func (WebsocketInterface) TableName() string {
	return "biz_websocket_interface"
}

type WebsocketInterfaceParam struct {
	BaseModel
	InterfaceParamBase
}

func (WebsocketInterfaceParam) TableName() string {
	return "biz_websocket_interface_param"
}

type WebsocketInterfaceHeader struct {
	BaseModel
	InterfaceHeaderBase
}

func (WebsocketInterfaceHeader) TableName() string {
	return "biz_websocket_interface_header"
}

type WebsocketInterfaceRequest struct {
	BaseModel

	CreatedBy uint `json:"createdBy"`
	UpdatedBy uint `json:"updatedBy"`

	// adv mode
	Namespace string `json:"Namespace"`
	Room      string `json:"room"`
	EmitEvent string `json:"emitEvent"`

	RequestContent  string `gorm:"type:text" json:"requestContent"`
	ResponseContent string `gorm:"type:text" json:"responseContent"` // multi-line

	InterfaceId uint `json:"interfaceId"`
}

func (WebsocketInterfaceRequest) TableName() string {
	return "biz_websocket_interface_request"
}

type WebsocketCache struct {
	BaseModel

	CreatedBy uint `json:"createdBy"`
	UpdatedBy uint `json:"updatedBy"`

	SerialNumber string `json:"serialNumber"`
	Name         string `json:"name"`
	Desc         string `gorm:"type:text" json:"desc"`
	Data         string `gorm:"type:text" json:"data"`
}

func (WebsocketCache) TableName() string {
	return "biz_websocket_cache"
}
