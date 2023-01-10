package model

type ProcessorInvocation struct {
	BaseModel
	InvocationBase

	ProcessorId          uint `json:"processorId"`
	ProcessorInterfaceId uint `json:"processorInterfaceId"`
}

func (ProcessorInvocation) TableName() string {
	return "biz_processor_invocation"
}
