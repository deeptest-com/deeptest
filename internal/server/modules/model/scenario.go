package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type Scenario struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor      Processor         `json:"processor" yaml:"processor" gorm:"-"`
	CategoryId     uint              `json:"categoryId"`
	ServeId        uint              `json:"serveId"`
	ProjectId      uint              `json:"projectId"`
	SerialNumber   string            `json:"serialNumber"`
	Status         consts.TestStatus `json:"status"`
	CreateUserId   uint              `json:"createUserId"`
	CreateUserName string            `json:"createUserName"`
	Priority       string            `json:"priority"`
	Type           consts.TestType   `json:"type"`
}

func (Scenario) TableName() string {
	return "biz_scenario"
}

type ScenarioDetail struct {
	Scenario
	CategoryName string `json:"categoryName"`
}
