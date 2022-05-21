package model

type TestSet struct {
	BaseModel

	TestConfig

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor TestProcessor `json:"processor" yaml:"processor"`
}

func (TestSet) TableName() string {
	return "test_set"
}
