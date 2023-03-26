package model

type Debug struct {
	BaseModel
	InvocationBase

	ProcessorId          uint `json:"processorId"`
	ProcessorInterfaceId uint `json:"processorInterfaceId"`
}

func (Debug) TableName() string {
	return "biz_debug"
}
