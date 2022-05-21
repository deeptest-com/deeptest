package model

type TestScenario struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor TestProcessor `json:"processor" yaml:"processor"`
}

func (TestScenario) TableName() string {
	return "test_scenario"
}
