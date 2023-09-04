package model

type MockJsExpression struct {
	BaseModel

	Name       string `json:"name"`
	Expression string `json:"expression"`
	Format     string `json:"format"`
	Desc       string `json:"desc"`
	Ordr       int    `json:"ordr"`

	Result interface{} `gorm:"-" json:"result"`
}

func (MockJsExpression) TableName() string {
	return "biz_mock_js_expression"
}
