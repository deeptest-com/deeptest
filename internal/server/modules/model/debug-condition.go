package model

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
)

type DebugPreCondition struct {
	BaseModel
	DebugReferIds

	EntityType consts.ConditionType `json:"entityType"`
	EntityId   uint                 `json:"entityId"`

	Name string `json:"name"`
	Desc string `json:"desc"`
	Ordr int    `json:"ordr"`
}

func (DebugPreCondition) TableName() string {
	return "biz_debug_condition_pre"
}

type DebugPostCondition struct {
	BaseModel
	DebugReferIds

	EntityType consts.ConditionType `json:"entityType"`
	EntityId   uint                 `json:"entityId"`

	Name string `json:"name"`
	Desc string `json:"desc"`
	Ordr int    `json:"ordr"`
}

func (DebugPostCondition) TableName() string {
	return "biz_debug_condition_post"
}

type DebugConditionExtractor struct {
	BaseModel
	ConditionId uint `json:"conditionId"`

	domain.ExtractorBase

	Scope consts.ExtractorScope `json:"scope" gorm:"default:public"`
}

func (DebugConditionExtractor) TableName() string {
	return "biz_debug_condition_extractor"
}

type DebugConditionCheckpoint struct {
	BaseModel
	ConditionId uint `json:"conditionId"`

	Type consts.CheckpointType `json:"type"`

	Expression        string `json:"expression"`
	ExtractorVariable string `json:"extractorVariable"`

	Operator consts.ComparisonOperator `json:"operator"`
	Value    string                    `json:"value"`

	ActualResult string              `json:"actualResult"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
}

func (DebugConditionCheckpoint) TableName() string {
	return "biz_debug_condition_checkpoint"
}

type DebugConditionScript struct {
	BaseModel
	ConditionId uint `json:"conditionId"`

	ConditionType consts.ConditionType `json:"conditionType"` // per | post
	Content       string               `json:"content"`

	Output       string              `json:"output"`
	ResultStatus consts.ResultStatus `json:"resultStatus"`
}

func (DebugConditionScript) TableName() string {
	return "biz_debug_condition_script"
}

type DebugReferIds struct {
	UsedBy consts.UsedBy `json:"usedBy"`

	DebugInterfaceId uint `gorm:"default:0" json:"debugInterfaceId"`

	// debug for Endpoint Interface
	EndpointInterfaceId uint `gorm:"default:0" json:"endpointInterfaceId"`

	// debug in Scenario Processor
	ScenarioProcessorId uint `gorm:"default:0" json:"scenarioProcessorId"`
	ScenarioId          uint `gorm:"default:0" json:"scenarioId"`

	// debug for Test Interface
	DiagnoseInterfaceId uint `gorm:"default:0" json:"diagnoseInterfaceId"`
}
