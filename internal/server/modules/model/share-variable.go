package model

import "github.com/aaronchen2k/deeptest/internal/pkg/consts"

type ShareVariable struct {
	BaseModel

	Name      string                     `json:"name"`
	Value     string                     `gorm:"type:text" json:"value"`
	ValueType consts.ExtractorResultType `json:"valueType"`

	InvokeId            uint `gorm:"default:0" json:"invokeId"`
	DebugInterfaceId    uint `gorm:"default:0" json:"debugInterfaceId"`
	CaseInterfaceId     uint `gorm:"default:0" json:"caseInterfaceId"`
	EndpointInterfaceId uint `json:"endpointInterfaceId"`
	ServeId             uint `json:"serveId"` // for interface debug

	ScenarioProcessorId uint `gorm:"default:0" json:"scenarioProcessorId"` // for scenario
	ScenarioId          uint `gorm:"default:0" json:"scenarioId"`          // for scenario

	Scope  consts.ExtractorScope `json:"scope" gorm:"default:private"` // debug
	UsedBy consts.UsedBy         `json:"usedBy"`
}

func (ShareVariable) TableName() string {
	return "biz_share_variable"
}
