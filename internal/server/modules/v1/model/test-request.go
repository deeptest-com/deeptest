package model

type TestRequest struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc,omitempty"`

	ReqContent  string `json:"reqContent,omitempty"`
	RespContent string `json:"respContent,omitempty"`

	InterfaceId uint `json:"interfaceId,omitempty"`
	ProjectId   uint `json:"projectId,omitempty"`
}

func (TestRequest) TableName() string {
	return "biz_test_request"
}
