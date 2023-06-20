package model

type MockInvocation struct {
	BaseModel
	InvocationBase
}

func (MockInvocation) TableName() string {
	return "biz_mock_invocation"
}
