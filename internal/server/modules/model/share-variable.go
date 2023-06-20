package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ShareVariable struct {
	BaseModel

	Name  string `json:"name"`
	Value string `json:"value"`

	InterfaceId uint `json:"interfaceId"`
	ServeId     uint `json:"serveId"` // for interface debug

	ScenarioProcessorId uint `gorm:"default:0" json:"scenarioProcessorId"` // for scenario
	ScenarioId          uint `gorm:"default:0" json:"scenarioId"`          // for scenario

	Scope consts.ExtractorScope `json:"scope" gorm:"default:private"` // debug
}

func (ShareVariable) TableName() string {
	return "biz_share_variable"
}
