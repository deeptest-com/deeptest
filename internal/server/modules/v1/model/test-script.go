package model

type TestScript struct {
	BaseModel

	Version int `json:"version"`
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`

	ProjectId    uint `json:"projectId"`
}

func (TestScript) TableName() string {
	return "biz_test_script"
}
