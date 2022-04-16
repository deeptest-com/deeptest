package model

type TestRequest struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	ReqContent  string `json:"reqContent"`
	RespContent string `json:"respContent"`

	InterfaceId uint `json:"interfaceId"`
	ProjectId   uint `json:"projectId"`
}

func (TestRequest) TableName() string {
	return "biz_test_request"
}
