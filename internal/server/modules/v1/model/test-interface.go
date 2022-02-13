package model

type TestInterface struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	ParentID  uint `json:"parentId"`
	ProjectId uint `json:"projectId"`
	UseID     uint `json:"useId"`

	Ordr     int              `json:"ordr"`
	Children []*TestInterface `json:"fields"`
}

func (TestInterface) TableName() string {
	return "biz_test_interface"
}
