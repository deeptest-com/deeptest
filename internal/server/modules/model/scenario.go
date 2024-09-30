package model

import "github.com/deeptest-com/deeptest/internal/pkg/consts"

type Scenario struct {
	BaseModel

	Version float64 `json:"version" yaml:"version"`
	Name    string  `json:"name" yaml:"name"`
	Desc    string  `json:"desc" yaml:"desc"`

	Processor      Processor         `json:"processor" yaml:"processor" gorm:"-"`
	CategoryId     int64             `json:"categoryId"`
	ServeId        uint              `json:"serveId"`
	ProjectId      uint              `json:"projectId"`
	SerialNumber   string            `json:"serialNumber"`
	Status         consts.TestStatus `json:"status"`
	CreateUserId   uint              `json:"createUserId"`
	CreatorName    string            `gorm:"-" json:"creatorName"`
	CreateUserName string            `json:"createUserName"` //用户登录名
	UpdateUserId   uint              `json:"updateUserId"`
	UpdateUserName string            `json:"updateUserName"`
	Priority       string            `json:"priority"`
	Type           consts.TestType   `json:"type"`
	CurrEnvId      uint              `json:"currEnvId"`
}

func (Scenario) TableName() string {
	return "biz_scenario"
}

type ScenarioDetail struct {
	Scenario
	CategoryName string `json:"categoryName"`
	RefId        uint   `json:"refId"`
}
