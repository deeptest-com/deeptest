package model

import (
	"github.com/deeptest-com/deeptest/internal/pkg/consts"
	serverConsts "github.com/deeptest-com/deeptest/internal/server/consts"
)

type GrpcInterface struct {
	BaseModel

	CreatedBy uint `json:"createdBy"`
	UpdatedBy uint `json:"updatedBy"`

	Name            string `json:"name"`
	Address         string `json:"address"`
	UseTls          *bool  `json:"useTls"`
	RestartConn     *bool  `json:"restartConn"`
	RequestMetadata string `json:"requestMetadata"`

	Service string `gorm:"default:''" json:"service"`
	Method  string `gorm:"default:''" json:"method"`

	Schema   string `json:"schema"`
	Template string `json:"template"`
	Message  string `json:"message"`

	ProtoSrc  string `gorm:"default:'remote'" json:"protoSrc"`
	ProtoName string `json:"protoName"`
	ProtoPath string `json:"protoPath"`
	// ProtoContent    string `json:"protoContent"`

	UsedBy consts.UsedBy                      `json:"usedBy"`
	Type   serverConsts.DiagnoseInterfaceType `json:"type"`

	DiagnoseInterfaceId uint `json:"diagnoseInterfaceId"`
	EnvironmentId       uint `json:"environmentId"`
	ServeId             uint `json:"serveId"`
	ServerId            uint `json:"serverId"`
	ProjectId           uint `json:"projectId"`
}

type GrpcInterfaceRequest struct {
	BaseModel

	CreatedBy uint `json:"createdBy"`
	UpdatedBy uint `json:"updatedBy"`

	Service         string `json:"service"`
	Method          string `json:"method"`
	RequestContent  string `json:"requestContent"`
	ResponseContent string `json:"responseContent"`

	InterfaceId uint `json:"interfaceId"`
}

func (GrpcInterface) TableName() string {
	return "biz_grpc_interface"
}
func (GrpcInterfaceRequest) TableName() string {
	return "biz_grpc_interface_request"
}
