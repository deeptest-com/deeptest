package serverDomain

type MockReqJson struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type MockJsExpression struct {
	ID uint `json:"id"`

	Name       string `json:"name"`
	Expression string `json:"expression"`
	Format     string `json:"format"`
	Desc       string `json:"desc"`
	Ordr       int    `json:"ordr"`

	Result interface{} `gorm:"-" json:"result"`
}
