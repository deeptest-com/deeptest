package model

type TestRequest struct {
	BaseModel

	Name string `json:"name"`
	Desc string `json:"desc"`

	// config
	Url               string `json:"url"`
	Method            string `gorm:"default:GET" json:"method"`
	Body              string `gorm:"default:{}" json:"body"`
	BodyType          string `gorm:"default:json" json:"bodyType"`
	AuthorizationType string `gorm:"default:''" json:"authorizationType"`
	PreRequestScript  string `gorm:"default:''" json:"preRequestScript"`
	ValidationScript  string `gorm:"default:''" json:"validationScript"`
}

func (TestRequest) TableName() string {
	return "biz_test_request"
}
