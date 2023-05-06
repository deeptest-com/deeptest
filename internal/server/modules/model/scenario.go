package model

type Scenario struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor    Processor `json:"processor" yaml:"processor" gorm:"-"`
	CategoryId   uint      `json:"categoryId"`
	ServeId      uint      `json:"serveId"`
	ProjectId    uint      `json:"projectId"`
	SerialNumber string    `json:"serialNumber"`
}

func (Scenario) TableName() string {
	return "biz_scenario"
}
