package model

type TestScript struct {
	BaseModel

	Version int    `json:"version"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Desc    string `json:"desc"`

	ProjectId uint       `json:"projectId"`
	Steps     []TestStep `json:"steps" gorm:"foreignKey:script_id"`
}

func (TestScript) TableName() string {
	return "biz_test_script"
}
